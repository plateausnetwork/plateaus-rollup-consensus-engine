package lottery

//go:generate mockgen -source=$GOFILE -destination=./manager_mock.go -package=$GOPACKAGE

import (
	"context"
	"errors"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/config"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/database"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/ipfs"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/network"
	"log"
)

type Manager struct {
	s    Subscriber
	r    Register
	dn   *network.Delegator
	dr   database.DataRepository
	ipfs *ipfs.Service
}

func NewManager(s Subscriber, r Register, dn *network.Delegator, dr database.DataRepository, i *ipfs.Service) *Manager {
	return &Manager{
		s:    s,
		r:    r,
		dn:   dn,
		dr:   dr,
		ipfs: i,
	}
}

func (m Manager) SubscribePeer(peer string, networks []string) error {
	log.Printf("started to process")
	defer log.Printf("finished to process")

	latestBlock, err := m.s.GetLatestBlock()

	if err != nil {
		return err
	}

	height := latestBlock.GetHeight()

	if height <= 0 {
		err := errors.New("height must be greater than 0")
		log.Printf("coud not process Subscribe: %s", err)

		return err
	}

	ok, err := m.s.IsAvailable(latestBlock)

	if !ok || err != nil {
		return err
	}

	if err := m.s.Subscribe(height, peer, networks); err != nil {
		return err
	}

	return nil
}

func (m Manager) RegisterTx(peer string, networks []config.Network) error {
	log.Printf("lottery register txs: %v", networks)

	if ok, err := m.r.IsClosed(); !ok || err != nil {
		return err
	}

	if err := m.r.PickWinner(); err != nil {
		return err
	}

	networkWinner, err := m.r.GetLotteryWinners(peer)

	log.Printf("network winner %s", networkWinner)

	if err != nil {
		return err
	}

	latestBlock, err := m.s.GetLatestBlock()

	if err != nil {
		return err
	}

	data, err := m.dr.Get()

	if err != nil {
		return err
	}

	subscribeBlocks := NewSubscribeBlocks(latestBlock.GetHeight(), data.LastBlockSubscribed, database.EachBlock)

	r, txsMapped, err := m.r.GenerateRoot(subscribeBlocks)

	if err != nil {
		return err
	}

	if wasMinted, err := m.dn.WasMinted(networkWinner, r); err != nil || wasMinted == true {
		return nil
	}

	log.Printf("nft can be minted: %s", r)

	url, err := m.ipfs.Upload(context.TODO(), txsMapped)

	if err != nil {
		return err
	}

	if err := m.dn.MintNFT(networkWinner, r, url, subscribeBlocks.GetCurrentMinHeight(), subscribeBlocks.GetCurrentMaxHeight()); err != nil {
		return err
	}

	return nil
}

func (m Manager) CheckNetworkBalances(networks []config.Network) []string {
	var allowedNetworks []string

	for _, n := range networks {
		//TODO: could be executed in parallel
		balance, err := m.dn.GetAccountBalance(n.Name)

		if balance <= n.MinAmount || err != nil {
			log.Printf("%s balance account is less than min amount desired: current %f - min: %f", n.Name, balance, n.MinAmount)
			continue
		}

		allowedNetworks = append(allowedNetworks, n.Name)
	}

	return allowedNetworks
}
