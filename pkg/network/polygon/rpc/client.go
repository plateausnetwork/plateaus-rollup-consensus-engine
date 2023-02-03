package rpc

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/contract"
	"log"
	"math/big"
)

type Client struct {
	rpcClient   *ethclient.Client
	chainID     *big.Int
	contract    *contract.LotteryValidation
	fromAddress common.Address
	privateKey  *ecdsa.PrivateKey
}

func New(rpc *ethclient.Client, chainId *big.Int, contract *contract.LotteryValidation, fromAddress common.Address, privateKey *ecdsa.PrivateKey) *Client {
	return &Client{
		rpcClient:   rpc,
		chainID:     chainId,
		contract:    contract,
		fromAddress: fromAddress,
		privateKey:  privateKey,
	}
}

func (c Client) Mint(hash string, img string, url string, minHeight, maxHeight int64) error {
	auth, err := c.createAuth()

	if err != nil {
		return err
	}

	bigMinHeight := big.NewInt(minHeight)
	bigMaxHeight := big.NewInt(maxHeight)

	tx, err := c.contract.Mint(auth, img, hash, url, bigMinHeight, bigMaxHeight)

	if err != nil {
		log.Printf("could not mint peer to network %s: %s", hash, err)
		return err
	}

	log.Printf("peer minted for hash %s: %s", hash, tx.Hash())

	return nil
}

func (c Client) AccountBalance() (int64, error) {
	balance, err := c.rpcClient.BalanceAt(context.TODO(), c.fromAddress, nil)

	if err != nil {
		log.Printf("could not get balance from account: %s", err)
		return 0, nil
	}

	return balance.Int64(), nil
}

func (c Client) BalanceOf() (int64, error) {
	balance, err := c.contract.BalanceOf(nil, c.fromAddress)

	if err != nil {
		log.Printf("could not get balance of: %s", err)

		return 0, err
	}

	return balance.Int64(), nil
}

func (c Client) TokenURI(tokenID int64) (string, error) {
	bigTokenID := big.NewInt(tokenID)
	tokenURI, err := c.contract.TokenURI(nil, bigTokenID)

	if err != nil {
		log.Printf("could not token uri: %s", err)

		return "", err
	}

	return tokenURI, nil
}

//TODO: duplicated from pkg/plateaus/rpc/client.go fix this!
func (c Client) createAuth() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)

	if err != nil {
		log.Printf("could not bind.NewKeyedTransactorWithChainID: %s", err)
		return nil, err
	}

	auth.Value = big.NewInt(0)

	return auth, nil
}
