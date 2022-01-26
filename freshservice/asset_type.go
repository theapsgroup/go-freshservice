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
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(assetTypeIdUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	at := new(assetTypeWrapper)
	res, err := s.client.SendRequest(req, &at)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &at.Details, res, nil
}

// ListAssetTypes will return paginated/filtered AssetTypes using ListAssetTypesOptions
func (s *AssetService) ListAssetTypes(opt *ListAssetTypesOptions) (*AssetTypes, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, assetTypesUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	ats := new(AssetTypes)
	res, err := s.client.SendRequest(req, &ats)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return ats, res, nil
}

// CreateAssetType creates and returns a new AssetType based on CreateAssetTypeModel
func (s *AssetService) CreateAssetType(newAssetType CreateAssetTypeModel) (*AssetType, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, assetTypesUrl, newAssetType)
	if err != nil {
		return nil, nil, err
	}

	at := new(assetTypeWrapper)
	res, err := s.client.SendRequest(req, &at)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &at.Details, res, nil
}

// UpdateAssetType updates and returns an AssetType matching id based on UpdateAssetTypeModel
func (s *AssetService) UpdateAssetType(id int, updatedAssetType UpdateAssetTypeModel) (*AssetType, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(assetTypeIdUrl, id), updatedAssetType)
	if err != nil {
		return nil, nil, err
	}

	at := new(assetTypeWrapper)
	res, err := s.client.SendRequest(req, &at)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &at.Details, res, nil
}

// DeleteAssetType irrecoverably deletes an AssetType from FreshService matching the id
func (s *AssetService) DeleteAssetType(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(assetTypeIdUrl, id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
