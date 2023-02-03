package http

import (
	"fmt"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/http"
	"io"
	"log"
	http2 "net/http"
	"net/url"
)

type Client struct {
	http.ClientDoer
	url string
}

func NewClient(c http.ClientDoer) *Client {
	return &Client{
		ClientDoer: c,
		url:        "http://sentry-nodes.rhizom.me:1317",
	}
}

func (c Client) do(req *http2.Request) ([]byte, error) {
	res, err := c.Do(req)

	if err != nil {
		log.Printf("could not execute /blocks/latests: %s", err)
		return nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("could not io.ReadAll res.Body: %s", err)

		return nil, err
	}

	return b, nil
}

func (c Client) GetLatestBlock() ([]byte, error) {
	log.Println("getting latest block")

	req, err := http2.NewRequest(http2.MethodGet, fmt.Sprintf("%s/blocks/latest", c.url), nil)

	if err != nil {
		log.Printf("could not retrieve GetBlocks: %s", err)
		return nil, err
	}

	return c.do(req)
}

func (c Client) GetTransactions(minHeight, maxHeight, offset, limit int) ([]byte, error) {
	log.Printf("getting transactions on page %d : %d - %d", offset, minHeight, maxHeight)

	url := url.Values{}
	url.Set("events", fmt.Sprintf("tx.height>=%d", minHeight))
	url.Add("events", fmt.Sprintf("tx.height<=%d", maxHeight))
	url.Set("pagination.limit", fmt.Sprintf("%d", limit))
	url.Set("pagination.offset", fmt.Sprintf("%d", offset))
	url.Set("order_by", "ORDER_BY_ASC")

	log.Println(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs?%s", c.url, url.Encode()))

	req, err := http2.NewRequest(http2.MethodGet, fmt.Sprintf("%s/cosmos/tx/v1beta1/txs?%s", c.url, url.Encode()), nil)

	if err != nil {
		log.Printf("could not retrieve GetTransactions [minHeight: %d - maxHeight: %d - offset: %d]: %s", minHeight, maxHeight, offset, err)
		return nil, err
	}

	return c.do(req)
}
