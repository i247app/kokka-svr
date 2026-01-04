package response

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"kokka.com/kokka/internal/shared/constant/status"
)

func WriteJson(w http.ResponseWriter, ctx context.Context, data any, err error, statusCode status.Code) {
	payload := make(map[string]any)

	if data != nil {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Printf("WriteJson: failed to marshal data: %v\n", err)
			return
		}
		var tmp map[string]any
		err = json.Unmarshal(dataBytes, &tmp)
		if err == nil || tmp != nil {
			payload = tmp
		} else {
			payload["result"] = data
		}
	}

	if err != nil {
		payload["error"] = err.Error()
	}

	// Default to not set if not set
	if statusCode != 0 {
		payload["mstatus"] = statusCode
	}

	if (payload["mmessage"] == "Unknown" || payload["message"] == "") && err != nil {
		payload["mmessage"] = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}
