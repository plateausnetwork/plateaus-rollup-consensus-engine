package hash

//go:generate mockgen -source=$GOFILE -destination=./hash_mock.go -package=$GOPACKAGE

import (
	"encoding/hex"
)

type Generator interface {
	GenerateByCollection(elements *[]string) (Hash, error)
	GenerateByMap(elements *map[string]string) (*map[string]string, error)
}

type Hash []byte

func (r Hash) String() string {
	return hex.EncodeToString(r)
}
