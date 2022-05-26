package main

import (
	"avito-bot/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
<<<<<<< HEAD
	fmt.Print("Hello")
	fmt.Print("Hello Leksa")
=======
	tgClient := telegram.New(tgBotHost, mustToken())
	// fetcher = fetcher.New()
	// processor = processor.New()
	// consumer.Start(fetcher, processor)

}
func mustToken() string {
	// bot -tg-bot-token 'my token'
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
>>>>>>> desktop
}
