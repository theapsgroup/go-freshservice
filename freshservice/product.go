package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	productsUrl  = "products"
	productIdUrl = "products/%d"
)

// ProductService API Docs: https://api.freshservice.com/#products
type ProductService struct {
	client *Client
}

// Products contains Collection an array of Product
type Products struct {
	Collection []Product `json:"products"`
}

// productWrapper contains Details of a Product
type productWrapper struct {
	Details Product `json:"product"`
}

// Product represents a FreshService Product
type Product struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	AssetTypeID        int       `json:"asset_type_id"`
	Manufacturer       string    `json:"manufacturer"`
	Status             string    `json:"status"`
	ModeOfProcurement  string    `json:"mode_of_procurement"`
	DepreciationTypeID int       `json:"depreciation_type_id"`
	DescriptionText    string    `json:"description_text"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// CreateProductModel is a data struct for creating a new Product
type CreateProductModel struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	AssetTypeID        int    `json:"asset_type_id"`
	Manufacturer       string `json:"manufacturer"`
	Status             string `json:"status"`
	ModeOfProcurement  string `json:"mode_of_procurement"`
	DepreciationTypeID int    `json:"depreciation_type_id"`
	DescriptionText    string `json:"description_text"`
}

// UpdateProductModel is a data struct for updating a Product
type UpdateProductModel struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	AssetTypeID        int    `json:"asset_type_id"`
	Manufacturer       string `json:"manufacturer"`
	Status             string `json:"status"`
	ModeOfProcurement  string `json:"mode_of_procurement"`
	DepreciationTypeID int    `json:"depreciation_type_id"`
	DescriptionText    string `json:"description_text"`
}

// ListProductsOptions represents filters/pagination for Products
type ListProductsOptions struct {
	ListOptions
}

// GetProduct will return a Product by id
func (s *ProductService) GetProduct(id int) (*Product, *http.Response, error) {
	o := new(productWrapper)
	res, err := s.client.Get(fmt.Sprintf(productIdUrl, id), &o)
	return &o.Details, res, err
}

// ListProducts will return paginated/filtered Products using ListProductsOptions
func (s *ProductService) ListProducts(opt *ListProductsOptions) (*Products, *http.Response, error) {
	o := new(Products)
	res, err := s.client.List(productsUrl, opt, &o)
	return o, res, err
}

// CreateProduct will create and return a new Product based on CreateProductModel
func (s *ProductService) CreateProduct(newProduct *CreateProductModel) (*Product, *http.Response, error) {
	o := new(productWrapper)
	res, err := s.client.Post(productsUrl, newProduct, &o)
	return &o.Details, res, err
}

// UpdateProduct will update and return a Product matching id based UpdateProductModel
func (s *ProductService) UpdateProduct(id int, product *UpdateLocationModel) (*Product, *http.Response, error) {
	o := new(productWrapper)
	res, err := s.client.Put(fmt.Sprintf(productIdUrl, id), product, &o)
	return &o.Details, res, err
}

// DeleteProduct will completely remove a Product from FreshService matching id
func (s *ProductService) DeleteProduct(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(productIdUrl, id))
	return success, res, err
}
