package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profiles-service/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) SkillsRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	skills, err := h.service.GetSkills(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	} else if skills == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.Response{Error: true, Message: fmt.Sprintf("not found skills with profile_code '%d' from db", profileCode)})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Read skills with profile code %d", profileCode), Data: skills})
}

func (h *Handler) SkillCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var skill model.Skill
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: "failed parse body request"})
		return
	}
	skill.ProfileCode = profileCode

	respondSkill, err := h.service.CreateSkill(&skill)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Create skill with id %d", respondSkill.ID), Data: respondSkill})
}

func (h *Handler) SkillDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	idSkillStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idSkillStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid skill id: %v", err), http.StatusBadRequest)
		return
	}
	idSkill := int64(id)

	err = h.service.DeleteSkill(idSkill, profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	}

	var data map[string]any = map[string]any{
		"profileCode": profileCode,
	}
	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Success delete skill with id %d", idSkill), Data: data})
}
