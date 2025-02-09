package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func RespondWithError(wrt http.ResponseWriter, err string, statusCode int) error {
	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(statusCode)
	errorResponse := ErrorResponse{Reason: err}
	errJSON := json.NewEncoder(wrt).Encode(errorResponse)
	return errJSON
}

func SendBadReq(wrt http.ResponseWriter) error {
	err := "invalid request format or query parameters"
	errResp := RespondWithError(wrt, err, http.StatusBadRequest)
	return errResp
}
