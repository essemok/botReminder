package main

import (
	"botReminder/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(tgBotHost, mustToken())
	//fetcher = fetcher.New(tgClient)
	//processor = processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"provide your token for access to telegram token",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not provided")
	}

	return *token
}
