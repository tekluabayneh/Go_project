package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPiKey(headers  http.Header) (string, error){ 
	val := headers.Get("Authorization") 

	if val == ""{ 
		 return "" , errors.New("apiey is mandatory")
		}
		Vals := strings.Split(val, " ")
		if len(Vals) != 2 || Vals[0] != "ApiKey" { 
			return "" , errors.New("malformed headers")
	}

		
 return Vals[1], nil
	
}