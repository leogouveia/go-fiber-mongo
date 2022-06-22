package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leogouveia/go-fiber-mongo/config"
	"github.com/leogouveia/go-fiber-mongo/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConectaMongo()
	routes.HandleRequests()
}
