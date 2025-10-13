package routes

import (
	"encoding/json"
	"net/http"
)

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"message": "Image compression server endpoint (stub). Frontend compresses in-browser for now.",
	})
}
