package tools

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main/internal/app/movie/model"
)

func ConvertIdParamsToInt(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func JsonResponse(w http.ResponseWriter, httpStatusCode int, message string, data interface{}) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.Response{
		Message: message,
		Data:    dataJson,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
