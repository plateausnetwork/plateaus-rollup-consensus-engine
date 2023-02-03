package crypto

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	cryptoethereum "github.com/ethereum/go-ethereum/crypto"
	"log"
)

func AddressFromPrivateKey(pKey string) (*common.Address, *ecdsa.PrivateKey, error) {
	privateKey, err := cryptoethereum.HexToECDSA(pKey)

	if err != nil {
		log.Printf("could not get address crypto.HexToECDSA: %s", err)
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		err := errors.New(fmt.Sprintf("could not get ecdsa.PublicKey: %t", ok))
		log.Println(err)
		return nil, nil, err
	}

	fromAddress := cryptoethereum.PubkeyToAddress(*publicKeyECDSA)

	return &fromAddress, privateKey, nil
}
