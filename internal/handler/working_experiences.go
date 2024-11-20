package handler

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) WorkingExperienceRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed to convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	workingExperience, err := h.service.GetWorkingExperiences(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	} else if workingExperience == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.ResponseBasic{Error: true, Message: fmt.Sprintf("not found working experience with profile_code '%d' from db", profileCode)})
		return
	}

	var bodyResponse map[string]any = map[string]any{"workingExperience": workingExperience.Experience}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyResponse})
}

func (h *Handler) WorkingExperienceCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var workingExperience model.WorkingExperience
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed to convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&workingExperience); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "failed to parse body request"})
		return
	}
	workingExperience.ProfileCode = profileCode

	workingExperienceResp, err := h.service.CreateWorkingExperience(profileCode, &workingExperience)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	var bodyRespond model.CreateResponseBasic = model.CreateResponseBasic{
		ProfileCode: profileCode,
		ID:          workingExperienceResp.ID,
	}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}

func (h *Handler) WorkingExperienceUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var workingExperience model.WorkingExperience
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed to convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&workingExperience); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "failed to parse body request"})
		return
	}
	workingExperience.ProfileCode = profileCode

	workingExperienceResp, err := h.service.UpdateWorkingExperience(profileCode, &workingExperience)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}
	var bodyResponse map[string]any = map[string]any{"profileCode": workingExperienceResp.Experience}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyResponse})
}
