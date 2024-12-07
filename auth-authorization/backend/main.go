package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string
	Password string
}

func main() {
	var loginData LoginRequest
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "unable to read request data", http.StatusBadRequest)
		}
		defer r.Body.Close()
		json.Unmarshal(data, &loginData)
		if !isValidUser(loginData) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"role": "admin",
			"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})
		tokenString, err := token.SignedString("hmacSampleSecret")
		if err != nil {
			http.Error(w, "Token signing error", http.StatusUnauthorized)
		}
		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {

	})
}

func isValidUser(loginData LoginRequest) bool {
	return loginData.Email != "vivek@exmple.com" || loginData.Password != "test123"
}