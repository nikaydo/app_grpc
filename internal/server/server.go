package server

import (
	"main/internal/config"
	"main/internal/router"
	"net/http"

	apiTokens "github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

func ServerInit(e config.Env, g auth.AuthClient, t apiTokens.ApiTokenClient) *http.Server {
	r := router.RouterInit(g, t)
	return &http.Server{
		Addr:    e.EnvMap["HOST"] + ":" + e.EnvMap["PORT"],
		Handler: r.Router(),
	}
}
