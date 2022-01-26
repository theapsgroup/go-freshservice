package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	assetsUrl          = "assets"
	assetIdUrl         = "assets/%d"
	assetRestoreUrl    = "assets/%d/restore"
	assetDeleteUrl     = "assets/%d/delete_forever"
	assetComponentsUrl = "assets/%d/components"
	assetContractsUrl  = "assets/%d/contracts"
	assetTypesUrl      = "asset_types"
	assetTypeIdUrl     = "asset_types/%d"
)

// AssetService API Docs: https://api.freshservice.com/#assets https://api.freshservice.com/#asset-types
type AssetService struct {
	client *Client
}

// Assets contains Collection an array of Asset
type Assets struct {
	Collection []Asset `json:"assets"`
}

// assetWrapper contains Details of one Asset
type assetWrapper struct {
	Details Asset `json:"asset"`
}

// Asset represents a FreshService Asset
type Asset struct {
	ID           int       `json:"id"`
	DisplayID    int       `json:"display_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	AssetTypeID  int       `json:"asset_type_id"`
	AssetTag     string    `json:"asset_tag"`
	Impact       string    `json:"impact"`
	AuthorType   string    `json:"author_type"`
	UsageType    string    `json:"usage_type"`
	UserID       int       `json:"user_id"`
	LocationID   int       `json:"location_id"`
	DepartmentID int       `json:"department_id"`
	AgentID      int       `json:"agent_id"`
	GroupID      int       `json:"group_id"`
	AssignedOn   time.Time `json:"assigned_on"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateAssetModel is the data structure required to create a new Asset
type CreateAssetModel struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	AssetTypeID  int       `json:"asset_type_id"`
	AssetTag     string    `json:"asset_tag"`
	Impact       string    `json:"impact"`
	UsageType    string    `json:"usage_type"`
	UserID       int       `json:"user_id"`
	LocationID   int       `json:"location_id"`
	DepartmentID int       `json:"department_id"`
	AgentID      int       `json:"agent_id"`
	GroupID      int       `json:"group_id"`
	AssignedOn   time.Time `json:"assigned_on"`
}

// UpdateAssetModel is the data structure required to update an Asset
type UpdateAssetModel struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	AssetTypeID  int       `json:"asset_type_id"`
	AssetTag     string    `json:"asset_tag"`
	Impact       string    `json:"impact"`
	UsageType    string    `json:"usage_type"`
	UserID       int       `json:"user_id"`
	LocationID   int       `json:"location_id"`
	DepartmentID int       `json:"department_id"`
	AgentID      int       `json:"agent_id"`
	GroupID      int       `json:"group_id"`
	AssignedOn   time.Time `json:"assigned_on"`
}

// ListAssetsOptions represents filters/pagination for Assets
type ListAssetsOptions struct {
	ListOptions
}

// GetAsset will return a single Asset by displayId
func (s *AssetService) GetAsset(displayId int) (*Asset, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(assetIdUrl, displayId), nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(assetWrapper)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// ListAssets will return paginated/filtered Assets using ListAssetsOptions
func (s *AssetService) ListAssets(opt *ListAssetsOptions) (*Assets, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, assetsUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	as := new(Assets)
	res, err := s.client.SendRequest(req, &as)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return as, res, nil
}

// CreateAsset will create and return a new Asset based on CreateAssetModel
func (s *AssetService) CreateAsset(newAsset *CreateAssetModel) (*Asset, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, assetsUrl, newAsset)
	if err != nil {
		return nil, nil, err
	}

	a := new(assetWrapper)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// UpdateAsset will update and return an Asset matching displayId based on UpdateAssetModel
func (s *AssetService) UpdateAsset(displayId int, asset *UpdateAssetModel) (*Asset, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(assetIdUrl, displayId), asset)
	if err != nil {
		return nil, nil, err
	}

	a := new(assetWrapper)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// TrashAsset will trash the Asset matching the displayId (non-permanent delete)
func (s *AssetService) TrashAsset(displayId int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(assetIdUrl, displayId), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}

// RestoreAsset will restore a previously Trashed Asset by displayId
func (s *AssetService) RestoreAsset(displayId int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(assetRestoreUrl, displayId), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}

// DeleteAsset irrecoverably removes an Asset from FreshService matching the displayId
func (s *AssetService) DeleteAsset(displayId int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(assetDeleteUrl, displayId), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
