package handles

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

func (h Handlers) CheckJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		result, err := h.Auth.ValidateJWT(context.Background(), &auth.ValidateJWTRequest{Token: c.Value, Refresh: false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !result.Expired {
			next.ServeHTTP(w, r)
			return
		}
		user, err := h.Auth.CheckUser(context.Background(), &auth.CheckUserRequest{Login: result.Login, Password: "", WithPass: false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err = h.Auth.ValidateJWT(context.Background(), &auth.ValidateJWTRequest{Token: user.User.Refresh, Refresh: true}); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := h.Auth.CreateTokens(context.Background(), &auth.CreateTokensRequest{Id: user.User.Id, Login: user.User.Login, Role: "user"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cockie, err := strconv.Atoi(h.Env.EnvMap["COCKIE_TTL"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.SetCookie(w, MakeCookie("jwt", token.JwtToken, time.Duration(cockie*int(time.Minute))))
		next.ServeHTTP(w, r)
	})
}

func (h Handlers) CheckApiToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
	 
		result, err := h.ApiTokens.Verify(context.Background(), &apiToken.VerifyRequest{Token: token})
		if err != nil {
			writeErrorResponse(w, fmt.Errorf("недействительный api токен"), http.StatusBadRequest)
			return
		}
		if !result.Result {
			writeErrorResponse(w, fmt.Errorf("недействительный api токен"), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
