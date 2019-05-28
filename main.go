package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
)

type AuthenticationPayload struct {
	Username string
	Password string
}

func authenticate(payload AuthenticationPayload) bool {
	return payload.Username == "jairvercosa" && payload.Password == "xyz"
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestBody := AuthenticationPayload{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	if !authenticate(requestBody) {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	token, err := issueJWT(requestBody.Username)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, token)
}

func issueJWT(username string) (string, error) {
	now := time.Now()
	minutesToAdd := time.Minute * time.Duration(5)
	exp := now.Add(minutesToAdd)

	privateKeyFile, err := os.Open("jwt-key")
	if err != nil {
		glog.Warningf("Private key not found.")
		return "", err
	}

	pemfileinfo, _ := privateKeyFile.Stat()

	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	privateKeyFile.Close()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      exp.Unix(),
	})

	tokenString, err := token.SignedString(pembytes)
	if err != nil {
		glog.Warningf("Unable to sign the JWT")
	}

	return tokenString, err
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	addr := []string{host, port}

	mux := http.NewServeMux()
	mux.HandleFunc("/authenticate", authHandler)

	server := &http.Server{
		Addr:              strings.Join(addr, ":"),
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      300 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	glog.Fatal(server.ListenAndServe())
}
