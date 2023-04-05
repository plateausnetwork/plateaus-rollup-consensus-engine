package network

//go:generate mockgen -source=$GOFILE -destination=./delegator_mock.go -package=$GOPACKAGE

import (
	"errors"
	"fmt"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"log"
	"math"
	"math/big"
	"reflect"
)

type Delegated interface {
	GetAccountBalance() (int64, error)
	GetNetwork() string
	MintNFT(hash hash.Hash, imgURL, url string, minHeight, maxHeight int) error
	WasMinted(hash hash.Hash) (bool, error)
	Supports(network string) bool
}

type Delegator struct {
	networkServices []Delegated
	imageGenerators []nft.ImageGenerator
}

func NewDelegator(services []Delegated, generators []nft.ImageGenerator) *Delegator {
	return &Delegator{
		networkServices: services,
		imageGenerators: generators,
	}
}

func (d Delegator) GetAccountBalance(network string) (float64, error) {
	for _, s := range d.networkServices {
		if s.Supports(network) == false {
			//log.Printf("could not get balance on %s because network is not supported by %s", network, reflect.TypeOf(s))
			continue
		}

		balance, err := s.GetAccountBalance()

		if err != nil {
			return 0, err
		}

		bigBalance := new(big.Float)
		bigBalance.SetInt64(balance)
		balanceValue := new(big.Float).Quo(bigBalance, big.NewFloat(math.Pow10(18)))
		value, _ := balanceValue.Float64()

		return value, nil
	}

	return 0, nil
}

func (d Delegator) MintNFT(network string, hash hash.Hash, imgURL, url string, minHeight, maxHeight int) error {
	for _, s := range d.networkServices {
		if s.Supports(network) == false {
			log.Printf("could not register NFT on %s because network is not supported by %s", network, s.GetNetwork())
			continue
		}

		if err := s.MintNFT(hash, imgURL, url, minHeight, maxHeight); err != nil {
			log.Printf("MintNFT %s: %s", network, err)
			return err
		}

		log.Printf("NFT Minted %s on %s : minheight: %d - maxheight: %d", hash, network, minHeight, maxHeight)

		return nil
	}

	return nil
}

func (d Delegator) GenerateImageLotteryValidation(h hash.Hash) (*nft.LotteryValidation, error) {
	// TODO: implement a way to treat error and generate usgin default image generator (DefaultImageGenerator)
	for _, generator := range d.imageGenerators {
		img, err := generator.Generate(h.String())

		if err != nil {
			return nil, err
		}

		return img, nil
	}

	return nil, errors.New(fmt.Sprintf("could not generate LotteryValidation for hash %s", h))
}

func (d Delegator) WasMinted(network string, hash hash.Hash) (bool, error) {
	for _, s := range d.networkServices {
		if s.Supports(network) == false {
			log.Printf("could not check is NFT was minted on %s because network is not supported by %s", network, reflect.TypeOf(s))
			continue
		}

		wasMinted, err := s.WasMinted(hash)

		if err != nil {
			return true, err
		}

		if wasMinted {
			log.Printf("already minted on %s: %s", network, hash)

			return true, err
		}

		return false, nil
	}

	return false, nil
}
