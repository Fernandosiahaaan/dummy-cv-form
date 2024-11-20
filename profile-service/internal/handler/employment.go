package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profiles-service/internal/model"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) EmploymentsRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	employments, err := h.service.GetEmployment(profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	} else if employments == nil {
		model.CreateResponseHttp(w, r, http.StatusNotFound, model.Response{Error: true, Message: fmt.Sprintf("not found employments with profile_code '%d' from db", profileCode)})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Read employments with profile code %d", profileCode), Data: employments})
}

func (h *Handler) EmploymentCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var employment model.Employment
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&employment); err != nil {
		fmt.Println(err)
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: "failed parse body request"})
		return
	}
	employment.ProfileCode = profileCode

	respondEmp, err := h.service.CreateEmployment(&employment)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	}

	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Create employment with id %d", respondEmp.ID), Data: respondEmp})
}

func (h *Handler) EmploymentDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.Response{Error: true, Message: "failed convert profile code"})
	}
	profileCode := int64(profileCodeInt)

	idEmpStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idEmpStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid employment id: %v", err), http.StatusBadRequest)
		return
	}
	idEmployment := int64(id)

	err = h.service.DeleteEmployment(idEmployment, profileCode)
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.Response{Error: true, Message: err.Error()})
		return
	}

	var data map[string]any = map[string]any{
		"profileCode": profileCode,
	}
	model.CreateResponseHttp(w, r, http.StatusOK, model.Response{Error: false, Message: fmt.Sprintf("Success delete employment with id %d", id), Data: data})
}
