package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	serviceCatalogItemUrl       = "service_catalog/items/%d"
	serviceCatalogItemsUrl      = "service_catalog/items"
	serviceCatalogItemSearchUrl = "service_catalog/items/search"
	serviceCategoriesUrl        = "service_catalog/categories"
)

// ServiceCatalogService API Docs: https://api.freshservice.com/#service-catalog
type ServiceCatalogService struct {
	client *Client
}

// ServiceItems contains Collection an array of ServiceItem
type ServiceItems struct {
	Collection []ServiceItem `json:"service_items"`
}

// serviceItemWrapper contains Details of a ServiceItem
type serviceItemWrapper struct {
	Details ServiceItem `json:"service_item"`
}

// ServiceItem represents a ServiceItem from the ServiceCatalog in FreshService
type ServiceItem struct {
	ID                     int       `json:"id"`
	Name                   string    `json:"name"`
	DeliveryTime           int       `json:"delivery_time"`
	DisplayID              int       `json:"display_id"`
	CategoryID             int       `json:"category_id"`
	ProductID              int       `json:"product_id"`
	Quantity               int       `json:"quantity"`
	Deleted                bool      `json:"deleted"`
	GroupVisibility        int       `json:"group_visibility"`
	ItemType               int       `json:"item_type"`
	CITypeID               int       `json:"ci_type_id"`
	CostVisibility         bool      `json:"cost_visibility"`
	DeliveryTimeVisibility bool      `json:"delivery_time_visibility"`
	Botified               bool      `json:"botified"`
	Visibility             int       `json:"visibility"`
	AllowAttachments       bool      `json:"allow_attachments"`
	AllowQuantity          bool      `json:"allow_quantity"`
	IsBundle               bool      `json:"is_bundle"`
	CreateChild            bool      `json:"create_child"`
	Description            string    `json:"description"`
	ShortDescription       string    `json:"short_description"`
	Cost                   float32   `json:"cost"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

// ServiceCategories contains Collection an array of ServiceCategory
type ServiceCategories struct {
	Collection []ServiceCategory `json:"service_categories"`
}

// ServiceCategory represents a ServiceCategory in FreshService
type ServiceCategory struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Position    int       `json:"position"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ServiceItemSearch struct {
	ListOptions
	SearchTerm string `json:"search_term,omitempty" url:"search_term,omitempty"`
	UserEmail  string `json:"user_email,omitempty" url:"user_email,omitempty"`
}

// GetServiceItem will return a ServiceItem by displayId
func (s *ServiceCatalogService) GetServiceItem(displayId int) (*ServiceItem, *http.Response, error) {
	o := new(serviceItemWrapper)
	res, err := s.client.Get(fmt.Sprintf(serviceCatalogItemUrl, displayId), &o)
	return &o.Details, res, err
}

// ListServiceItems will return ServiceItems
func (s *ServiceCatalogService) ListServiceItems() (*ServiceItems, *http.Response, error) {
	o := new(ServiceItems)
	res, err := s.client.List(serviceCatalogItemsUrl, nil, &o)
	return o, res, err
}

// SearchServiceItems will return paginated/filtered ServiceItems based on ServiceItemSearch
func (s *ServiceCatalogService) SearchServiceItems(search *ServiceItemSearch) (*ServiceItems, *http.Response, error) {
	o := new(ServiceItems)
	res, err := s.client.List(serviceCatalogItemSearchUrl, search, &o)
	return o, res, err
}
