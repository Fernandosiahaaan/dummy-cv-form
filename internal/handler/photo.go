package handler

import (
	"dummy-cv-form/internal/model"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (h *Handler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var requestData model.BodyUploadRequest

	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "Invalid request body"})
		return
	}

	base64Data := strings.Split(requestData.Base64Img, ",")[1] // Remove prefix `data:image/png;base64,`
	imgData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "Failed to decode base64 image"})
		return
	}
	photoURL, err := h.service.SavePhoto(profileCode, imgData)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	// Respond with the uploaded photo URL
	response := model.BodyUploadResponse{
		ProfileCode: profileCode,
		PhotoURL:    photoURL,
	}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: response})
}

func (h *Handler) DownloadPhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	responseBody, err := h.service.StorePhoto(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: err.Error()})
	}

	w.Header().Set("Content-Type", "image/png")
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: responseBody})
}

func (h *Handler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)
	responseBody, err := h.service.DeletePhoto(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: responseBody})
}