package handles

import (
	"main/internal/database"

	apiTokens "github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

type Handlers struct {
	Db        database.Database
	Auth      auth.AuthClient
	ApiTokens apiTokens.ApiTokenClient
}
