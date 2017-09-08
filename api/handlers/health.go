package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/loongy/jaguar/actions"
)

// Health returns a health check handler. It always responds with a status of
// 200, and a JSON object `{ "healthy": true }`.
func Health(ctx actions.Context) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"healthy": true,
		})
	}
}
