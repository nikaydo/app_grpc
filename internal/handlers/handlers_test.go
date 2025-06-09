package handles_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	handles "main/internal/handlers"
	mocks "main/internal/mocks"
	"main/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/nikaydo/grpc-contract/gen/apiToken"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

func TestHandlers_Token_GET(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/token", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "valid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), &auth.ValidateJWTRequest{Token: "valid-jwt-token", Refresh: false}).
		Return(&auth.ValidateJWTResponse{Id: 123}, nil)

	mockApiTokens.EXPECT().
		Create(gomock.Any(), &apiToken.CreateRequest{Id: 123}).
		Return(&apiToken.CreateResponse{Token: "generated-token"}, nil)

	h.Token(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK, got %v", res.StatusCode)
	}
}

func TestHandlers_Token_GET_NoCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/token", nil)
	w := httptest.NewRecorder()

	h.Token(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status Unauthorized, got %v", res.StatusCode)
	}
}

func TestHandlers_Token_GET_ValidateJWTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/token", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "invalid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), &auth.ValidateJWTRequest{Token: "invalid-jwt-token", Refresh: false}).
		Return(nil, errors.New("invalid token"))

	h.Token(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}

func TestHandlers_Token_DELETE(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	tokenToDelete := models.Token{Token: "token-to-delete"}
	bodyBytes, _ := json.Marshal(tokenToDelete)
	req := httptest.NewRequest(http.MethodDelete, "/token", bytes.NewReader(bodyBytes))
	w := httptest.NewRecorder()

	mockApiTokens.EXPECT().
		Delete(gomock.Any(), &apiToken.DeleteRequest{Token: "token-to-delete"}).
		Return(&apiToken.DeleteResponse{}, nil)

	h.Token(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK, got %v", res.StatusCode)
	}
}

func TestHandlers_Token_DELETE_BadJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodDelete, "/token", strings.NewReader("not-json"))
	w := httptest.NewRecorder()

	h.Token(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}

func TestHandlers_Token_DELETE_DeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	tokenToDelete := models.Token{Token: "token-to-delete"}
	bodyBytes, _ := json.Marshal(tokenToDelete)
	req := httptest.NewRequest(http.MethodDelete, "/token", bytes.NewReader(bodyBytes))
	w := httptest.NewRecorder()

	mockApiTokens.EXPECT().
		Delete(gomock.Any(), &apiToken.DeleteRequest{Token: "token-to-delete"}).
		Return(nil, errors.New("delete failed"))

	h.Token(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}

func TestHandlers_Token_UnsupportedMethod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodPost, "/token", nil)
	w := httptest.NewRecorder()

	h.Token(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected status MethodNotAllowed, got %v", res.StatusCode)
	}
}

func TestHandlers_GetTokens_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/tokens", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "valid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), &auth.ValidateJWTRequest{Token: "valid-jwt-token", Refresh: false}).
		Return(&auth.ValidateJWTResponse{Id: 123}, nil)

	mockApiTokens.EXPECT().
		Get(gomock.Any(), &apiToken.GetRequest{Id: 123}).
		Return(&apiToken.GetResponse{Tokens: &apiToken.Tokens{Tokens: []string{"token1", "token2"}}}, nil)

	h.GetTokens(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK, got %v", res.StatusCode)
	}
}

func TestHandlers_GetTokens_NoCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/tokens", nil)
	w := httptest.NewRecorder()

	h.GetTokens(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status Unauthorized, got %v", res.StatusCode)
	}
}

func TestHandlers_GetTokens_ValidateJWTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/tokens", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "invalid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), &auth.ValidateJWTRequest{Token: "invalid-jwt-token", Refresh: false}).
		Return(nil, errors.New("invalid token"))

	h.GetTokens(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}

func TestHandlers_GetTokens_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodGet, "/tokens", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "valid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), &auth.ValidateJWTRequest{Token: "valid-jwt-token", Refresh: false}).
		Return(&auth.ValidateJWTResponse{Id: 123}, nil)

	mockApiTokens.EXPECT().
		Get(gomock.Any(), &apiToken.GetRequest{Id: 123}).
		Return(nil, errors.New("get failed"))

	h.GetTokens(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}
func TestHandlers_VerifyToken_NoCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodPost, "/verify?api-token=some-token", nil)
	w := httptest.NewRecorder()

	h.VerifyToken(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status Unauthorized, got %v", res.StatusCode)
	}
}

func TestHandlers_VerifyToken_ValidateJWTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodPost, "/verify?api-token=some-token", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "invalid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("invalid token"))

	h.VerifyToken(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}

func TestHandlers_VerifyToken_VerifyError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mocks.NewMockAuthClient(ctrl)
	mockApiTokens := mocks.NewMockApiTokenClient(ctrl)

	h := handles.Handlers{
		Auth:      mockAuth,
		ApiTokens: mockApiTokens,
	}

	req := httptest.NewRequest(http.MethodPost, "/verify?api-token=some-token", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "valid-jwt-token"})
	w := httptest.NewRecorder()

	mockAuth.EXPECT().
		ValidateJWT(gomock.Any(), gomock.Any()).
		Return(&auth.ValidateJWTResponse{Id: 123}, nil)

	mockApiTokens.EXPECT().
		Verify(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("verify failed"))

	h.VerifyToken(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status BadRequest, got %v", res.StatusCode)
	}
}
