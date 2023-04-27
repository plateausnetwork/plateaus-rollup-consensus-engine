package network

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/network"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/rpc"
	"log"
)

type LotteryValidationService struct {
	network string
	rpc     rpc.LotteryValidation
}

// Verifying interface Compliance
var _ network.Delegated = (*LotteryValidationService)(nil)

func NewLotteryValidationService(name string, r rpc.LotteryValidation) *LotteryValidationService {
	return &LotteryValidationService{
		network: name,
		rpc:     r,
	}
}

func (s LotteryValidationService) GetNetwork() string {
	return s.network
}

func (s LotteryValidationService) Supports(network string) bool {
	return s.GetNetwork() == network
}

func (s LotteryValidationService) MintNFT(hash hash.Hash, imgURL, url string, minHeight, maxHeight int) error {
	if err := s.rpc.Mint(hash.String(), imgURL, url, int64(minHeight), int64(maxHeight)); err != nil {
		return err
	}

	return nil
}

func (s LotteryValidationService) WasMinted(hash hash.Hash) (bool, error) {
	bal, err := s.rpc.TotalBalance()

	if bal == 0 {
		return false, nil
	}

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

// TODO: this method should not be here
func (s LotteryValidationService) GetAccountBalance() (int64, error) {
	balance, err := s.rpc.AccountBalance()

	if err != nil {
		return 0, err
	}

	return balance, nil
}
