package freshservice

import (
	"fmt"
	"net/http"
)

// AssociatedAssets contains Collection an array of Asset
type AssociatedAssets struct {
	Collection []Asset `json:"associated_assets"`
}

func (s *ContractService) ListContractAssociatedAssets(id int) (*AssociatedAssets, *http.Response, error) {
	o := new(AssociatedAssets)
	res, err := s.client.List(fmt.Sprintf(contractAssociatedAssetsUrl, id), nil, &o)
	return o, res, err
}
