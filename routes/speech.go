package routes

import (
	"encoding/json"
	"net/http"
)

func SpeechHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"message": "Speech-to-text endpoint (stub). Connect to Whisper or other ASR later.",
	})
}
