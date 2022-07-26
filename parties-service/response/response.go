package response

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int                    `json:"status"`
		Message    map[string]interface{} `json:"msg"`
	}
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(msg), &jsonMap)

	temp := &errdata{Statuscode: 200, Message: jsonMap}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

func SuccessArrRespond(fields []*interface{}, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)
	type data struct {
		Parties    []*interface{} `json:"data"`
		Statuscode int            `json:"status"`
		Message    string         `json:"msg"`
	}
	temp := &data{Parties: fields, Statuscode: 200, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
	}

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

	//Send header, status code and output to writer
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
