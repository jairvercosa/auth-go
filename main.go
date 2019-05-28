package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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

	if authenticate(requestBody) {
		fmt.Fprint(w, "welcome")
	} else {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

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
