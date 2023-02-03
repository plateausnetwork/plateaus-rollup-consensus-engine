package plateaus

import "time"

//go:generate mockgen -source=$GOFILE -destination=./client_mock.go -package=$GOPACKAGE

type HTTPClient interface {
	GetLatestBlock() ([]byte, error)
	GetTransactions(minHeight, maxHeight, offset, limit int) ([]byte, error)
}

type RPCClient interface {
	IsClosed(date time.Time) (bool, error)
	IsOpen(date time.Time) (bool, error)
	WasPicked() (bool, error)
	Subscribe(network string) error
	PickWinner() error
	GetLotteryWinners() (string, error)
}
