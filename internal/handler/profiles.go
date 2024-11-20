package handler

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) ProfileRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	profile, err := h.service.GetProfile(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	} else if profile == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.ResponseBasic{Error: true, Message: fmt.Sprintf("not found profile with profile_code '%d' from db", profileCode)})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: profile})
}

func (h *Handler) ProfileCreate(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	var err error
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		fmt.Println(err)
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "failed parse body request"})
		return
	}

	profileCode, err := h.service.CreateNewProfile(&profile)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}
	var bodyRespond model.OnlyProfileCodeResponse = model.OnlyProfileCodeResponse{ProfileCode: profileCode}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}

func (s *Handler) ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	var err error
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewDecoder(r.Body).Decode(&profile); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "failed parse body request"})
		return
	}
	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
		return
	}
	profile.ProfileCode = int64(profileCodeInt)

	profileCode, err := s.service.UpdateProfile(profile.ProfileCode, &profile)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	var bodyRespond model.OnlyProfileCodeResponse = model.OnlyProfileCodeResponse{ProfileCode: *profileCode}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}
