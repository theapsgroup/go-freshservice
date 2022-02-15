package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	contractsUrl                = "contracts"
	contractIdUrl               = "contracts/%d"
	contractTypesUrl            = "contract_types"
	contractSubmitApprovalUrl   = "contracts/%d?operation=submit-for-approval"
	contractApproveUrl          = "contracts/%d?operation=approve"
	contractRejectUrl           = "contracts/%d?operation=reject"
	contractAssociatedAssetsUrl = "contracts/%d/associated_assets"
)

// ContractService API Docs: https://api.freshservice.com/#contracts
type ContractService struct {
	client *Client
}

// Contracts contains Collection an array of Contract
type Contracts struct {
	Collection []Contract `json:"contracts"`
}

// contractWrapper contains Details of one Contract
type contractWrapper struct {
	Details Contract `json:"contract"`
}

// Contract represents a Contract in the FreshService instance.
type Contract struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	VendorID        int              `json:"vendor_id"`
	AutoRenew       bool             `json:"auto_renew"`
	NotifyExpiry    bool             `json:"notify_expiry"`
	NotifyBefore    int              `json:"notify_before"`
	ApproverID      int              `json:"approver_id"`
	StartDate       time.Time        `json:"start_date"`
	EndDate         time.Time        `json:"end_date"`
	Cost            float32          `json:"cost"`
	Status          string           `json:"status"`
	ContractNumber  string           `json:"contract_number"`
	ContractTypeID  int              `json:"contract_type_id"`
	VisibleToID     int              `json:"visible_to_id"`
	NotifyTo        []string         `json:"notify_to"`
	ExpiryNotified  bool             `json:"expiry_notified"`
	RequesterID     int              `json:"requester_id"`
	DelegateeID     int              `json:"delegatee_id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	SoftwareID      int              `json:"software_id"`
	LicenseType     string           `json:"license_type"`
	BillingCycle    string           `json:"billing_cycle"`
	LicenseKey      string           `json:"license_key"`
	ItemCostDetails []ItemCostDetail `json:"item_cost_details"`
}

// CreateContractModel is the data structure required to create a new Contract
type CreateContractModel struct {
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`
	VendorID       int       `json:"vendor_id"`
	AutoRenew      bool      `json:"auto_renew,omitempty"`
	NotifyExpiry   bool      `json:"notify_expiry,omitempty"`
	NotifyBefore   int       `json:"notify_before,omitempty"`
	ApproverID     int       `json:"approver_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Cost           float32   `json:"cost"`
	ContractNumber string    `json:"contract_number"`
	ContractTypeID int       `json:"contract_type_id"`
	VisibleToID    int       `json:"visible_to_id"`
	NotifyTo       []string  `json:"notify_to,omitempty"`
	SoftwareID     int       `json:"software_id,omitempty"`
	LicenseType    string    `json:"license_type,omitempty"`
	BillingCycle   string    `json:"billing_cycle,omitempty"`
	LicenseKey     string    `json:"license_key,omitempty"`
}

// UpdateContractModel is the data structure required to update a Contract
type UpdateContractModel struct {
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`
	VendorID       int       `json:"vendor_id"`
	AutoRenew      bool      `json:"auto_renew,omitempty"`
	NotifyExpiry   bool      `json:"notify_expiry,omitempty"`
	NotifyBefore   int       `json:"notify_before,omitempty"`
	ApproverID     int       `json:"approver_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Cost           float32   `json:"cost"`
	ContractNumber string    `json:"contract_number"`
	ContractTypeID int       `json:"contract_type_id"`
	VisibleToID    int       `json:"visible_to_id"`
	NotifyTo       []string  `json:"notify_to,omitempty"`
	SoftwareID     int       `json:"software_id,omitempty"`
	LicenseType    string    `json:"license_type,omitempty"`
	BillingCycle   string    `json:"billing_cycle,omitempty"`
	LicenseKey     string    `json:"license_key,omitempty"`
}

// ItemCostDetail represents a line-item cost for a Contract
type ItemCostDetail struct {
	ID           int       `json:"id"`
	ItemName     string    `json:"item_name"`
	PricingModel string    `json:"pricing_model"`
	Cost         float32   `json:"cost"`
	Count        int       `json:"count"`
	Comments     string    `json:"comments"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ListContractsOptions is for filtering/pagination of Contracts
type ListContractsOptions struct {
	ListOptions
}

// GetContract will return a single Contract by id
func (s *ContractService) GetContract(id int) (*Contract, *http.Response, error) {
	o := new(contractWrapper)
	res, err := s.client.Get(fmt.Sprintf(contractIdUrl, id), &o)
	return &o.Details, res, err
}

// ListContracts will return paginated/filtered Contracts using ListContractsOptions
func (s *ContractService) ListContracts(opt *ListContractsOptions) (*Contracts, *http.Response, error) {
	o := new(Contracts)
	res, err := s.client.List(contractsUrl, opt, &o)
	return o, res, err
}

// CreateContract will create and return a new Contract based on CreateContractModel
func (s *ContractService) CreateContract(contract *CreateContractModel) (*Contract, *http.Response, error) {
	o := new(contractWrapper)
	res, err := s.client.Post(contractsUrl, contract, &o)
	return &o.Details, res, err
}

// UpdateContract will update and return a Contract matching id based on UpdateContractModel
func (s *ContractService) UpdateContract(id int, contract *UpdateContractModel) (*Contract, *http.Response, error) {
	o := new(contractWrapper)
	res, err := s.client.Put(fmt.Sprintf(contractIdUrl, id), contract, &o)
	return &o.Details, res, err
}

// SubmitContractApproval allows for a Contract to be submitted for approval
func (s *ContractService) SubmitContractApproval(id int) (bool, *http.Response, error) {
	res, err := s.client.Put(fmt.Sprintf(contractSubmitApprovalUrl, id), nil, nil)
	success, _ := isSuccessful(res)
	return success, res, err
}

// ApproveContract allows for a Contract to be Approved
func (s *ContractService) ApproveContract(id int) (bool, *http.Response, error) {
	res, err := s.client.Put(fmt.Sprintf(contractApproveUrl, id), nil, nil)
	success, _ := isSuccessful(res)
	return success, res, err
}

// RejectContract rejects the Contract that was submitted for approval
func (s *ContractService) RejectContract(id int) (bool, *http.Response, error) {
	res, err := s.client.Put(fmt.Sprintf(contractRejectUrl, id), nil, nil)
	success, _ := isSuccessful(res)
	return success, res, err
}
