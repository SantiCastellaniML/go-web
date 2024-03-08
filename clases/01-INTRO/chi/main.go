package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//handlers:
	handlerBase := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}

	handlerHealth := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

	router := chi.NewRouter()
	router.Get("/", handlerBase)
	router.Get("/health", handlerHealth)
	http.ListenAndServe(":8080", nil)
}
