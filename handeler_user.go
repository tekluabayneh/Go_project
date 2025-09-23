package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	database "github.com/tekluabayneh/Go_project/internal/db"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (apiCfg *apiConfig) handelerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:name`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	 err := decoder.Decode(&params)
	 if err != nil{ 
		handeErrorResponse(w, 400, fmt.Sprintf("error parsing json %v\n", err))
		return 
	 }
	NewUser , err := apiCfg.DB.CreateUser(r.Context(),  database.CreateUserParams{
		ID:uuid.New(), 
		CreatedAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
          Name : params.Name,
	 })
	 if err != nil{ 
		handeErrorResponse(w, 400, fmt.Sprintf("couldn't create user %v\n:", err));
		return 
	 }

	responseWithJSON(w, 200, NewUser)

}


func (apiCfg *apiConfig) handelergrGetUser(w http.ResponseWriter, r *http.Request, user database.GetUserByAPIkeyRow) {
	
	responseWithJSON(w, 200, (user))

}
