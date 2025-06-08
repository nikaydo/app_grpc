package server

import (
	"main/internal/config"
	"main/internal/database"
	"main/internal/router"
	"net/http"

	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

func ServerInit(e config.Env, db database.Database, g auth.AuthClient) *http.Server {
	r := router.RouterInit(db, g)
	return &http.Server{
		Addr:    e.EnvMap["host"] + ":" + e.EnvMap["PORT"],
		Handler: r.Router(),
	}
}
