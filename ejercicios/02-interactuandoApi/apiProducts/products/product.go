package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ProductStorage struct {
	filename string
	Products []Product
}

func NewProductStorage(filename string) ProductStorage {
	return ProductStorage{
		filename: filename,
		Products: []Product{},
	}
}

func (ps *ProductStorage) AddProduct(p Product) {
	ps.Products = append(ps.Products, p)
}

func (ps *ProductStorage) LoadProducts() (err error) {
	file, err := os.Open("./products/" + ps.filename)
	if err != nil {
		return err
	}

	defer file.Close()

	//reading JSON data:
	var jsonData = json.NewDecoder(file)
	err = jsonData.Decode(&ps.Products)

	return err
}

func (ps *ProductStorage) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	var prods string
	for _, p := range ps.Products {
		prods += fmt.Sprint(p) + "\n"
	}
	w.Write([]byte(prods))

}

func (ps *ProductStorage) GetProductByKey(w http.ResponseWriter, r *http.Request) {
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
}

func (ps *ProductStorage) GetProductsGreaterThan(w http.ResponseWriter, r *http.Request) {
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
}

func (p Product) String() string {
	return fmt.Sprintf("ID: %d\nName: %s\nQuantity: %d\nCode_value: %s\nIs_published: %t\nExpiration: %s\nPrice: %.2f\n", p.ID, p.Name, p.Quantity, p.Code_value, p.Is_published, p.Expiration, p.Price)
}
