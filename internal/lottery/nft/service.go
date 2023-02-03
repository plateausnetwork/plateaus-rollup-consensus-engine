package nft

import (
	"fmt"
)

//go:generate mockgen -source=$GOFILE -destination=./service_mock.go -package=$GOPACKAGE

type ImageGenerator interface {
	Generate(hash string) (*LotteryValidation, error)
}

type DefaultImageGenerator struct{}

func (d DefaultImageGenerator) Generate(hash string) (*LotteryValidation, error) {
	return &LotteryValidation{
		Image: d.generateLocalSVG(hash),
	}, nil
}

func (d DefaultImageGenerator) generateLocalSVG(hash string) []byte {
	return []byte(fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" preserveAspectRatio=\"xMinYMin meet\" viewBox=\"0 0 350 350\"><style>.base { fill: white; font-family: serif; font-size: 14px; }</style><rect width=\"100%%\" height=\"100%%\" fill=\"black\" /><text x=\"50%%\" y=\"50%%\" class=\"base\" dominant-baseline=\"middle\" text-anchor=\"middle\">%s</text></svg>", hash))
}
