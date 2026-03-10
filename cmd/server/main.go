package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/echo-logger/service/ent"
	"github.com/echo-logger/service/internal/handler"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/echodb?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Close()

	echoHandler := handler.NewEchoHandler(client)
	http.HandleFunc("/requests", echoHandler.ListRequests)
	http.Handle("/", echoHandler)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
