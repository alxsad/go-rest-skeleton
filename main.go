package main

import (
	"github.com/alxsad/go-rest-skeleton/config"
	"github.com/alxsad/go-rest-skeleton/handler"
	"github.com/alxsad/go-rest-skeleton/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main()  {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		addr    = config.GetEnvString("LISTEN_ADDR", ":8888")
		timeout = config.GetEnvInt("HTTP_TIMEOUT", 5)
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/panic", handler.PanicHandler)

	log.Printf("Listening: %s\n", addr)

	handler := middleware.Logging(mux)
	handler = middleware.PanicRecovery(handler)

	server := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: time.Duration(timeout) * time.Second,
		ReadTimeout:  time.Duration(timeout) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
