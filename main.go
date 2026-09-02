package main

import (
	// "errors"
	// "context"
	// "fmt"
	"log"
	"net/http"
	// "os"
	"www.echosystem/views"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := views.HeaderTemplate("person").Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func main() {
	log.Println("Starting server on :8080")

	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
