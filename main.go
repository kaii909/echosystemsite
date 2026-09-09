package main

import (
	// "errors"
	// "context"
	// "fmt"
	// "io/fs"
	"log"
	"net/http"

	// "os"
	"www.echosystem/views"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := views.HeaderTemplate("echosystem", "person").Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
  mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", homeHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
