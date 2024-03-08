package repository

import (
	"apiProductsStructure/internal"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// try to implement this. Currently, it's not working because the Marshal is not working with the map[int]internal.Product type.
func NewProductMap(db map[int]internal.Product, filename string) *ProductMap {
	if db == nil {
		db = map[int]internal.Product{}
	}

	return &ProductMap{
		db:       db,
		filename: filename,
	}
}

type ProductMap struct {
	db       map[int]internal.Product
	filename string
}

func (p *ProductMap) Save(product *internal.Product) (err error) {
	err = p.validateCodeValue(product.Code_value)
	if err != nil {
		return
	}

	product.ID = len(p.db) + 1
	p.AddProduct(product)

	return
}

func (ps *ProductMap) validateCodeValue(codeValue string) (err error) {
	for _, p := range ps.db {
		if p.Code_value == codeValue {
			err = errors.New("code_value already exists")
			return
		}
	}
	return
}

func (ps *ProductMap) AddProduct(p *internal.Product) {
	//ps = append(ps.db, p)
	ps.db[p.ID] = *p
}

func (ps *ProductMap) LoadProducts() (err error) {
	file, err := os.Open("../internal/products.json")
	if err != nil {
		return err
	}

	defer file.Close()

	//reading JSON data:
	var jsonData = json.NewDecoder(file)
	err = jsonData.Decode(&ps.db)

	return err
}

func (ps *ProductMap) GetProducts() (data string) {
	for _, p := range ps.db {
		data += fmt.Sprint(p) + "\n"
	}
	return
}

func (ps *ProductMap) GetProductByKey(key int) (product internal.Product, err error) {
	product, ok := ps.db[key]
	if !ok {
		err = errors.New("product not found")
		return
	}
	return
}

func (ps *ProductMap) GetProductsGreaterThanPrice(price float64) (data string) {
	for _, p := range ps.db {
		if p.Price > price {
			data += fmt.Sprint(p) + "\n"
		}
	}
	return
}
