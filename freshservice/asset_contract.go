package freshservice

import (
    "fmt"
    "net/http"
)

// AssetContracts contains collection of AssetContract
type AssetContracts struct {
    Collection []AssetContract `json:"contracts"`
}

// AssetContract contains small information about Contracts assigned to the Asset
type AssetContract struct {
    ID             int    `json:"id"`
    ContractID     string `json:"contract_id"`
    ContractType   string `json:"contract_type"`
    ContractName   string `json:"contract_name"`
    ContractStatus string `json:"contract_status"`
}

// GetAssetContracts obtains AssetContracts for a given Asset by displayId
func (s *AssetService) GetAssetContracts(displayId int) (*AssetContracts, *http.Response, error) {
    req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("assets/%v/contracts", displayId), nil)
    if err != nil {
        return nil, nil, err
    }

    acs := new(AssetContracts)
    res, err := s.client.SendRequest(req, &acs)
    if b, s := isSuccessful(res); !b {
        return nil, res, fmt.Errorf("%s: %v", s, err)
    }

    return acs, res, nil
}
