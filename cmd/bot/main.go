package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Error message: %s", err.Error())
		return
	}
	sess.Identify.Intents = discordgo.IntentMessageContent
	// sess.AddHandler(NewMessage)
	sess.Identify.Intents = discordgo.IntentsGuildMessages
	if err = sess.Open(); err != nil {
		log.Fatalf("Error message: %s", err.Error())
		return
	}
	defer sess.Close()
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func initApp() {
	if err := godotenv.Load("../.env"); err != nil {
		println(err)
		slog.Error("Error to load environment variables", "err", err)
	}
}
