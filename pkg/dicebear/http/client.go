package http

import (
	"fmt"
	http2 "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/http"
	"io"
	"log"
	"net/http"
)

type Client struct {
	http2.ClientDoer
	Url string
}

const avatarTypeBottts = "bottts"

func NewClient(c http2.ClientDoer) *Client {
	return &Client{
		ClientDoer: c,
		Url:        "https://avatars.dicebear.com/api",
	}
}

func (c Client) do(req *http.Request) ([]byte, error) {
	res, err := c.Do(req)

	if err != nil {
		log.Printf("could not execute /api: %s", err)
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

func (c Client) GetAvatar(hash string) ([]byte, error) {
	log.Println("getting dicebear avatar")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s.svg", c.Url, avatarTypeBottts, hash), nil)

	if err != nil {
		log.Printf("could not retrieve Dicebear avatar: %s", err)
		return nil, err
	}

	return c.do(req)
}
