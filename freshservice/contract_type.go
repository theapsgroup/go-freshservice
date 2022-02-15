package freshservice

import (
	"net/http"
	"time"
)

// ContractTypes contains Collection an array of ContractType
type ContractTypes struct {
	Collection []ContractType `json:"contract_types"`
}

// ContractType represents a type of Contract
type ContractType struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	NeedsApproval bool      `json:"needs_approval"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ListContractTypes will return ContractTypes
func (s *ContractService) ListContractTypes() (*ContractTypes, *http.Response, error) {
	o := new(ContractTypes)
	res, err := s.client.List(contractTypesUrl, nil, &o)
	return o, res, err
}
