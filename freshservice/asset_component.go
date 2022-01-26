package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AssetComponents contains Collection an array of AssetComponent
type AssetComponents struct {
	Collection []AssetComponent `json:"components"`
}

// AssetComponent represents each component of an Asset
type AssetComponent struct {
	ID            int         `json:"id"`
	ComponentType string      `json:"component_type"`
	ComponentData interface{} `json:"component_data"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// ListAssetComponents will return all AssetComponents for a given Asset by displayId
func (s *AssetService) ListAssetComponents(displayId int) (*AssetComponents, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(assetComponentsUrl, displayId), nil)
	if err != nil {
		return nil, nil, err
	}

	acs := new(AssetComponents)
	res, err := s.client.SendRequest(req, &acs)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return acs, res, nil
}
