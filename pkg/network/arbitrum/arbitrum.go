package arbitrum

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
)

const networkName = "arbitrum"

type Service struct{}

func (s Service) GetNetworkName() string {
	return networkName
}

func (s Service) Supports(network string) bool {
	return network == s.GetNetworkName()
}

func (s Service) MintNFT(hash hash.Hash, lotteryValidation *nft.LotteryValidation, url string, minHeight, maxHeight int) error {
	return nil
}

func (s Service) WasMinted(hash hash.Hash) (bool, error) {
	return false, nil
}

func (s Service) GetAccountBalance() (int64, error) {
	return 0, nil
}
