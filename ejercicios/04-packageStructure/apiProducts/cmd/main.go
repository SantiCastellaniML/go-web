package main

import (
	"apiProductsStructure/internal/handlers"
	"apiProductsStructure/internal/repository"
	"apiProductsStructure/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//rp := repository.NewProductMap(nil, "products.json")

	rp := repository.NewProductSlice(nil, "products.json")
	sv := service.NewProductService(rp)
	hd := handlers.NewProductHandler(*sv)

	err := rp.LoadProducts()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	//gets:
	router.Get("/ping", HandlePing)
	router.Get("/products", hd.GetProducts)
	router.Get("/products/{id}", hd.GetProductByKey)
	router.Get("/products/search", hd.GetProductsGreaterThanPrice)

	//posts:
	router.Post("/products", hd.PostProduct)

	//puts:
	router.Put("/products/{id}", hd.PutProduct)

	//patches:
	router.Patch("/products/{id}", hd.PatchProduct)

	//deletes:
	router.Delete("/products/{id}", hd.DeleteProduct)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\nStatus was: " + fmt.Sprint(http.StatusOK)))
}
