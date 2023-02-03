package polygon

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/rpc"
	"log"
)

const networkName = "polygon"

type Service struct {
	rpc *rpc.Client
}

func NewService(r *rpc.Client) *Service {
	return &Service{
		rpc: r,
	}
}

func (s Service) GetNetworkName() string {
	return networkName
}

func (s Service) Supports(network string) bool {
	return s.GetNetworkName() == network
}

func (s Service) MintNFT(hash hash.Hash, lotteryValidation *nft.LotteryValidation, url string, minHeight, maxHeight int) error {
	if err := s.rpc.Mint(hash.String(), lotteryValidation.Encode(), url, int64(minHeight), int64(maxHeight)); err != nil {
		return err
	}

	return nil
}

func (s Service) WasMinted(hash hash.Hash) (bool, error) {
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

func (s Service) GetAccountBalance() (int64, error) {
	balance, err := s.rpc.AccountBalance()

	if err != nil {
		return 0, err
	}

	return balance, nil
}
