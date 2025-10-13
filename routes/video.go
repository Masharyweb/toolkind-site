package routes

import (
	"encoding/json"
	"net/http"
)

func VideoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"message": "Video endpoint (stub). FFmpeg integration can be added here later.",
	})
}
