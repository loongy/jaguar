package handlers

import "net/http"

// Health returns a health check handler. It always responds with a status of
// 201, and an empty JSON object.
func Health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"alive\":true}"))
	})
}
