package handles

import (
	"main/internal/config"

	apiTokens "github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

type Handlers struct {
	Env       config.Env
	Auth      auth.AuthClient
	ApiTokens apiTokens.ApiTokenClient
}
