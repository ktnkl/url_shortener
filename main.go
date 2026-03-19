package main

import (
	"url-shortener/database"
	"url-shortener/server"

	"github.com/ktnkl/dotenv_validator/pkg/dotenv_validator"
)

func main() {
	dotenv_validator.ValidateEnv([]string{"DSN"}, ".env")

	database.Connect()

	server.StartServer()
}
