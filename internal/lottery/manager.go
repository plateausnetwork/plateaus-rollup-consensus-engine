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
	"sync"
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

// TODO: refactor that function
func (m Manager) RegisterTx(peer string, networks []string) error {
	log.Printf("lottery register txs: %d", len(networks))

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

	// TODO: can be used buffered channels for it m.s.GetLatestBlock and m.dr.Get
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

	txsURL, err := m.ipfs.UploadConsensusData(context.TODO(), txsMapped)

	if err != nil {
		return err
	}

	imgBytes, err := m.dn.GenerateImageLotteryValidation(r)

	if err != nil {
		return err
	}

	imageURL, err := m.ipfs.UploadImage(context.TODO(), imgBytes.Image)

	if err != nil {
		return err
	}

	if err := m.dn.MintNFT(networkWinner, r, imageURL, txsURL, subscribeBlocks.GetCurrentMinHeight(), subscribeBlocks.GetCurrentMaxHeight()); err != nil {
		return err
	}

	return nil
}

func (m Manager) CheckNetworkBalances(networks []config.Network) []string {
	var allowedNetworks []string
	wg := sync.WaitGroup{}
	wg.Add(len(networks))

	for _, n := range networks {
		go func(n config.Network, allowedNetworks *[]string) {
			balance, err := m.dn.GetAccountBalance(n.Name)

			if balance <= n.MinAmount || err != nil {
				log.Printf("%s balance account is less than min amount desired: current %f - min: %f", n.Name, balance, n.MinAmount)
				wg.Done()
				return
			}

			*allowedNetworks = append(*allowedNetworks, n.Name)

			wg.Done()
		}(n, &allowedNetworks)
	}

	wg.Wait()

	return allowedNetworks
}
