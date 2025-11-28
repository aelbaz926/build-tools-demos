package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AppInfo struct {
	Message  string `json:"message"`
	Tool     string `json:"tool"`
	Language string `json:"language"`
	Version  string `json:"version"`
}

type DependencyInfo struct {
	Dependencies map[string]string `json:"dependencies"`
	Description  string            `json:"description"`
	Endpoints    []string          `json:"endpoints"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	info := AppInfo{
		Message:  "Go Modules Demo Application",
		Tool:     "Go Modules",
		Language: "Go",
		Version:  "1.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := DependencyInfo{
		Dependencies: map[string]string{
			"gorilla/mux": "v1.8.1",
		},
		Description: "REST API demonstrating Go modules",
		Endpoints:   []string{"/", "/info", "/health"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/info", infoHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")

	fmt.Println("=== Go Modules Demo Application ===\n")
	fmt.Println("✅ Server running on http://localhost:8080")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  GET http://localhost:8080/")
	fmt.Println("  GET http://localhost:8080/info")
	fmt.Println("  GET http://localhost:8080/health")
	fmt.Println("\nPress Ctrl+C to stop\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}
