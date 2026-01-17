package main

import (
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addrs string
}

func (app application) run() error {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:  app.config.addrs,
		Handler : mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
	}

	return  server.ListenAndServe()
}