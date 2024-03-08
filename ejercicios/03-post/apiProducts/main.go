package main

import (
	"apiProducts/products"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	ps := products.NewProductStorage("products.json")
	err := ps.LoadProducts()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	//gets:
	router.Get("/ping", HandlePing)
	router.Get("/products", ps.GetProducts)
	router.Get("/products/{id}", ps.GetProductByKey)
	router.Get("/products/search", ps.GetProductsGreaterThan)

	//de esta forma el path sería: localhost:8080/products?id=1, no se especifica el parámetro que recibe en el endpoint. (esto es un query param)
	//para obtener el query param se utiliza r.URL.Query().Get("id")
	/*
		router.Get("/products/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.Write([]byte("Error: " + err.Error()))
				return
			}

			if id <= 0 || id > len(ps.Products) {
				w.Write([]byte("Error: Invalid ID"))
				return
			}

			w.Write([]byte(fmt.Sprint(ps.Products[id-1])))
		})
	*/

	//router.Post("/products", ps.PostProduct)
	//http.HandleFunc("/products", ps.PostProduct)
	router.Post("/hi", ps.Hi)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\nStatus was: " + fmt.Sprint(http.StatusOK)))
}
