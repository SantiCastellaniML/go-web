package main

import (
	"go-packages/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	//dependencies:
	handler := handlers.NewDefaultTask(nil, 0)

	//create router
	router := chi.NewRouter()
	//associate handlers:
	router.Post("/tasks", handler.Create())

	//start server
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
