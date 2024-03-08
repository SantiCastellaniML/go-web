package service

import (
	"apiProductsStructure/internal"
)

type ProductServiceIF interface {
}

type ProductService struct {
	rp internal.ProductRepository
}

func NewProductService(rp internal.ProductRepository) *ProductService {
	return &ProductService{
		rp: rp,
	}
}

func (ps *ProductService) Save(product *internal.Product) (err error) {
	err = product.Validate()
	if err != nil {
		return
	}

	//this goes in repository
	/*
		err = ps.validateCodeValue(product.Code_value)
		if err != nil {
			return
		}
	*/

	err = ps.rp.Save(product)
	return
}

func (ps *ProductService) GetProducts() (data string) {
	data = ps.rp.GetProducts()

	return
}

func (ps *ProductService) GetProductByKey(key int) (product internal.Product, err error) {
	product, err = ps.rp.GetProductByKey(key)
	return
}

func (ps *ProductService) GetProductsGreaterThanPrice(price float64) (data string) {
	data = ps.rp.GetProductsGreaterThanPrice(price)

	return
}

func (ps *ProductService) UpdateProduct(product *internal.Product, id int) (err error) {
	err = product.Validate()
	if err != nil {
		return
	}

	err = ps.rp.UpdateProduct(product, id)
	return
}

func (ps *ProductService) PatchProduct(dataMap map[string]any, id int) (err error) {
	err = ps.rp.PatchProduct(dataMap, id)
	return
}

func (ps *ProductService) DeleteProduct(id int) (err error) {
	err = ps.rp.DeleteProduct(id)
	return
}
