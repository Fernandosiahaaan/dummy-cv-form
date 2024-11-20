package main

import (
	"dummy-cv-form/internal/handler"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func router(handler *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}).Methods(http.MethodGet)

	// PROFILES
	router.HandleFunc("/api/profile", handler.ProfileCreate).Methods(http.MethodPost)               // create
	router.HandleFunc("/api/profile/{profile_code}", handler.ProfileRead).Methods(http.MethodGet)   // read
	router.HandleFunc("/api/profile/{profile_code}", handler.ProfileUpdate).Methods(http.MethodPut) // update

	// PHOTO
	router.HandleFunc("/api/photo/{profile_code}", handler.DownloadPhoto).Methods(http.MethodGet)  // download
	router.HandleFunc("/api/photo/{profile_code}", handler.UploadPhoto).Methods(http.MethodPut)    // upload
	router.HandleFunc("/api/photo/{profile_code}", handler.DeletePhoto).Methods(http.MethodDelete) // delete

	// WORKING EXPERIENCE
	router.HandleFunc("/api/working-experience/{profile_code}", handler.WorkingExperienceRead).Methods(http.MethodGet)    // read
	router.HandleFunc("/api/working-experience/{profile_code}", handler.WorkingExperienceCreate).Methods(http.MethodPost) // create
	router.HandleFunc("/api/working-experience/{profile_code}", handler.WorkingExperienceUpdate).Methods(http.MethodPut)  // create

	// EMPLOYMENT
	router.HandleFunc("/api/employment/{profile_code}", handler.EmploymentsRead).Methods(http.MethodGet)     // read
	router.HandleFunc("/api/employment/{profile_code}", handler.EmploymentCreate).Methods(http.MethodPost)   // create
	router.HandleFunc("/api/employment/{profile_code}", handler.EmploymentDelete).Methods(http.MethodDelete) // delete

	// EDUCATION
	router.HandleFunc("/api/education/{profile_code}", handler.EducationsRead).Methods(http.MethodGet)     // read
	router.HandleFunc("/api/education/{profile_code}", handler.EducationCreate).Methods(http.MethodPost)   // create
	router.HandleFunc("/api/education/{profile_code}", handler.EducationDelete).Methods(http.MethodDelete) // delete

	// SKILL
	router.HandleFunc("/api/skill/{profile_code}", handler.SkillsRead).Methods(http.MethodGet)     // read
	router.HandleFunc("/api/skill/{profile_code}", handler.SkillCreate).Methods(http.MethodPost)   // create
	router.HandleFunc("/api/skill/{profile_code}", handler.SkillDelete).Methods(http.MethodDelete) // delete

	return router
}
