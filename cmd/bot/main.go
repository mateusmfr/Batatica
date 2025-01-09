package main

import (
	"fmt"
	"log"
	"os"

	"batatica/internal/bot"
)

func main() {
	botToken, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: DISCORD_BOT_TOKEN")
	}

	err := bot.InitializeBot(botToken)
	if err != nil {
		log.Fatal("Erro ao inicializar bot:", err)
	}

	fmt.Println("Batatica rodando...")
	select {}
}
