package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Faild to marsh Json res: %v\n", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)

}

func handeErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Response with 5XX code error:", msg)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, statusCode, ErrorResponse{
		Error: msg})
}
