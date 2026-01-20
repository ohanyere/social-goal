package main

import (
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
  	"github.com/go-chi/chi/v5/middleware"

)

type application struct {
	config config
}

type config struct {
	addrs string
}

func (app *application) mount() *chi.Mux{
	  r := chi.NewRouter()
  // A good base middleware stack
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
		
	// Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  	r.Use(middleware.Timeout(60 * time.Second))
	
    r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.HealthCheck)
	})
  	
	return  r
}

func (app *application) run() error {
	mux := app.mount()
	server := &http.Server{
		Addr:  app.config.addrs,
		Handler : mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	return  server.ListenAndServe()
}