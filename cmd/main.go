package main

import (
	"context"
	"log"
	"main/internal/config"
	"main/internal/database"
	GRPc "main/internal/grpc"
	s "main/internal/server"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	Env := config.ReadEnv()
	err := Env.SetTimeout()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.InitDB(Env)
	if err != nil {
		log.Fatal(err)
	}
	auth := GRPc.Run()
	server := s.ServerInit(Env, db, auth)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	<-stop
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server gracefully stopped")
}
