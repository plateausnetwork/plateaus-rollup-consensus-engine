package rpc

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/contracts"
	"log"
	"math/big"
)

type PlateausValidation interface {
	BalanceOf() (int64, error)
}

type PlateausValidationClient struct {
	rpcClient   *ethclient.Client
	chainID     *big.Int
	contract    *contracts.PlateausValidation
	fromAddress common.Address
	privateKey  *ecdsa.PrivateKey
}

func NewPlateausValidation(rpc *ethclient.Client, chainId *big.Int, contract *contracts.PlateausValidation, fromAddress common.Address, privateKey *ecdsa.PrivateKey) PlateausValidation {
	return &PlateausValidationClient{
		rpcClient:   rpc,
		chainID:     chainId,
		contract:    contract,
		fromAddress: fromAddress,
		privateKey:  privateKey,
	}
}

func (c PlateausValidationClient) BalanceOf() (int64, error) {
	balance, err := c.contract.BalanceOf(nil, c.fromAddress)

	if err != nil {
		log.Printf("could not get balance of: %s", err)

		return 0, err
	}

	return balance.Int64(), nil
}
