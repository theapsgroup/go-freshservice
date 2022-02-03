package freshservice

import (
	"fmt"
	"net/http"
)

// AssetContracts contains Collection an array of AssetContract
type AssetContracts struct {
	Collection []AssetContract `json:"contracts"`
}

// AssetContract represents each contract link of an Asset
type AssetContract struct {
	ID             int    `json:"id"`
	ContractID     string `json:"contract_id"`
	ContractType   string `json:"contract_type"`
	ContractName   string `json:"contract_name"`
	ContractStatus string `json:"contract_status"`
}

// ListAssetContracts will return all AssetContracts for a given Asset by displayId
func (s *AssetService) ListAssetContracts(displayId int) (*AssetContracts, *http.Response, error) {
	o := new(AssetContracts)
	res, err := s.client.List(fmt.Sprintf(assetContractsUrl, displayId), nil, &o)
	return o, res, err
}
