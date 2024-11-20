package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseBasic struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateResponseBasic struct {
	ProfileCode int64 `json:"profileCode"`
	ID          int64 `json:"id"`
}

type OnlyProfileCodeResponse struct {
	ProfileCode int64 `json:"profileCode"`
}

func CreateResponseHttp(w http.ResponseWriter, r *http.Request, statusCode int, response ResponseBasic) {
	w.WriteHeader(statusCode)
	if response.Error {
		json.NewEncoder(w).Encode(response)
		fmt.Printf("❌  [%s] uri = '%s'; status code = %d; message = %s\n", r.Method, r.RequestURI, statusCode, response.Message)
		return
	}

	json.NewEncoder(w).Encode(response.Data)
	fmt.Printf("✅  [%s] uri = '%s'; status code = %d; message = %s\n", r.Method, r.RequestURI, statusCode, response.Message)
}
