package app

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/config"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/ipfs"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/network"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/database/filesystem"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/dicebear"
	httpDicebear "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/dicebear/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/ethereum"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/ethereum/crypto"
	merkletree "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/hash"
	httpDefault "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/arbitrum"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon"
	contractPolygon "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/contract"
	rpcPolygon "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/rpc"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/contract"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/rpc"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/web3storage"
	"log"
	"os"
	"time"
)

type Container struct {
	PlateausHTTPClient *http.Client
	PlateausRPCClient  *rpc.Client
	LotteryManager     *lottery.Manager
	NetworkDelegator   *network.Delegator
	IPFSClient         ipfs.Client
}

var container Container

// TODO: decouple this container.go in many other providers separated by context
func init() {
	ctx := context.Background()
	cfg := config.GetConfig()
	dataPath := "data/data.json"
	_, err := os.Stat(dataPath)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dr := filesystem.NewDataRepository(dataPath)

	c := &httpDefault.Client{}
	c.Timeout = 10 * time.Second
	container.PlateausHTTPClient = http.NewClient(c)

	clientPlateaus, err := ethereum.Dial(cfg.PlateausRPC)

	if err != nil {
		log.Printf("could not ethereum.Dial on Plateaus: %s", err)
		os.Exit(1)
	}

	chainIdPlateaus, err := clientPlateaus.ChainID(ctx)

	if err != nil {
		log.Printf("could not client.ChainID: %s", err)
		os.Exit(1)
	}

	addressLotteryPlateaus := common.HexToAddress(cfg.LotteryContractAddress)
	lotteryContract, err := contract.NewLottery(addressLotteryPlateaus, clientPlateaus)

	if err != nil {
		log.Printf("could not contract.NewLottery: %s", err)
		os.Exit(1)
	}

	fromAddressPlateaus, privateKeyPlateaus, err := crypto.AddressFromPrivateKey(cfg.PlateausPrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	container.PlateausRPCClient = rpc.New(chainIdPlateaus, lotteryContract, *fromAddressPlateaus, privateKeyPlateaus)

	tf := merkletree.NewFactory()

	s := lottery.NewService(container.PlateausHTTPClient, container.PlateausRPCClient, dr, tf)
	r := lottery.NewRegisterService(container.PlateausHTTPClient, container.PlateausRPCClient, tf)

	// RPC Polygon
	clientPolygon, err := ethereum.Dial(cfg.PolygonRPC)

	if err != nil {
		log.Printf("could not ethereum.Dial on Polygon: %s", err)
		os.Exit(1)
	}

	chainIdPolygon, err := clientPolygon.ChainID(ctx)

	if err != nil {
		log.Printf("could not client.ChainID: %s", err)
		os.Exit(1)
	}

	fromAddressPolygon, privateKeyPolygon, err := crypto.AddressFromPrivateKey(cfg.PlateausPrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	addressLotteryPolygon := common.HexToAddress(cfg.LotteryValidationContractAddress)
	lotteryValidationContract, err := contractPolygon.NewLotteryValidation(addressLotteryPolygon, clientPolygon)

	if err != nil {
		log.Printf("could not contract.NewLotteryValidation: %s", err)
		os.Exit(1)
	}

	polygonRPC := rpcPolygon.New(clientPolygon, chainIdPolygon, lotteryValidationContract, *fromAddressPolygon, privateKeyPolygon)

	var networkServices = []network.Delegated{
		arbitrum.Service{},
		polygon.NewService(polygonRPC),
	}

	delegatedNetworkServices := networkServices

	var imageGenerators []nft.ImageGenerator

	imageGenerators = append(imageGenerators, dicebear.NewService(httpDicebear.NewClient(c)))
	imageGenerators = append(imageGenerators, nft.DefaultImageGenerator{})

	d := network.NewDelegator(delegatedNetworkServices, imageGenerators)

	container.NetworkDelegator = d

	ipfsClient, err := web3storage.NewClient(cfg.IPFSToken)
	if err != nil {
		log.Printf("could not web3storage.NewClient: %s", err)
		os.Exit(1)
	}

	ipfsService := ipfs.NewService(ipfsClient)

	container.LotteryManager = lottery.NewManager(s, r, d, dr, ipfsService)
}

func GetContainer() *Container {
	return &container
}
