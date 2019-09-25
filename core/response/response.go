package response

import (
	"encoding/json"
	"net/http"
)

func Err(w http.ResponseWriter, code int, message string, statusCode ...string) {
	payload := map[string]interface{}{
		"status":  "errors",
		"message": message,
	}

	if len(statusCode) > 0 {
		payload["status_code"] = statusCode[0]
	}

	JSON(w, code, payload)
}

func JSON(w http.ResponseWriter, code int, payload map[string]interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
