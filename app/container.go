package app

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/config"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/ipfs"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/network"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/database/filesystem"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/ethereum"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/ethereum/crypto"
	merkletree "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/hash"
	httpDefault "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/contract"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/rpc"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/web3storage"
	"log"
	"os"
	"time"
)

type Container struct {
	PlateausHTTPClient        *http.Client
	PlateausRPCClient         *rpc.Client
	PlateausValidationService *polygon.PlateausValidationService
	LotteryManager            *lottery.Manager
	NetworkDelegator          *network.Delegator
	IPFSClient                ipfs.Client
}

var container Container

// TODO: decouple this container.go in many other providers separated by context
// TODO: move from app/container.go to app/container/container.go and networks.go too
func init() {
	var dataPath = config.GetAbsolutePath("data/data.json")
	ctx := context.Background()
	cfg := config.GetConfig()
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

	addressLotteryPlateaus := common.HexToAddress(cfg.PlateausLotteryContractAddress)
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

	// Register Networks
	registerNetworks(ctx, cfg)

	// LotteryManager
	ipfsClient, err := web3storage.NewClient(cfg.IPFSToken)
	if err != nil {
		log.Printf("could not web3storage.NewClient: %s", err)
		os.Exit(1)
	}

	ipfsService := ipfs.NewService(ipfsClient)

	container.LotteryManager = lottery.NewManager(s, r, container.NetworkDelegator, dr, ipfsService)
}

func GetContainer() *Container {
	return &container
}
