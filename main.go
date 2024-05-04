package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/v3ronez/zilean/cmd/bot"
)

func main() {
	initEnv()
	bot.Run()
}

func initEnv() {
	path, err := os.Getwd()
	if err != nil {
		slog.Error("Error to load environment variables", "err", err)
		return
	}
	if err := godotenv.Load(path + "/.env"); err != nil {
		slog.Error("Error to load environment variables", "err", err)
		return
	}
}
