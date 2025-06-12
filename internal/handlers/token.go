package handles

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	"net/http"

	apiTokens "github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

func (h *Handlers) Token(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c, err := r.Cookie("jwt")
		if err != nil {
			log.Println("Error getting cookie:", err)
			writeErrorResponse(w, err, http.StatusUnauthorized)
			return
		}
		result, err := h.Auth.ValidateJWT(context.Background(), &auth.ValidateJWTRequest{Token: c.Value, Refresh: false})
		if err != nil {
			writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		_, err = h.ApiTokens.Create(context.Background(), &apiTokens.CreateRequest{Id: result.Id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		var t models.VideoData
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		_, err := h.ApiTokens.Delete(context.Background(), &apiTokens.DeleteRequest{Token: t.Token})
		if err != nil {
			writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		writeErrorResponse(w, fmt.Errorf("allow methods: Get, Delete"), http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) GetTokens(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("jwt")
	if err != nil {
		log.Println("Error getting cookie:", err)
		writeErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	result, err := h.Auth.ValidateJWT(context.Background(), &auth.ValidateJWTRequest{Token: c.Value, Refresh: false})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	tok, err := h.ApiTokens.Get(context.Background(), &apiTokens.GetRequest{Id: int32(result.Id)})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	var t models.Tokens
	t.Token = tok.Tokens.Tokens
	writeJSONResponse(w, t, 200)
}

func (h *Handlers) VerifyToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("jwt")
	if err != nil {
		log.Println("Error getting cookie:", err)
		writeErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	token := r.FormValue("api-token")
	_, err = h.Auth.ValidateJWT(context.Background(), &auth.ValidateJWTRequest{Token: c.Value, Refresh: false})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	result, err := h.ApiTokens.Verify(context.Background(), &apiTokens.VerifyRequest{Token: token})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	_, err = w.Write(fmt.Appendf(nil, "%t", result.Result))
	if err != nil {
		log.Println("error writing response", err)
	}
}
