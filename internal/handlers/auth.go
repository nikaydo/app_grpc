package handles

import (
	"context"
	"encoding/json"
	"fmt"
	"main/internal/models"
	"net/http"
	"time"

	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

func (h *Handlers) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	n, err := h.Auth.SignIn(context.Background(), &auth.SignInRequest{Login: user.Login, Password: user.Pass})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	http.SetCookie(w, MakeCookie("jwt", n.Token, time.Duration(10*time.Minute)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(n.Token))
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}

func (h *Handlers) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	n, err := h.Auth.SignUp(context.Background(), &auth.SignUpRequest{Login: user.Login, Password: user.Pass})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	if n.UserId == 0 {
		writeErrorResponse(w, fmt.Errorf("ошибка"), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
