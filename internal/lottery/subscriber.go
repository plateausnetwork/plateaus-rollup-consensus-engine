package lottery

//go:generate mockgen -source=$GOFILE -destination=./subscriber_mock.go -package=$GOPACKAGE

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/database"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/plateaus"
	"log"
	"time"
)

type Subscriber interface {
	IsAvailable(latestBlock *LatestBlock) (bool, error)
	GetLatestBlock() (*LatestBlock, error)
	Subscribe(height int, peer string, networks []string) error
}

type SubscribeService struct {
	dr   database.DataRepository
	http plateaus.HTTPClient
	rpc  plateaus.RPCClient
	hg   hash.Generator
}

// NewService returns a new LotterySubscriber lottery.SubscribeService
func NewService(http plateaus.HTTPClient, rpc plateaus.RPCClient, dr database.DataRepository, hg hash.Generator) Subscriber {
	return &SubscribeService{
		http: http,
		dr:   dr,
		hg:   hg,
		rpc:  rpc,
	}
}

func (s SubscribeService) Subscribe(height int, peer string, networks []string) error {
	log.Printf("started to subscribe peer: %s on netowkrs %s", peer, networks)

	for _, network := range networks {
		//TODO: can be executed in parallel
		if err := s.rpc.Subscribe(network); err != nil {
			log.Printf("could not subscribe peer %s on network: %s - %s", peer, network, err)

			return err
		}
	}

	err := s.dr.Store(&database.Data{
		LastBlockSubscribed: height,
	})

	log.Printf("last block subscribed: %d", height)

	if err != nil {
		return err
	}

	log.Printf("finished to subscribe peer")

	return nil
}

// IsAvailable is to check if the peer is available to subscribe on Plateaus
func (s SubscribeService) IsAvailable(latestBlock *LatestBlock) (bool, error) {
	if ok, err := s.checkBlock(latestBlock); !ok || err != nil {
		return ok, err
	}

	if ok, err := s.isOpen(); !ok || err != nil {
		return false, err
	}

	// get subscription from peer: c.IsSubscribed

	return true, nil
}

func (s SubscribeService) GetLatestBlock() (*LatestBlock, error) {
	b, err := s.http.GetLatestBlock()

	if err != nil {
		return nil, err
	}

	var latestBlock LatestBlock

	if err := json.Unmarshal(b, &latestBlock); err != nil {
		log.Printf("json.Unmarshal GetBlocks response: %s", err)
		return nil, err
	}

	return &latestBlock, nil
}

func (s SubscribeService) checkBlock(latestBlock *LatestBlock) (bool, error) {
	height := latestBlock.GetHeight()

	if height <= 0 {
		err := errors.New("height could not be zero")
		log.Print(err)
		return false, err
	}

	//TODO: analisar a forma de armazenar o ultimo range de blocos da loteria
	data, err := s.dr.Get()

	if err != nil {
		return false, err
	}

	if height <= data.LastBlockSubscribed {
		log.Println(fmt.Sprintf("height is less than the last subscribed block: %d", height))
		return false, nil
	}

	subscribedBlock := NewSubscribeBlocks(latestBlock.GetHeight(), data.LastBlockSubscribed, database.EachBlock)

	if height < subscribedBlock.GetHeightWaiting() {
		log.Println(fmt.Sprintf("waiting for the next block to subscribe peer last subscribed block: %d current block: %d next block: %d", data.LastBlockSubscribed, height, subscribedBlock.GetHeightWaiting()))
		return false, nil
	}

	log.Printf("block was checked: height: %d - lastblocksubscribed: %d - nextminblock: %d", height, data.LastBlockSubscribed, subscribedBlock.GetCurrentMinHeight())

	return true, nil
}

func (s SubscribeService) isOpen() (bool, error) {
	return s.rpc.IsOpen(time.Now())
}
