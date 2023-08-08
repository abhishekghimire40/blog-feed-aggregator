package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// routes staring with /v1
	versionRouter := chi.NewRouter()
	versionRouter.Get("/readiness", handleReadiness())
	versionRouter.Get("/err", handleError())
	router.Mount("/v1", versionRouter)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Println("serving on port:", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error starting server")
	}
}
