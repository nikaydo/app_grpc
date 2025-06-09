package grpc

import (
	"log"

	apiTokens "github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunAuth(addr string) auth.AuthClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("grpc auth start")
	return auth.NewAuthClient(conn)
}

func RunApiToken(addr string) apiTokens.ApiTokenClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("grpc apiTokens start")
	return apiTokens.NewApiTokenClient(conn)
}
