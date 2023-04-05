package ipfs

import (
	"context"
	"encoding/json"
	"log"
)

const (
	ConsensusFile = "plateaus-consensus"
	ImgFile       = "plateaus-consensus-image"
)

type Service struct {
	c Client
}

func NewService(c Client) *Service {
	return &Service{
		c: c,
	}
}

func (s Service) UploadConsensusData(ctx context.Context, txs *map[string]string) (string, error) {
	txsBytes, err := json.Marshal(txs)

	if err != nil {
		log.Printf("could not json.Marshal: %s", err)
		return "", err
	}

	url, err := s.c.Put(ctx, ConsensusFile, txsBytes)

	if err != nil {
		return "", err
	}

	return url, nil
}

func (s Service) UploadImage(ctx context.Context, content []byte) (string, error) {
	url, err := s.c.Put(ctx, ImgFile, content)

	if err != nil {
		return "", err
	}

	return url, nil
}
