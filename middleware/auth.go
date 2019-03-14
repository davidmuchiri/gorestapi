package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"

	"github.com/dinobambino7/gorestapi/users"
	"github.com/dinobambino7/gorestapi/utils"
)

//JwtAuthentication func
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// End points that dont require authentication
		noAuth := []string{"/api/users/register", "/api/users/login"}

		// Current path
		requestPath := r.URL.Path

		// Check if req does not need authentication
		// Serve the request if it doesnt need authentication

		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})

		// Grab the token from header
		tokenHeader := r.Header.Get("Authorization")

		//If Token is missing, returns with error 403 unauthorized
		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Response(w, utils.Message(false, "Missing auth token", nil))
			return
		}

		// The token normally comes in the format `Bearer {token-body}`
		// We check if the retrieved token matches this requirement

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			w.WriteHeader(http.StatusForbidden)
			response = utils.Message(false, "Invalid/Malformed auth token", nil)
			utils.Response(w, response)
			return
		}

		// If token is okay
		// Grab the part that we need
		tokenPart := splitted[1]

		var tk users.Token

		var e = godotenv.Load("config/config.env")

		if e != nil {
			fmt.Println(e)
		}

		token, err := jwt.ParseWithClaims(tokenPart, &tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		// Malformed token
		if err != nil {
			response = utils.Message(false, "Malformed authentication token", nil)
			w.WriteHeader(http.StatusForbidden)
			utils.Response(w, response)
			return
		}

		// If token is invalid
		if !token.Valid {
			response = utils.Message(false, "Token is not valid.", nil)
			w.WriteHeader(http.StatusForbidden)
			utils.Response(w, response)
			return
		}

		// If everything went well
		// Proceed with the request and set the caller to the user retrived from the parsed token

		//useful for monitoring
		fmt.Println("User ", tk.UserID)

		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}
