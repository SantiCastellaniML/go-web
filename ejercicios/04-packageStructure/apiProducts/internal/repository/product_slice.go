package repository

import (
	"apiProductsStructure/internal"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func NewProductSlice(db []internal.Product, filename string) *ProductSlice {
	if db == nil {
		db = []internal.Product{}
	}

	return &ProductSlice{
		db:       db,
		filename: filename,
		//i should have a lastID for when i delete an element, i don't repeat ids for the new ones.
		//currently, it's calculated with the length of the slice, but if i delete an element, the length will be smaller than the lastID.
	}
}

type ProductSlice struct {
	db       []internal.Product
	filename string
}

func (ps *ProductSlice) Save(product *internal.Product) (err error) {
	err = ps.validateCodeValue(product.Code_value)
	if err != nil {
		return
	}

	product.ID = len(ps.db)
	ps.AddProduct(product)

	return
}

func (ps *ProductSlice) UpdateProduct(product *internal.Product, id int) (err error) {
	err = ps.validateCodeValue(product.Code_value)
	if err != nil {
		return
	}

	if id <= 0 || id > len(ps.db) {
		err = errors.New("id nonexistent")
		return
	}

	if id == len(ps.db) { //if it didn't exist, add it
		ps.AddProduct(product)
		return
	}

	ps.db[id-1] = *product
	return
}

func (ps *ProductSlice) validateCodeValue(codeValue string) (err error) {
	for _, p := range ps.db {
		if p.Code_value == codeValue {
			err = errors.New("code_value already exists")
			return
		}
	}
	return
}

func (ps *ProductSlice) AddProduct(p *internal.Product) {
	ps.db = append(ps.db, *p)
}

func (ps *ProductSlice) LoadProducts() (err error) {
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

func (ps *ProductSlice) GetProducts() (data string) {
	for _, p := range ps.db {
		data += fmt.Sprint(p) + "\n"
	}
	return
}

func (ps *ProductSlice) GetProductByKey(key int) (product internal.Product, err error) {
	for _, p := range ps.db {
		if p.ID == key {
			product = p
			return
		}
	}

	err = errors.New("product not found")
	return
}

func (ps *ProductSlice) GetProductsGreaterThanPrice(price float64) (data string) {
	for _, p := range ps.db {
		if p.Price > price {
			data += fmt.Sprint(p) + "\n"
		}
	}
	return
}

func (ps *ProductSlice) PatchProduct(dataMap map[string]any, id int) (err error) {
	if id <= 0 || id > len(ps.db) {
		err = errors.New("id nonexistent")
		return
	}

	//debería revisar las claves del dataMap y actualizar los campos que me interesan.

	//debería validar que el codeValue no existiera o si existe, que sea la misma clave
	p := &ps.db[id-1]

	if p.Name != "" {
		p.Name = p.Name
	}

	/*
		Quantity     int     `json:"quantity"`
		Code_value   string  `json:"code_value"`
		Is_published bool    `json:"is_published"`
		Expiration   string  `json:"expiration"`
		Price
	*/

	if p.Quantity != 0 {
		p.Quantity = p.Quantity
	}

	if p.Price != 0 {
		p.Price = p.Price
	}

	if p.Code_value != "" {
		p.Code_value = p.Code_value
	}

	return
}

func (ps *ProductSlice) DeleteProduct(id int) (err error) {
	if id <= 0 || id > len(ps.db) {
		err = errors.New("id nonexistent")
		return
	}

	//if the id is positional this shouldn't be needed, but if i delete an element, the ids of every other element should change or i should have a lastID field.
	for _, prod := range ps.db {
		if prod.ID == id {
			ps.db = append(ps.db[:id-1], ps.db[id:]...)
			return
		}
	}

	err = errors.New("product not found")
	return
}
