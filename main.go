package main

import (
	"flag"
	"github.com/KDias-code/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token which need to access tg bot")

	if *token != "" {
		log.Fatal("token is no specified")
	}

	return *token
}
