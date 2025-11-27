package main

import (
	"embed"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Sadotib/cc_project/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public/*/*
var publicFS embed.FS

func init() {
	// initializers.LoadEnv()
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("HTTP Server started", "listenAddress", listenAddr)

	r := chi.NewMux()

	// // Serve ./assets as /assets/*
	// fs := http.FileServer(http.Dir("./assets"))
	// r.Handle("assets/*", http.StripPrefix("assets/", fs))

	r.Handle("/public/*", public())

	r.Get("/", handlers.HomeHandler)
	r.Get("/cloud", handlers.CloudHandler)

	log.Fatal(http.ListenAndServe(listenAddr, r))

}

func public() http.Handler {
	fmt.Println("building static files for production")
	return http.FileServerFS(publicFS)
}
