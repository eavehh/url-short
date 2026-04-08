package res

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, data any, StatusCode int) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(StatusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		println("Error: ", err)
	}
}
