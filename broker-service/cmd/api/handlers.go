package main

import (
	"encoding/json"
	"net/http"
)


type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (app *Config) Broker(writer http.ResponseWriter, request *http.Request) {

	payload := jsonResponse{
		Error: false,
		Message: "Hit the Broker",
	}

	output, _ := json.MarshalIndent(payload, "", "\t")

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(output)
}
