package polygon

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/rpc"
	"log"
)

const networkName = "polygon"

type LotteryValidationService struct {
	rpc rpc.LotteryValidation
}

func NewLotteryValidationService(r rpc.LotteryValidation) *LotteryValidationService {
	return &LotteryValidationService{
		rpc: r,
	}
}

func (s LotteryValidationService) GetNetworkName() string {
	return networkName
}

func (s LotteryValidationService) Supports(network string) bool {
	return s.GetNetworkName() == network
}

func (s LotteryValidationService) MintNFT(hash hash.Hash, lotteryValidation *nft.LotteryValidation, url string, minHeight, maxHeight int) error {
	if err := s.rpc.Mint(hash.String(), lotteryValidation.Encode(), url, int64(minHeight), int64(maxHeight)); err != nil {
		return err
	}

	return nil
}

func (s LotteryValidationService) WasMinted(hash hash.Hash) (bool, error) {
	bal, err := s.rpc.BalanceOf()

	if err != nil {
		return true, err
	}

	tokenURI, err := s.rpc.TokenURI(bal)

	if err != nil {
		return true, err
	}

	metadata, err := nft.NewMetadata(tokenURI)

	if err != nil {
		log.Printf("could not nft.NewMetadata: %s", err)
		return true, err
	}

	return metadata.CheckHash(hash.String()), nil
}

//TODO: this method should not be here
func (s LotteryValidationService) GetAccountBalance() (int64, error) {
	balance, err := s.rpc.AccountBalance()

	if err != nil {
		return 0, err
	}

	return balance, nil
}

// PlateausValidationService implementation of PlateausValidation
type PlateausValidationService struct {
	rpc rpc.PlateausValidation
}

func NewPlateausValidationService(r rpc.PlateausValidation) *PlateausValidationService {
	return &PlateausValidationService{
		rpc: r,
	}
}

func (s PlateausValidationService) GetBalance() (int64, error) {
	bal, err := s.rpc.BalanceOf()

	if err != nil {
		log.Printf("could not rpc.BalanceOf: %s", err)
		return 0, err
	}

	return bal, nil
}
