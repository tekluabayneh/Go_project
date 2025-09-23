package main

import (
	"encoding/json"
	"log"
	"net/http"
)

 


func handlerReadyness(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("oops! body is empty")
		return
	}

	defer r.Body.Close()
	responseWithJSON(w, 200, user)

}
