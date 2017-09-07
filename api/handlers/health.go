package handlers

import (
	"encoding/json"
	"net/http"
)

// Health returns a health check handler. It always responds with a status of
// 200, and a JSON object `{ "healthy": true }`.
func Health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"healthy": true,
		}); err != nil {
			w.WriteHeader(500)
		}
	})
}
