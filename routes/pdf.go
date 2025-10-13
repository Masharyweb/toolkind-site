package routes

import (
	"encoding/json"
	"net/http"
)

func PDFHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"message": "PDF server endpoint (stub). Use for large PDFs or background jobs later.",
	})
}
