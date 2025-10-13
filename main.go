package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"ngokit/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/convert", routes.ConvertHandler)
	mux.HandleFunc("/api/compress", routes.CompressHandler)
	mux.HandleFunc("/api/pdf", routes.PDFHandler)
	mux.HandleFunc("/api/images", routes.ImagesHandler)
	mux.HandleFunc("/api/video", routes.VideoHandler)
	mux.HandleFunc("/api/speech", routes.SpeechHandler)

	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", spa(fs))

	port := os.Getenv("PORT")
	if port == "" { port = "3000" }
	log.Println("NGOKit server on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func spa(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".") || strings.HasPrefix(r.URL.Path, "/api/") {
			next.ServeHTTP(w, r)
			return
		}
		r.URL.Path = "/index.html"
		next.ServeHTTP(w, r)
	})
}
