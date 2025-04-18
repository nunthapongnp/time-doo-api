package main

import (
	"log"
	"net/http"
	"time-doo-api/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal(err)
	}
	cfg, router := app.InitializeApp()

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: router,
	}

	log.Printf("Server running at http://localhost:%s", cfg.App.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to run server: %v", err)
	}
}
