package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pruebaTecnica/src/core/common"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}, message interface{}) {

	response := common.Response{
		Success: true,
		Code:    statusCode,
		Data:    data,
		Message: message,
	}

	if data == nil {
		response.Success = false
		response.Message = nil
		response.Error = message
	}

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, nil, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil, nil)
}
