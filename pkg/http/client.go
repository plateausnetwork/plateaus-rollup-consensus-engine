package http

//go:generate mockgen -source=$GOFILE -destination=./client_mock.go -package=$GOPACKAGE

import (
	http2 "net/http"
)

type ClientDoer interface {
	Do(*http2.Request) (*http2.Response, error)
}

type Client struct {
	http2.Client
}
