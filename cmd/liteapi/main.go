package main

import (
	"context"
	"encoding/base64"
	"github.com/kosrk/ton-http-liteapi/api"
	"github.com/kosrk/ton-http-liteapi/config"
	log "github.com/sirupsen/logrus"
	"github.com/startfellows/tongo/liteclient"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.GetConfig()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	publicKey, err := base64.StdEncoding.DecodeString(config.Config.LiteServerKey)
	if err != nil {
		log.Fatalf("public key decoding error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	c, err := liteclient.NewConnection(ctx, publicKey, config.Config.LiteServer)
	if err != nil {
		log.Fatalf("liteserver connection error: %v", err)
	}
	client := liteclient.NewClient(c)

	apiMux := http.NewServeMux()
	h := api.NewHandler(client)
	api.RegisterHandlers(apiMux, h)
	go func() {
		err := http.ListenAndServe(config.Config.APIHost, apiMux)
		if err != nil {
			log.Fatalf("api error: %v", err)
		}
	}()
	log.Info("Started")
	<-sigChannel
	log.Info("SIGTERM received")
}
