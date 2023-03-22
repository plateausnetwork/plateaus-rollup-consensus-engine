package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type Dialer func(url string) (*ethclient.Client, error)

func Dial(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	defer client.Close()

	if err != nil {
		log.Printf("could not ethclient.Dial to url: %s", err)
		return nil, err
	}

	return client, nil
}
