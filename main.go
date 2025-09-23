package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	database "github.com/tekluabayneh/Go_project/internal/db"
)
type apiConfig struct { 
	DB *database.Queries 
}

func main() {

	godotenv.Load()

	PORT := os.Getenv("PORT")
	DbUrl := os.Getenv("DB_URL")
	if DbUrl == ""{ 
	 log.Fatal("DbUrl is not found in the enviroment ")
	}
	 db, err := sql.Open("postgres", DbUrl)
	 queries := database.New(db)

if err != nil {
    log.Fatal("Database connection failed:", err)
}
defer db.Close()
 

    apiCfg := apiConfig{ 
		DB: queries,
	}

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
	v1Router.Post("/user", apiCfg.handelerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handelergrGetUser))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Server is running on port %s\n", PORT)
	srv.ListenAndServe()
	

}
