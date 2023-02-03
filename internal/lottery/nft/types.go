package nft

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Metadata struct {
	Name        string `json:"name"`
	ImageData   string `json:"image_data"`
	ExternalUrl string `json:"external_url"`
	Description string `json:"description"`
}

func NewMetadata(tokenURI string) (*Metadata, error) {
	splitedURI := strings.Split(tokenURI, ",")
	decodedURI, err := base64.StdEncoding.DecodeString(splitedURI[1])

	if err != nil {
		return nil, err
	}

	metadata := &Metadata{}
	if err := json.Unmarshal(decodedURI, metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

func (m Metadata) CheckHash(hash string) bool {
	return strings.Contains(m.Description, hash)
}

type LotteryValidation struct {
	Image []byte
}

func NewLotteryValidation(img []byte) *LotteryValidation {
	return &LotteryValidation{
		Image: img,
	}
}

func (i LotteryValidation) Encode() string {
	return fmt.Sprintf("data:image/svg+xml;base64,%s", base64.StdEncoding.EncodeToString(i.Image))
}
