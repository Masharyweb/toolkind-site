package routes

import (
	"encoding/json"
	"net/http"
)

type ConvertReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	URL  string `json:"url"`
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(map[string]any{
			"ok": true,
			"message": "Use POST with JSON: {from,to,url}. This is a stub API ready for future integration.",
		})
		return
	}
	var req ConvertReq
	_ = json.NewDecoder(r.Body).Decode(&req)
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"received": req,
		"note": "Server-side conversion can be plugged here (FFmpeg/ImageMagick/WASM).",
	})
}
