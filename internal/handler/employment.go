package handler

import (
	"dummy-cv-form/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (h *Handler) EmploymentsRead(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: model.ErrParseProfileCode})
	}
	profileCode := int64(profileCodeInt)

	employments, err := h.service.GetEmployments(profileCode)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.HasPrefix(err.Error(), model.ProfileCodeErr01) {
			statusCode = http.StatusNotFound
		}
		model.CreateResponseHttp(w, r, statusCode, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	var bodyRespond model.OnlyDataResponse = model.OnlyDataResponse{Data: employments}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}

func (h *Handler) EmploymentCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var employment model.Employment
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: model.ErrParseProfileCode})
	}
	profileCode := int64(profileCodeInt)

	if err := json.NewDecoder(r.Body).Decode(&employment); err != nil {
		fmt.Println(err)
		model.CreateResponseHttp(w, r, http.StatusBadRequest, model.ResponseBasic{Error: true, Message: model.ErrParseJson})
		return
	}
	employment.ProfileCode = profileCode

	respondEmp, err := h.service.CreateEmployment(&employment)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.HasPrefix(err.Error(), model.ProfileCodeErr01) {
			statusCode = http.StatusNotFound
		}
		model.CreateResponseHttp(w, r, statusCode, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	var bodyRespond model.CreateResponseBasic = model.CreateResponseBasic{
		ProfileCode: profileCode,
		ID:          respondEmp.ID,
	}
	model.CreateResponseHttp(w, r, http.StatusCreated, model.ResponseBasic{Error: false, Data: bodyRespond})
}

func (h *Handler) EmploymentDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	profileCodeInt, err := strconv.Atoi(vars["profile_code"])
	if err != nil {
		model.CreateResponseHttp(w, r, http.StatusInternalServerError, model.ResponseBasic{Error: true, Message: model.ErrParseProfileCode})
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
		statusCode := http.StatusInternalServerError
		if strings.HasPrefix(err.Error(), model.EmploymentErr01) || strings.HasPrefix(err.Error(), model.ProfileCodeErr01) {
			statusCode = http.StatusNotFound
		}
		model.CreateResponseHttp(w, r, statusCode, model.ResponseBasic{Error: true, Message: err.Error()})
		return
	}

	bodyRespond := model.OnlyProfileCodeResponse{ProfileCode: profileCode}
	model.CreateResponseHttp(w, r, http.StatusOK, model.ResponseBasic{Error: false, Data: bodyRespond})
}
