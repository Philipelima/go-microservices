package main

import (
	"fmt"
	"log"
	"net/http"
)

const servicePort = "4000"

type Config struct {}

func main() {

	appConfig := Config{}

	log.Printf("Starting Broker service on port %s", servicePort)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", servicePort),
		Handler: appConfig.ApiRoutes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}