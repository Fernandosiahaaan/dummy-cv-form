package handler

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) EducationsRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	educations, err := h.service.GetEducation(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	} else if educations == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.ResponseBasic{Error: true, Message: fmt.Sprintf("not found educations with profile_code '%d' from db", profileCode)})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: educations})
}

func (h *Handler) EducationCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var education model.Education
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&education); err != nil {
		fmt.Println(err)
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "failed parse body request"})
		return
	}
	education.ProfileCode = profileCode

	respondEducation, err := h.service.CreateEducation(&education)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	var bodyRespond model.CreateResponseBasic = model.CreateResponseBasic{
		ProfileCode: respondEducation.ProfileCode,
		ID:          respondEducation.ID,
	}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}

func (h *Handler) EducationDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	idEmpStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idEmpStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid education id: %v", err), http.StatusBadRequest)
		return
	}
	idEducations := int64(id)

	err = h.service.DeleteEducation(idEducations, profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	bodyRespond := model.OnlyProfileCodeResponse{ProfileCode: profileCode}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}
