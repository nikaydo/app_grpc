package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/config"
	GRPc "main/internal/grpc"
	s "main/internal/server"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	Env := config.ReadEnv()
	auth := GRPc.RunAuth(fmt.Sprintf("%s:%s", Env.EnvMap["AUTH_HOST"], Env.EnvMap["AUTH_PORT"]))
	apiToken := GRPc.RunApiToken(fmt.Sprintf("%s:%s", Env.EnvMap["APITOKEN_HOST"], Env.EnvMap["APITOKEN_PORT"]))
	server := s.ServerInit(Env, auth, apiToken)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		log.Printf("Starting server on %s:%s", Env.EnvMap["HOST"], Env.EnvMap["PORT"])
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s:%s: %v\n", Env.EnvMap["HOST"], Env.EnvMap["PORT"], err)
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
