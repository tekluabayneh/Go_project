package main

import (
	"fmt"
	"net/http"

	"github.com/tekluabayneh/Go_project/internal/auth"
	database "github.com/tekluabayneh/Go_project/internal/db"
)

type authHandler func(http.ResponseWriter, *http.Request,  database.GetUserByAPIkeyRow)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{ 
  return func(w http.ResponseWriter, r *http.Request){ 
		apiKey, err := auth.GetAPiKey(r.Header)

	if err != nil{ 
		handeErrorResponse(w, 403, fmt.Sprintf("auth error: %v\n", err))
		return 
	}

	user, err := apiCfg.DB.GetUserByAPIkey(r.Context(),  apiKey)
	 if err != nil{ 
		handeErrorResponse(w, 400, fmt.Sprintf("couldn't create user %v\n:", err));
		return 
	 }
	 handler(w, r, user)
  }

}