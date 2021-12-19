package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
	//token = flags.Get(token)
	//tgClient = telegram.New(token)
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
