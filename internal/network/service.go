package network

//go:generate mockgen -source=$GOFILE -destination=./service_mock.go -package=$GOPACKAGE

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
)

type Service interface {
	Register(root hash.Hash) error
}
