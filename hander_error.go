package main

import "net/http"

func handerError(w http.ResponseWriter, r *http.Request) {
	handeErrorResponse(w, 500, "Something went wrong")
}
