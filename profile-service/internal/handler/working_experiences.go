package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profiles-service/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) WorkingExperienceRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed to convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	workingExperience, err := h.service.GetWorkingExperiences(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	} else if workingExperience == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.Response{Error: true, Message: fmt.Sprintf("not found working experience with profile_code '%d' from db", profileCode)})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Read working experience with profile code %d", profileCode), Data: workingExperience})
}

func (h *Handler) WorkingExperienceCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var workingExperience model.WorkingExperience
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed to convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&workingExperience); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: "failed to parse body request"})
		return
	}
	workingExperience.ProfileCode = profileCode

	createdWorkingExperience, err := h.service.CreateWorkingExperience(profileCode, &workingExperience)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Created working experience with id %d", createdWorkingExperience.ID), Data: createdWorkingExperience})
}

func (h *Handler) WorkingExperienceUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var workingExperience model.WorkingExperience
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed to convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&workingExperience); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: "failed to parse body request"})
		return
	}
	workingExperience.ProfileCode = profileCode

	updatedWorkingExperience, err := h.service.UpdateWorkingExperience(profileCode, &workingExperience)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Updated working experience with id %d", updatedWorkingExperience.ID), Data: updatedWorkingExperience})
}
