package app

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/config"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/network"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/dicebear"
	httpDicebear "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/dicebear/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/ethereum"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/ethereum/crypto"
	httpDefault "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/http"
	externalnetwork "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/contracts"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon"
	contractPolygon "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/contracts"
	rpcPolygon "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/rpc"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/rpc"
	"log"
	"os"
	"sync"
	"time"
)

func registerNetworks(ctx context.Context, cfg *config.Config) {
	wg := sync.WaitGroup{}
	wg.Add(len(cfg.Networks))

	var delegatedNetworkServices []network.Delegated

	for _, n := range cfg.Networks {
		if n.Name == externalnetwork.Polygon {
			registerPlateausNodeValidator(ctx, n.RPC, n.GetPrivateKey(), cfg.PlateausNodeValidatorContractAddress)
		}

		go func(n config.Network) {
			lv := registerNetwork(ctx, n.GetPrivateKey(), n.RPC, n.GetLotteryValidationContractAddress())
			delegatedNetworkServices = append(delegatedNetworkServices, externalnetwork.NewLotteryValidationService(n.Name, lv))
			wg.Done()
		}(n)
	}

	wg.Wait()

	c := &httpDefault.Client{}
	c.Timeout = 10 * time.Second

	var imageGenerators []nft.ImageGenerator

	imageGenerators = append(imageGenerators, dicebear.NewService(httpDicebear.NewClient(c)))
	imageGenerators = append(imageGenerators, nft.DefaultImageGenerator{})

	d := network.NewDelegator(delegatedNetworkServices, imageGenerators)

	container.NetworkDelegator = d
}

func registerPlateausNodeValidator(ctx context.Context, polygonRPC string, polygonPrivateKey string, plateausValidationContractAddress string) {
	clientPolygon, err := ethereum.Dial(polygonRPC)

	if err != nil {
		log.Printf("could not ethereum.Dial on ZK: %s", err)
		os.Exit(1)
	}

	chainIDPolygon, err := clientPolygon.ChainID(ctx)

	if err != nil {
		log.Printf("could not client.ChainID: %s", err)
		os.Exit(1)
	}

	fromAddressPolygon, privateKeyPolygon, err := crypto.AddressFromPrivateKey(polygonPrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	addressPlateausValidationPolygon := common.HexToAddress(plateausValidationContractAddress)
	plateausValidationContract, err := contractPolygon.NewPlateausValidation(addressPlateausValidationPolygon, clientPolygon)

	if err != nil {
		log.Printf("could not contract.NewLotteryValidation: %s", err)
		os.Exit(1)
	}

	var plateausValidationPolygonRPC rpcPolygon.PlateausValidation
	plateausValidationPolygonRPC = rpcPolygon.NewPlateausValidation(clientPolygon, chainIDPolygon, plateausValidationContract, *fromAddressPolygon, privateKeyPolygon)

	container.PlateausValidationService = polygon.NewPlateausValidationService(plateausValidationPolygonRPC)
}

func registerNetwork(
	ctx context.Context,
	networkPrivateKey string,
	networkRPC string,
	lotteryValidationContractAddress string,
) rpc.LotteryValidation {
	clientNetwork, err := ethereum.Dial(networkRPC)

	if err != nil {
		log.Printf("could not ethereum.Dial on ZK: %s", err)
		os.Exit(1)
	}

	chainId, err := clientNetwork.ChainID(ctx)

	if err != nil {
		log.Printf("could not client.ChainID: %s", err)
		os.Exit(1)
	}

	fromAddressNetwork, privateKeyNetwork, err := crypto.AddressFromPrivateKey(networkPrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	addressLotteryValidationNetwork := common.HexToAddress(lotteryValidationContractAddress)
	lotteryValidationContract, err := contracts.NewLotteryValidation(addressLotteryValidationNetwork, clientNetwork)

	if err != nil {
		log.Printf("could not contract.NewLotteryValidation: %s", err)
		os.Exit(1)
	}

	var lotteryValidationNetworkRPC rpc.LotteryValidation
	lotteryValidationNetworkRPC = rpc.NewLotteryValidation(clientNetwork, chainId, lotteryValidationContract, *fromAddressNetwork, privateKeyNetwork)

	return lotteryValidationNetworkRPC
}
