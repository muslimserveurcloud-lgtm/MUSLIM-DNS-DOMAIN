package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/muslimserveurcloud-lgtm/MUSLIM-DNS-DOMAIN/internal/registry"
)

func main() {
	r := registry.New()

	data, err := os.ReadFile("configs/extensions.json")
	if err != nil {
		log.Fatal(err)
	}

	var config struct {
		Extensions []string `json:"extensions"`
	}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	for _, ext := range config.Extensions {
		if err := r.AddExtension(ext); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"service": "MUSLIM DNS REGISTRY",
			"status":  "online",
			"version": "1.0.0",
		})
	})

	http.HandleFunc("/v1/extensions", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r.ListExtensions())
	})

	log.Println("MUSLIM DNS REGISTRY listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
