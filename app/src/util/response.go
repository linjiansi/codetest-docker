package util

import (
	"encoding/json"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, status int, response any) {
	ReturnResponseWithHeaders(w, status, response, map[string]string{})
}

func ReturnResponseWithHeaders(w http.ResponseWriter, status int, response any, headers map[string]string) {
	w.Header().Set("Content-Type", "application/json")

	for key, value := range headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
