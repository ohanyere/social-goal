package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("server_port")
	cfg := config{
		addrs: port,
	}

	app := application{
		config: cfg,
	}

	log.Fatal(app.run())
}
