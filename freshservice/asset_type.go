package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AssetTypes contains Collection an array of AssetType
type AssetTypes struct {
	Collection []AssetType `json:"asset_types"`
}

// assetTypeWrapper contains Details of one AssetType
type assetTypeWrapper struct {
	Details AssetType `json:"asset_type"`
}

// AssetType represents a FreshService AssetType
type AssetType struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	ParentAssetTypeID int       `json:"parent_asset_type_id"`
	Visible           bool      `json:"visible"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// CreateAssetTypeModel is the data structure required to create a new AssetType
type CreateAssetTypeModel struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	ParentAssetTypeID int    `json:"parent_asset_type_id"`
}

// UpdateAssetTypeModel is the data structure required to update an AssetType
type UpdateAssetTypeModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Visible     bool   `json:"visible"`
}

// ListAssetTypesOptions represents filters/pagination for AssetTypes
type ListAssetTypesOptions struct {
	ListOptions
}

// GetAssetType returns an AssetType by id
func (s *AssetService) GetAssetType(id int) (*AssetType, *http.Response, error) {
	o := new(assetTypeWrapper)
	res, err := s.client.Get(fmt.Sprintf(fmt.Sprintf(assetTypeIdUrl, id), id), &o)
	return &o.Details, res, err
}

// ListAssetTypes will return paginated/filtered AssetTypes using ListAssetTypesOptions
func (s *AssetService) ListAssetTypes(opt *ListAssetTypesOptions) (*AssetTypes, *http.Response, error) {
	o := new(AssetTypes)
	res, err := s.client.List(assetTypesUrl, opt, &o)
	return o, res, err
}

// CreateAssetType creates and returns a new AssetType based on CreateAssetTypeModel
func (s *AssetService) CreateAssetType(newAssetType CreateAssetTypeModel) (*AssetType, *http.Response, error) {
	o := new(assetTypeWrapper)
	res, err := s.client.Post(assetTypesUrl, newAssetType, &o)
	return &o.Details, res, err
}

// UpdateAssetType updates and returns an AssetType matching id based on UpdateAssetTypeModel
func (s *AssetService) UpdateAssetType(id int, updatedAssetType UpdateAssetTypeModel) (*AssetType, *http.Response, error) {
	o := new(assetTypeWrapper)
	res, err := s.client.Put(fmt.Sprintf(assetTypeIdUrl, id), updatedAssetType, &o)
	return &o.Details, res, err
}

// DeleteAssetType irrecoverably deletes an AssetType from FreshService matching the id
func (s *AssetService) DeleteAssetType(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(assetTypeIdUrl, id))
	return success, res, err
}
