package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	purchaseOrdersUrl  = "purchase_orders"
	purchaseOrderIdUrl = "purchase_orders/%d"
)

// PurchaseOrderService API Docs: https://api.freshservice.com/#purchase-order
type PurchaseOrderService struct {
	client *Client
}

// PurchaseOrders contains Collection an array of PurchaseOrder
type PurchaseOrders struct {
	Collection []PurchaseOrder `json:"purchase_orders"`
}

// poWrapper contains Details of one PurchaseOrder
type poWrapper struct {
	Details PurchaseOrder `json:"purchase_order"`
}

// PurchaseOrder represents a PurchaseOrder in FreshService
type PurchaseOrder struct {
	ID                    int            `json:"id"`
	VendorID              int            `json:"vendor_id"`
	Name                  string         `json:"name"`
	PurchaseOrderNumber   string         `json:"po_number"`
	VendorDetails         string         `json:"vendor_details"`
	ExpectedDeliveryDate  time.Time      `json:"expected_delivery_date"`
	CreatedBy             int            `json:"created_by"`
	Status                int            `json:"status"`
	ShippingAddress       string         `json:"shipping_address"`
	BillingAddress        string         `json:"billing_address"`
	BillingSameAsShipping bool           `json:"billing_same_as_shipping"`
	CurrencyCode          string         `json:"currency_code"`
	ConversionRate        float32        `json:"conversion_rate"`
	DepartmentID          int            `json:"department_id"`
	DiscountPercentage    float32        `json:"discount_percentage"`
	TaxPercentage         float32        `json:"tax_percentage"`
	ShoppingCost          float32        `json:"shopping_cost"`
	PurchaseItems         []PurchaseItem `json:"purchase_items"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}

// PurchaseItem represents a line item on a PurchaseOrder
type PurchaseItem struct {
	ItemType      int     `json:"item_type"`
	ItemName      string  `json:"item_name"`
	Description   string  `json:"description"`
	Cost          float32 `json:"cost"`
	Quantity      int     `json:"quantity"`
	TaxPercentage float32 `json:"tax_percentage"`
}

// CreatePurchaseOrderModel is a data struct for creating a new PurchaseOrder
type CreatePurchaseOrderModel struct {
	VendorID              int            `json:"vendor_id"`
	Name                  string         `json:"name"`
	PurchaseOrderNumber   string         `json:"po_number"`
	VendorDetails         string         `json:"vendor_details"`
	ExpectedDeliveryDate  time.Time      `json:"expected_delivery_date"`
	ShippingAddress       string         `json:"shipping_address"`
	BillingAddress        string         `json:"billing_address"`
	BillingSameAsShipping bool           `json:"billing_same_as_shipping"`
	CurrencyCode          string         `json:"currency_code"`
	ConversionRate        float32        `json:"conversion_rate"`
	DepartmentID          int            `json:"department_id"`
	DiscountPercentage    float32        `json:"discount_percentage"`
	TaxPercentage         float32        `json:"tax_percentage"`
	ShoppingCost          float32        `json:"shopping_cost"`
	PurchaseItems         []PurchaseItem `json:"purchase_items"`
}

// UpdatePurchaseOrderModel is a data struct for updating a PurchaseOrder
type UpdatePurchaseOrderModel struct {
	VendorID              int            `json:"vendor_id"`
	Name                  string         `json:"name"`
	PurchaseOrderNumber   string         `json:"po_number"`
	VendorDetails         string         `json:"vendor_details"`
	ExpectedDeliveryDate  time.Time      `json:"expected_delivery_date"`
	ShippingAddress       string         `json:"shipping_address"`
	BillingAddress        string         `json:"billing_address"`
	BillingSameAsShipping bool           `json:"billing_same_as_shipping"`
	CurrencyCode          string         `json:"currency_code"`
	ConversionRate        float32        `json:"conversion_rate"`
	DepartmentID          int            `json:"department_id"`
	DiscountPercentage    float32        `json:"discount_percentage"`
	TaxPercentage         float32        `json:"tax_percentage"`
	ShoppingCost          float32        `json:"shopping_cost"`
	PurchaseItems         []PurchaseItem `json:"purchase_items"`
}

// ListPurchaseOrdersOptions represents filters/pagination for PurchaseOrders
type ListPurchaseOrdersOptions struct {
	ListOptions
}

// GetPurchaseOrder will return a single PurchaseOrder by id
func (s *PurchaseOrderService) GetPurchaseOrder(id int) (*PurchaseOrder, *http.Response, error) {
	o := new(poWrapper)
	res, err := s.client.Get(fmt.Sprintf(purchaseOrderIdUrl, id), &o)
	return &o.Details, res, err
}

// ListPurchaseOrders will return paginated/filtered PurchaseOrders using ListPurchaseOrdersOptions
func (s *PurchaseOrderService) ListPurchaseOrders(opt *ListPurchaseOrdersOptions) (*PurchaseOrders, *http.Response, error) {
	o := new(PurchaseOrders)
	res, err := s.client.List(purchaseOrdersUrl, opt, &o)
	return o, res, err
}

// CreatePurchaseOrder will create and return a new PurchaseOrder based on CreatePurchaseOrderModel
func (s *PurchaseOrderService) CreatePurchaseOrder(newPurchaseOrder *CreatePurchaseOrderModel) (*PurchaseOrder, *http.Response, error) {
	o := new(poWrapper)
	res, err := s.client.Post(purchaseOrdersUrl, newPurchaseOrder, &o)
	return &o.Details, res, err
}

// UpdatePurchaseOrder will update and return an PurchaseOrder matching id based on UpdatePurchaseOrderModel
func (s *PurchaseOrderService) UpdatePurchaseOrder(id int, purchaseOrder *UpdatePurchaseOrderModel) (*PurchaseOrder, *http.Response, error) {
	o := new(poWrapper)
	res, err := s.client.Put(fmt.Sprintf(purchaseOrderIdUrl, id), purchaseOrder, &o)
	return &o.Details, res, err
}

// DeletePurchaseOrder will completely remove an PurchaseOrder from FreshService
func (s *PurchaseOrderService) DeletePurchaseOrder(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(purchaseOrderIdUrl, id))
	return success, res, err
}
