package ipfs

import (
	"context"
)

type Service struct {
	c Client
}

func NewService(c Client) *Service {
	return &Service{
		c: c,
	}
}

func (s Service) Upload(ctx context.Context, txs *map[string]string) (string, error) {
	url, err := s.c.Put(ctx, txs)

	if err != nil {
		return "", err
	}

	return url, nil
}
