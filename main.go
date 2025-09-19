package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	godotenv.Load()

	PORT := os.Getenv("PORT")
	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	router.Mount("/v1", v1Router)
	v1Router.Get("/ready", handlerReadyness)
	v1Router.Get("/error", handerError)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Server is runnig on port %s\n", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
