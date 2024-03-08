package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"apiProducts/products"
)

func main() {

	ps := products.NewProductStorage("products.json")
	err := ps.LoadProducts()
	if err != nil {
		panic(err)
	}
	router := chi.NewRouter()

	router.Get("/ping", HandlePing)
	/*
		router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			var prods string
			for _, p := range ps.Products {
				prods += fmt.Sprint(p) + "\n"
			}
			w.Write([]byte(prods))

		})
	*/
	router.Get("/products", ps.GetProducts)

	//de esta forma el path sería: localhost:8080/products/1 (esto es un path param)
	//router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request)

	//pathParam: PARA OBTENER EL PARAMETRO QUE RECIBE EL ENDPOINT SE UTILIZA chi.URLParam(r, "id")
	//la cual internamente utiliza el context.
	/* 	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
	   		w.Header().Set("Content-Type", "text/plain")
	   		w.WriteHeader(http.StatusOK)
	   		idStr := chi.URLParam(r, "id")
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

	router.Get("/products/{id}", ps.GetProductByKey)
	//de esta forma el path sería: localhost:8080/products?id=1, no se especifica el parámetro que recibe en el endpoint. (esto es un query param)
	//para obtener el query param se utiliza r.URL.Query().Get("id")
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

	/* 	router.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
	   		w.Header().Set("Content-Type", "text/plain")
	   		w.WriteHeader(http.StatusOK)
	   		priceMinStr := r.URL.Query().Get("priceGT")

	   		price, err := strconv.Atoi(priceMinStr)
	   		if err != nil {
	   			w.Write([]byte("Error: " + err.Error()))
	   			return
	   		}

	   		priceFloat := float32(price)
	   		var prods string
	   		for _, p := range ps.Products {
	   			if p.Price > float64(priceFloat) {
	   				prods += fmt.Sprint(p) + "\n"
	   			}
	   		}

	   		if prods == "" {
	   			prods = "No products were found with a price higher than " + priceMinStr
	   		}
	   		w.Write([]byte(prods))
	   	})
	*/

	router.Get("/products/search", ps.GetProductsGreaterThan)

	http.ListenAndServe(":8080", router)
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\nStatus was: " + fmt.Sprint(http.StatusOK)))
}
