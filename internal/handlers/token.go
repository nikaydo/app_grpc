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
		t, err := h.ApiTokens.ApiTokenCreate(context.Background(), &apiTokens.ApiTokenCreateRequest{Id: result.Id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(t.Token)
	case http.MethodDelete:
		var t models.Token
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		_, err := h.ApiTokens.ApiTokenDelete(context.Background(), &apiTokens.ApiTokenDeleteRequest{Token: t.Token})
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
	tok, err := h.ApiTokens.ApiTokenGet(context.Background(), &apiTokens.ApiTokenGetRequest{Id: int32(result.Id)})
	fmt.Println(err)
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	writeJSONResponse(w, tok.Tokens, 200)
}
