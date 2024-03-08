package handlers

import (
	"apiProductsStructure/internal"
	"apiProductsStructure/internal/service"
	"apiProductsStructure/web"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	sv service.ProductService
}

func NewProductHandler(sv service.ProductService) *ProductHandler {
	return &ProductHandler{
		sv: sv,
	}
}

func (ph *ProductHandler) PostProduct(w http.ResponseWriter, r *http.Request) {

	body, err := ph.preprocessInsert(r)
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	/*
		dataMap := map[string]any{}
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
	*/

	var product internal.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	/* this goes in service
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
	*/

	// - adding the product to the storage
	if err := ph.sv.Save(&product); err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	// this goes in the repository
	//product.ID = len(ps.Products) + 1
	//ps.AddProduct(product)

	// - returning response
	web.ResponseJSON(w, http.StatusCreated, map[string]any{
		"message": "Product created successfully",
		"product": product,
	})
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

func (ps *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	data := ps.sv.GetProducts()
	w.Write([]byte(data))
}

func (ps *ProductHandler) GetProductByKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	prod, err := ps.sv.GetProductByKey(id)
	w.Write([]byte(fmt.Sprint(prod)))
}

func (ps *ProductHandler) GetProductsGreaterThanPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	priceMinStr := r.URL.Query().Get("priceGT")

	price, err := strconv.Atoi(priceMinStr)
	if err != nil {
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	priceFloat := float64(price)
	var prods string

	prods = ps.sv.GetProductsGreaterThanPrice(priceFloat)
	if prods == "" {
		prods = "No products were found with a price higher than " + priceMinStr
	}
	w.Write([]byte(prods))
}

func (ph *ProductHandler) preprocessInsert(r *http.Request) (body []byte, err error) {
	dataMap := map[string]any{}
	body, err = io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &dataMap)
	if err != nil {
		return
	}

	err = validateDataMap(dataMap)
	if err != nil {
		return
	}

	return
}

func (ph *ProductHandler) PutProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ph.preprocessInsert(r)
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": "Invalid id",
		})
		return
	}

	var product internal.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	// - adding the product to the storage
	if err := ph.sv.UpdateProduct(&product, idInt); err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Could not save the product",
		})
		return
	}

	// - returning response
	web.ResponseJSON(w, http.StatusCreated, map[string]any{
		"message": "Product created successfully",
		"product": product,
	})
}

func (ph *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	dataMap := map[string]any{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
	}

	//tendría que validar qué claves voy a modificar o pasarle el mapa de datos al service.

	err = json.Unmarshal(body, &dataMap)
	if err != nil {
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": "Invalid id",
		})
		return
	}

	var product internal.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}

	//reviso que los campos que hayan llegado en el body sean campos de internal.Product

	// - adding the product to the storage
	//if err := ph.sv.PatchProduct(&product, idInt); err != nil {
	if err := ph.sv.PatchProduct(dataMap, idInt); err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Could not save the product",
		})
		return
	}

	// - returning response
	web.ResponseJSON(w, http.StatusCreated, map[string]any{
		"message": "Product created successfully",
		"product": product,
	})
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"error": "Invalid id",
		})
		return
	}

	if err := ph.sv.DeleteProduct(id); err != nil {
		web.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Could not delete the product",
		})
		return
	}
}
