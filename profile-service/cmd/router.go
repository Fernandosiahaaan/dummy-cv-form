package main

import (
	"encoding/json"
	"net/http"
	"profiles-service/internal/handler"

	"github.com/gorilla/mux"
)

func kosongan(w http.ResponseWriter, r *http.Request) {}

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
	router.HandleFunc("/api/photo/{profile_code}", kosongan).Methods(http.MethodGet)    // download
	router.HandleFunc("/api/photo/{profile_code}", kosongan).Methods(http.MethodPut)    // upload
	router.HandleFunc("/api/photo/{profile_code}", kosongan).Methods(http.MethodDelete) // delete

	// WORKING EXPERIENCE
	router.HandleFunc("/api/working-experience/{profile_code}", kosongan).Methods(http.MethodGet) // read
	router.HandleFunc("/api/working-experience/{profile_code}", kosongan).Methods(http.MethodPut) // update

	// EMPLOYMENT
	router.HandleFunc("/api/employment/{profile_code}", kosongan).Methods(http.MethodGet)    // read
	router.HandleFunc("/api/employment/{profile_code}", kosongan).Methods(http.MethodPost)   // create
	router.HandleFunc("/api/employment/{profile_code}", kosongan).Methods(http.MethodDelete) // delete

	// EDUCATION
	router.HandleFunc("/api/education/{profile_code}", kosongan).Methods(http.MethodGet)    // read
	router.HandleFunc("/api/education/{profile_code}", kosongan).Methods(http.MethodPost)   // create
	router.HandleFunc("/api/education/{profile_code}", kosongan).Methods(http.MethodDelete) // delete

	// SKILL
	router.HandleFunc("/api/skill/{profile_code}", kosongan).Methods(http.MethodPost)   // create
	router.HandleFunc("/api/skill/{profile_code}", kosongan).Methods(http.MethodDelete) // delete

	return router
}