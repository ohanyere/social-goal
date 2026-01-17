package main

import "log"

func main() {
	cfg := config{
		addrs: ":8080",
	}

	app := application{
		config: cfg,
	}

	log.Fatal(app.run())
}
