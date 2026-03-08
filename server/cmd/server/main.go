package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ARitik/echelon/server/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := api.NewRouter()

	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
