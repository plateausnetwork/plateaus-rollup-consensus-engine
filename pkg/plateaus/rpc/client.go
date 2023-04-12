package rpc

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/plateaus/contract"
	"log"
	"math/big"
	"time"
)

type Client struct {
	chainID     *big.Int
	contract    *contract.Lottery
	fromAddress common.Address
	privateKey  *ecdsa.PrivateKey
}

func New(chainId *big.Int, contract *contract.Lottery, fromAddress common.Address, privateKey *ecdsa.PrivateKey) *Client {
	return &Client{
		chainID:     chainId,
		contract:    contract,
		fromAddress: fromAddress,
		privateKey:  privateKey,
	}
}

func (c Client) Subscribe(networks []string) error {
	auth, err := c.createAuth()

	if err != nil {
		return err
	}

	tx, err := c.contract.Subscribe(auth, networks)

	if err != nil {
		log.Printf("could not subscribe peer to network %s: %s", networks, err)
		return err
	}

	log.Printf("peer subscribed on %s: %s", networks, tx.Hash())

	return nil
}

func (c Client) GetLotteryWinners() (string, error) {
	r, err := c.contract.GetCurrentWinners(&bind.CallOpts{
		From: c.fromAddress,
	})

	if err != nil {
		log.Println("could not get current winners")
		return "", err
	}

	var winners = make(map[string]string)

	for _, sub := range r {
		winners[sub.Addr.String()] = sub.Network
	}

	winner, _ := winners[c.fromAddress.String()]

	return winner, nil
}

func (c Client) IsClosed(date time.Time) (bool, error) {
	bigDate := big.NewInt(date.Unix())

	r, err := c.contract.IsClosed(&bind.CallOpts{
		From: c.fromAddress,
	}, bigDate)

	if err != nil || r == false {
		log.Println("current lottery is not closed")
		return false, err
	}

	return r, err
}

func (c Client) IsOpen(date time.Time) (bool, error) {
	bigDate := big.NewInt(date.Unix())

	r, err := c.contract.IsOpen(&bind.CallOpts{
		From: c.fromAddress,
	}, bigDate)

	if err != nil || r == false {
		log.Println("current lottery is not open")
		return false, err
	}

	return r, err
}

func (c Client) WasPicked() (bool, error) {
	r, err := c.contract.WasPicked(&bind.CallOpts{
		From: c.fromAddress,
	})

	if err != nil || r == false {
		log.Println("winner was not picked")
		return false, err
	}

	log.Println("winner was picked")

	return r, err
}

func (c Client) PickWinner() error {
	auth, err := c.createAuth()

	if err != nil {
		return err
	}

	tx, err := c.contract.Winner(auth)

	if err != nil {
		log.Printf("could not contract.Winner: %s", err)
		return err
	}

	log.Printf("winner picked successfully: %s", tx.Hash())

	return err
}

func (c Client) createAuth() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)

	if err != nil {
		log.Printf("could not bind.NewKeyedTransactorWithChainID: %s", err)
		return nil, err
	}

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(0)

	return auth, nil
}
