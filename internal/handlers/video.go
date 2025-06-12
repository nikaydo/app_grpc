package handles

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"main/internal/models"
	"net/http"
	"time"

	"github.com/nikaydo/grpc-contract/gen/apiToken"
	"github.com/nikaydo/grpc-contract/gen/video"
)

func (h *Handlers) Video(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		VideoPost(w, r, h)
	case http.MethodDelete:
		VideoDelete(w, r, h)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) SearchVideo(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	token := r.FormValue("token")
	v, err := h.Vid.Get(context.Background(), &video.GetRequest{Token: token, VideoName: name})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	vids := models.VideoList{Video: v.Video.Video}
	w.WriteHeader(http.StatusOK)
	writeJSONResponse(w, vids, http.StatusOK)
}

func VideoPost(w http.ResponseWriter, r *http.Request, h *Handlers) {
	name := r.FormValue("name")
	token := r.FormValue("token")
	file, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusBadRequest)
		return
	}
	videoBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	result, err := h.ApiTokens.Verify(context.Background(), &apiToken.VerifyRequest{Token: token})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	if !result.Result {
		writeErrorResponse(w, fmt.Errorf("bad token"), http.StatusBadRequest)
		return
	}
	_, err = h.Vid.Add(context.Background(), &video.AddRequest{Token: token, Video: videoBytes, Name: name})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func VideoDelete(w http.ResponseWriter, r *http.Request, h *Handlers) {
	uuid := r.FormValue("uuid")
	token := r.FormValue("token")

	result, err := h.ApiTokens.Verify(context.Background(), &apiToken.VerifyRequest{Token: token})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	if !result.Result {
		writeErrorResponse(w, fmt.Errorf("bad token"), http.StatusBadRequest)
		return
	}
	_, err = h.Vid.Delete(context.Background(), &video.DeleteRequest{Uuid: uuid})
	if err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Stream(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
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
	vid, err := h.Vid.Stream(context.Background(), &video.StreamRequest{Uuid: uuid})
	if err != nil {

		writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Disposition", "inline")
	http.ServeContent(w, r, "video.mp4", time.Now(), bytes.NewReader(vid.Video))
}
