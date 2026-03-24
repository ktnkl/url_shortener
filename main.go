package main

import (
	"url-shortener/bot"
	"url-shortener/database"
	"url-shortener/server"

	_ "github.com/ktnkl/dotenv_validator/pkg/dotenv_validator"
)

func main() {
	// dotenv_validator.ValidateEnv([]string{"DSN", "TG_BOT_API_KEY", "HOST_NAME"}, ".env")

	database.Connect()

	go bot.InitTgBot()

	server.StartServer()
}
