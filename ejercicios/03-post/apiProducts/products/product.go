package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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

func (ps *ProductStorage) PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PostProduct"))
	/* dataMap := map[string]any{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	err = json.Unmarshal(body, &dataMap)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	err = validateDataMap(dataMap)
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	err = product.validate()
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	err = ps.validateCodeValue(product.Code_value)
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// - adding the product to the storage
	product.ID = len(ps.Products) + 1
	ps.AddProduct(product)

	// - returning response
	web.ResponseJSON(w, http.StatusCreated, map[string]any{
		"message": "Product created successfully",
		"product": product,
	})
	*/
}

func (ps ProductStorage) Hi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi"))
}

func (ps *ProductStorage) validateCodeValue(codeValue string) (err error) {
	for _, p := range ps.Products {
		if p.Code_value == codeValue {
			err = errors.New("code_value already exists")
			return
		}
	}
	return
}

func (p Product) validate() (err error) {

	if p.Name == "" {
		err = errors.New("name is empty")
		return
	}

	if p.Quantity <= 0 {
		err = errors.New("quantity is less than or equal to zero")
		return
	}

	if p.Code_value == "" {
		err = errors.New("code_value is empty")
		return
	}

	if p.Expiration == "" {
		err = errors.New("expiration is empty")
		return
	}
	if !validateExpiration(p.Expiration) {
		err = errors.New("expiration is invalid")
		return
	}

	if p.Price <= 0 {
		err = errors.New("price is less than or equal to zero")
		return
	}

	return
}

func validateExpiration(expiration string) bool {
	_, err := time.Parse("02/01/2006", expiration)
	return err == nil
}

func validateDataMap(data map[string]any) (err error) {
	if _, ok := data["name"]; !ok {
		err = errors.New("name is required")
		return
	}
	if _, ok := data["quantity"]; !ok {
		err = errors.New("quantity is required")
		return
	}
	if _, ok := data["code_value"]; !ok {
		err = errors.New("code_value is required")
		return
	}
	if _, ok := data["is_published"]; !ok {
		err = errors.New("is_published is required")
		return

	}
	if _, ok := data["expiration"]; !ok {
		err = errors.New("expiration is required")
		return
	}
	if _, ok := data["price"]; !ok {
		err = errors.New("price is required")
		return
	}
	return
}

func (p Product) String() string {
	return fmt.Sprintf("ID: %d\nName: %s\nQuantity: %d\nCode_value: %s\nIs_published: %t\nExpiration: %s\nPrice: %.2f\n", p.ID, p.Name, p.Quantity, p.Code_value, p.Is_published, p.Expiration, p.Price)
}
