package response

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(msg map[string]interface{}, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int                    `json:"status"`
		Message    map[string]interface{} `json:"msg"`
	}

	temp := &errdata{Statuscode: 200, Message: msg}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

func ErrorResponse(error string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &errdata{Statuscode: 400, Message: error}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}

func ServerErrResponse(error string, writer http.ResponseWriter) {
	type servererrdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &servererrdata{Statuscode: 500, Message: error}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(temp)
}

func ValidationResponse(fields map[string][]string, writer http.ResponseWriter) {
	response := make(map[string]interface{})
	response["errors"] = fields
	response["status"] = 422
	response["msg"] = "validation error"

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(writer).Encode(response)
}
