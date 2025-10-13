package routes

import (
	"encoding/json"
	"net/http"
)

func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"message": "Images endpoint (stub). Place for background removal AI or resizing on server.",
	})
}
