package server

import (
	"main/internal/config"
	GRPc "main/internal/grpc"
	"main/internal/router"
	"net/http"
)

func ServerInit(e config.Env, g GRPc.Service) *http.Server {
	r := router.RouterInit(g.Auth, g.ApiToken, g.Video, e)
	return &http.Server{
		Addr:    e.EnvMap["HOST"] + ":" + e.EnvMap["PORT"],
		Handler: r.Router(),
	}
}
