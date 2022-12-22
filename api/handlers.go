package api

import (
	"github.com/startfellows/tongo/liteclient"
)

//go:generate go run generator.go

type Handler struct {
	blockchain *liteclient.Client
}

func NewHandler(client *liteclient.Client) *Handler {
	return &Handler{blockchain: client}
}

