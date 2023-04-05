package web3storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/web3-storage/go-w3s-client"
	"log"
	"os"
	"path"
)

type Client struct {
	ipfs w3s.Client
}

func NewClient(token string) (*Client, error) {
	var opts []w3s.Option
	opts = append(opts, w3s.WithToken(token))
	ipfsClient, err := w3s.NewClient(opts...)

	if err != nil {
		log.Printf("some error while w3s.NewClient: %s", err)
		return nil, err
	}

	c := &Client{
		ipfs: ipfsClient,
	}

	return c, nil
}

func (c Client) Put(ctx context.Context, name string, content []byte) (string, error) {
	tempDir := fmt.Sprintf("%s/%s", os.TempDir(), "plateaus-consensus")

	if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Printf("could not os.Mkdir: %s", err)
		}
	}

	f, err := os.CreateTemp(tempDir, fmt.Sprintf("*-%s", name))

	if err != nil {
		log.Printf("could not os.CreateTemp: %s", err)
		return "", err
	}

	if _, err := f.Write(content); err != nil {
		log.Printf("could no f.Write: %s", err)
		return "", err
	}

	defer f.Close()

	basename := path.Base(f.Name())
	log.Printf("storing file on IPFS: %s", basename)

	fOpened, err := os.Open(f.Name())

	if err != nil {
		log.Printf("could not os.Open: %s", err)
		return "", err
	}

	cid, err := c.ipfs.Put(ctx, fOpened)

	if err != nil {
		log.Printf("could not upload file ipfs.Put: %s", err)
		return "", err
	}

	fmt.Printf("https://ipfs.io/ipfs/%s/%s\n", cid, basename)
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s/%s", cid, basename)

	return url, nil
}
