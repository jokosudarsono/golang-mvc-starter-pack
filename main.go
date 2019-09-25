package main

import (
	"log"

	"todo/core"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := core.App{}
	app.Run(":8000")
}
