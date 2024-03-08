package internal

import (
	"errors"
	"fmt"
	"time"
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

type ProductRepository interface {
	Save(product *Product) (err error)
	GetProducts() (data string)
	GetProductByKey(id int) (product Product, err error)
	GetProductsGreaterThanPrice(price float64) (data string)
	UpdateProduct(product *Product, id int) (err error)
	LoadProducts() (err error)
	PatchProduct(dataMap map[string]any, id int) (err error)
	DeleteProduct(id int) (err error)
}

type ProductService interface {
	Save(product *Product) (err error)
	GetProducts() (data string)
	GetProductByKey(id int) (product Product, err error)
	GetProductsGreaterThanPrice(price float64) (data string)
	UpdateProduct(product *Product, id int) (err error)
	PatchProduct(product *Product, id int) (err error)
	DeleteProduct(id int) (err error)
}

/*
func NewProductStorage(filename string) ProductStorage {
	return ProductStorage{
		filename: filename,
		Products: []Product{},
	}
}
*/

func (p Product) Validate() (err error) {

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

func (p Product) String() string {
	return fmt.Sprintf("ID: %d\nName: %s\nQuantity: %d\nCode_value: %s\nIs_published: %t\nExpiration: %s\nPrice: %.2f\n", p.ID, p.Name, p.Quantity, p.Code_value, p.Is_published, p.Expiration, p.Price)
}

func validateExpiration(expiration string) bool {
	_, err := time.Parse("02/01/2006", expiration)
	return err == nil
}
