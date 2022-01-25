package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AssetTypes contains collection of AssetType
type AssetTypes struct {
	Collection []AssetType `json:"asset_types"`
}

// SpecificAssetType contains Details of one specific AssetType
type SpecificAssetType struct {
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

// NewAssetType is the data structure required to create a new AssetType
type NewAssetType struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	ParentAssetTypeID int    `json:"parent_asset_type_id"`
}

// UpdateAssetType is the data structure required to update an AssetType
type UpdateAssetType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Visible     bool   `json:"visible"`
}

// GetAssetType returns an AssetType by id
func (s *AssetService) GetAssetType(id int) (*AssetType, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("asset_types/%v", id), nil)
	if err != nil {
		return nil, nil, err
	}

	at := new(SpecificAssetType)
	res, err := s.client.SendRequest(req, &at)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &at.Details, res, nil
}

// GetAssetTypes will return AssetTypes collection
func (s *AssetService) GetAssetTypes(opt *ListOptions) (*AssetTypes, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "asset_types", opt)
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

// CreateAssetType creates and returns a new AssetType
func (s *AssetService) CreateAssetType(newAssetType NewAssetType) (*AssetType, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "asset_types", newAssetType)
	if err != nil {
		return nil, nil, err
	}

	at := new(SpecificAssetType)
	res, err := s.client.SendRequest(req, &at)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &at.Details, res, nil
}

// UpdateAssetType updates and returns an AssetType
func (s *AssetService) UpdateAssetType(id int, updatedAssetType UpdateAssetType) (*AssetType, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("asset_types/%d", id), updatedAssetType)
	if err != nil {
		return nil, nil, err
	}

	at := new(SpecificAssetType)
	res, err := s.client.SendRequest(req, &at)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &at.Details, res, nil
}

// DeleteAssetType deletes an AssetType
func (s *AssetService) DeleteAssetType(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("asset_types/%d", id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
