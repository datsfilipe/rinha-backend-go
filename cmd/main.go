package main

import (
	"log"
	"net/http"
	"time"

	"github.com/datsfilipe/rinha-backend-go/pkg/api"
	"github.com/datsfilipe/rinha-backend-go/pkg/database"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Printf("Error opening database: %v", err)
	}

	router := api.SetupRouter(db)

	server := &http.Server{
		Addr:           ":80",
		Handler:        router,
		ReadTimeout:    1200 * time.Millisecond,
		WriteTimeout:   1200 * time.Millisecond,
	}
	
	server.ListenAndServe()
}
