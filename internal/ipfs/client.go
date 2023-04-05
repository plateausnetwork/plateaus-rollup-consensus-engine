package ipfs

//go:generate mockgen -source=$GOFILE -destination=./client_mock.go -package=$GOPACKAGE

import (
	"context"
)

type Client interface {
	Put(ctx context.Context, name string, content []byte) (string, error)
}
