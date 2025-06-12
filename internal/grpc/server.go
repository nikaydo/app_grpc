package grpc

import (
	"fmt"
	"log"
	"main/internal/config"

	apiTokens "github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
	"github.com/nikaydo/grpc-contract/gen/video"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Auth     auth.AuthClient
	ApiToken apiTokens.ApiTokenClient
	Video    video.VideoClient
}

func (s *Service) Init(Env config.Env) {
	s.Auth = RunAuth(fmt.Sprintf("%s:%s", Env.EnvMap["AUTH_HOST"], Env.EnvMap["AUTH_PORT"]))
	s.ApiToken = RunApiToken(fmt.Sprintf("%s:%s", Env.EnvMap["APITOKEN_HOST"], Env.EnvMap["APITOKEN_PORT"]))
	s.Video = RunVideo(fmt.Sprintf("%s:%s", Env.EnvMap["VIDEO_HOST"], Env.EnvMap["VIDEO_PORT"]))
}

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

func RunVideo(addr string) video.VideoClient {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(100<<20),
			grpc.MaxCallRecvMsgSize(100<<20),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("gRPC video client started")
	return video.NewVideoClient(conn)
}
