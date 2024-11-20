package handler

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) SkillsRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	skills, err := h.service.GetSkills(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	} else if skills == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.ResponseBasic{Error: true, Message: fmt.Sprintf("not found skills with profile_code '%d' from db", profileCode)})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: skills})
}

func (h *Handler) SkillCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var skill model.Skill
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
		return
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: "failed parse body request"})
		return
	}
	skill.ProfileCode = profileCode

	respondSkill, err := h.service.CreateSkill(&skill)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	var bodyRespond model.CreateResponseBasic = model.CreateResponseBasic{
		ProfileCode: respondSkill.ProfileCode,
		ID:          respondSkill.ID,
	}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}

func (h *Handler) SkillDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: "failed convert profile code"})
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
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	bodyRespond := model.OnlyProfileCodeResponse{ProfileCode: profileCode}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}
