package main

import (
	"flag"
	tgClient "github.com/KDias-code/clients/telegram"
	event_consumer "github.com/KDias-code/consumer/event-consumer"
	"github.com/KDias-code/events/telegram"
	"github.com/KDias-code/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started...")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped!", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token which need to access tg bot")

	if *token != "" {
		log.Fatal("token is no specified")
	}

	return *token
}
