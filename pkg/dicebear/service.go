package dicebear

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/dicebear/http"
)

type Service struct {
	c *http.Client
}

func NewService(c *http.Client) *Service {
	// TODO: high coupling with http.Client
	return &Service{
		c: c,
	}
}

func (s Service) Generate(hash string) (*nft.LotteryValidation, error) {
	bytes, err := s.c.GetAvatar(hash)

	if err != nil {
		return nil, err
	}

	return &nft.LotteryValidation{
		Image: bytes,
	}, nil
}
