package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, struct {
		Status      string `json:"status"`
		ProjectName string `json:"project_name"`
	}{"available", "Library API"})
}

func info(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, struct {
		Version    string `json:"version"`
		ServerTime string `json:"server_time"`
	}{"1.0.0", time.Now().Format(time.RFC3339)})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("This page does not exist"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", health)
	mux.HandleFunc("/api/v1/info", info)
	mux.HandleFunc("/notfound", notFound)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}