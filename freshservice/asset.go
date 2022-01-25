package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AssetService API Docs: https://api.freshservice.com/#assets https://api.freshservice.com/#asset-types
type AssetService struct {
	client *Client
}

// Assets contains collection of Asset
type Assets struct {
	Collection []Asset `json:"assets"`
}

// SpecificAsset contains Details of one specific Asset
type SpecificAsset struct {
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

// ListAssetOptions represents filters for Assets
type ListAssetOptions struct {
	ListOptions
}

// GetAsset will return a single Asset by displayId, assuming a record is found
func (s *AssetService) GetAsset(displayId int) (*Asset, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("assets/%v", displayId), nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAsset)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// GetAssets will return Assets collection
func (s *AssetService) GetAssets(opt *ListAssetOptions) (*Assets, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "assets", opt)
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

// CreateAsset will create a new Asset in FreshService
// TODO: Decide if need to implement a custom struct for newAsset since Asset struct has extraneous fields
func (s *AssetService) CreateAsset(newAsset *Asset) (*Asset, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "assets", newAsset)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAsset)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// UpdateAsset will update the Asset matching the displayId and return the updated Asset
func (s *AssetService) UpdateAsset(displayId int, asset *Asset) (*Asset, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("assets/%d", displayId), asset)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAsset)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// TrashAsset will trash the Asset matching the displayId (non-permanent delete)
func (s *AssetService) TrashAsset(displayId int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("assets/%d", displayId), nil)
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
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("assets/%d/restore", displayId), nil)
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
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("assets/%d/delete_forever", displayId), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
