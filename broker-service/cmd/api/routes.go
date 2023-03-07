package main

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5/middleware"

	chi "github.com/go-chi/chi/v5"
)

func (app *Config)ApiRoutes() http.Handler {
	
	router := chi.NewRouter()
	corsOptions := corsOptions()

	router.Use(cors.Handler(corsOptions))
	router.Use(middleware.Heartbeat("/ping"))
	

	router.Post("/", app.Broker)

	return router
}



func corsOptions() cors.Options {

	return cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}
}
