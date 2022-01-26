package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	vendorsUrl  = "vendors"
	vendorIdUrl = "vendors/%d"
)

// VendorService API Docs: https://api.freshservice.com/#vendors
type VendorService struct {
	client *Client
}

// Vendors contains Collection an array of Vendor
type Vendors struct {
	Collection []Vendor `json:"vendors"`
}

// vendorWrapper contains Details of one Vendor
type vendorWrapper struct {
	Details Vendor `json:"vendor"`
}

// Vendor represents a FreshService Vendor
type Vendor struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	PrimaryContactID int           `json:"primary_contact_id"`
	Address          VendorAddress `json:"address"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

// VendorAddress is an alternative to Address but with only one address line
type VendorAddress struct {
	Line1   string `json:"line1"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipcode"`
}

// CreateVendorModel is the data structure required to create a new Vendor
type CreateVendorModel struct {
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	PrimaryContactID int           `json:"primary_contact_id"`
	Address          VendorAddress `json:"address"`
}

// UpdateVendorModel is the data structure required to update a Vendor
type UpdateVendorModel struct {
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	PrimaryContactID int           `json:"primary_contact_id"`
	Address          VendorAddress `json:"address"`
}

// ListVendorsOptions represents filters/pagination for Vendors
type ListVendorsOptions struct {
	ListOptions
}

// GetVendor will return a single Vendor by id
func (s *VendorService) GetVendor(id int) (*Vendor, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(vendorIdUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(vendorWrapper)
	res, err := s.client.SendRequest(req, &v)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &v.Details, res, nil
}

// ListVendors will return paginated/filtered Vendors using ListVendorsOptions
func (s *VendorService) ListVendors(opt *ListVendorsOptions) (*Vendors, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, vendorsUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	vs := new(Vendors)
	res, err := s.client.SendRequest(req, &vs)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return vs, res, nil
}

// CreateVendor will create and return a new Vendor based on CreateVendorModel
func (s *VendorService) CreateVendor(newVendor *CreateVendorModel) (*Vendor, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, vendorsUrl, newVendor)
	if err != nil {
		return nil, nil, err
	}

	v := new(vendorWrapper)
	res, err := s.client.SendRequest(req, &v)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &v.Details, res, nil
}

// UpdateVendor will update and return a Vendor matching id based on UpdateVendorModel
func (s *VendorService) UpdateVendor(id int, vendor *UpdateVendorModel) (*Vendor, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(vendorIdUrl, id), vendor)
	if err != nil {
		return nil, nil, err
	}

	v := new(vendorWrapper)
	res, err := s.client.SendRequest(req, &v)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &v.Details, res, nil
}

// DeleteVendor will completely remove a Vendor from FreshService matching id
func (s *VendorService) DeleteVendor(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(vendorIdUrl, id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
