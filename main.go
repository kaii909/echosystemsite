package main

import (
	// "errors"
	// "context"
	// "fmt"
	// "io/fs"
	"log"
	"net/http"

	// "os"
	"encoding/json"

	"echosystem/views"
)

func desktopHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := views.DesktopRoot().Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	loadGuestbook()
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", desktopHandler)

	// guestbook API endpoints
	mux.HandleFunc("/api/guestbook", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			json.NewEncoder(w).Encode(getGuestbookEntries())
			return
		}

		if r.Method == "POST" {
			var req struct {
				Name string `json:"name"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid request", 400)
				return
			}
			if req.Name == "" {
				http.Error(w, "name required", 400)
				return
			}
			entry := addGuestbookEntry(req.Name)
			json.NewEncoder(w).Encode(entry)
			return
		}

		http.Error(w, "method not allowed", 405)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
