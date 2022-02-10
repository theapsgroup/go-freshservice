package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	changesUrl       = "changes"
	changeIdUrl      = "changes/%d"
	changeRestoreUrl = "changes/%d/restore"
	changeNotesUrl   = "changes/%d/notes"
	changeNoteIdUrl  = "changes/%d/notes/%d"
)

const (
	ChangeStatusOpen           = 1
	ChangeStatusPlanning       = 2
	ChangeStatusApproval       = 3
	ChangeStatusPendingRelease = 4
	ChangeStatusPendingReview  = 5
	ChangeStatusClosed         = 6
	ChangePriorityLow          = 1
	ChangePriorityMedium       = 2
	ChangePriorityHigh         = 3
	ChangePriorityUrgent       = 4
	ChangeImpactLow            = 1
	ChangeImpactMedium         = 2
	ChangeImpactHigh           = 3
	ChangeTypeMinor            = 1
	ChangeTypeStandard         = 2
	ChangeTypeMajor            = 3
	ChangeTypeEmergency        = 4
	ChangeRiskLow              = 1
	ChangeRiskMedium           = 2
	ChangeRiskHigh             = 3
	ChangeRiskVeryHigh         = 4
)

// ChangeService API Docs: https://api.freshservice.com/#changes
type ChangeService struct {
	client *Client
}

// Changes contains Collection an array of Change
type Changes struct {
	Collection []Change `json:"changes"`
}

// changeWrapper contains Details of one Change
type changeWrapper struct {
	Details Change `json:"change"`
}

// Change represents a Change request on FreshService
type Change struct {
	ID               int       `json:"id"`
	AgentID          int       `json:"agent_id"`
	Description      string    `json:"description"`
	DescriptionText  string    `json:"description_text"`
	RequesterID      int       `json:"requester_id"`
	GroupID          int       `json:"group_id"`
	Priority         int       `json:"priority"`
	Impact           int       `json:"impact"`
	Status           int       `json:"status"`
	Risk             int       `json:"risk"`
	ChangeType       int       `json:"change_type"`
	ApprovalStatus   int       `json:"approval_status"`
	PlannedStartDate time.Time `json:"planned_start_date"`
	PlannedEndDate   time.Time `json:"planned_end_date"`
	Subject          string    `json:"subject"`
	DepartmentID     int       `json:"department_id"`
	Category         string    `json:"category"`
	SubCategory      string    `json:"sub_category"`
	ItemCategory     string    `json:"item_category"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CreateChangeModel is a data struct for creating a new Change
type CreateChangeModel struct {
	AgentID          int       `json:"agent_id"`
	Description      string    `json:"description"`
	Subject          string    `json:"subject"`
	GroupID          int       `json:"group_id"`
	Priority         int       `json:"priority"`
	Impact           int       `json:"impact"`
	Status           int       `json:"status"`
	Risk             int       `json:"risk"`
	ChangeType       int       `json:"change_type"`
	ApprovalStatus   int       `json:"approval_status"`
	PlannedStartDate time.Time `json:"planned_start_date"`
	PlannedEndDate   time.Time `json:"planned_end_date"`
	DepartmentID     int       `json:"department_id"`
}

// UpdateChangeModel is a data struct for updating a Change
type UpdateChangeModel struct {
	AgentID          int       `json:"agent_id"`
	Description      string    `json:"description"`
	DescriptionText  string    `json:"description_text"`
	RequesterID      int       `json:"requester_id"`
	GroupID          int       `json:"group_id"`
	Priority         int       `json:"priority"`
	Impact           int       `json:"impact"`
	Status           int       `json:"status"`
	Risk             int       `json:"risk"`
	ChangeType       int       `json:"change_type"`
	ApprovalStatus   int       `json:"approval_status"`
	PlannedStartDate time.Time `json:"planned_start_date"`
	PlannedEndDate   time.Time `json:"planned_end_date"`
	Subject          string    `json:"subject"`
	DepartmentID     int       `json:"department_id"`
	Category         string    `json:"category"`
	SubCategory      string    `json:"sub_category"`
	ItemCategory     string    `json:"item_category"`
}

// ListChangesOptions represents filters/pagination for Changes
type ListChangesOptions struct {
	ListOptions
	Filter       *string    `json:"filter,omitempty" url:"filter,omitempty"`
	RequesterID  *int       `json:"requester_id,omitempty" url:"requester_id,omitempty"`
	UpdatedSince *time.Time `json:"updated_since,omitempty" url:"updated_since,omitempty"`
}

// GetChange will return a single Change by id
func (s *ChangeService) GetChange(id int) (*Change, *http.Response, error) {
	o := new(changeWrapper)
	res, err := s.client.Get(fmt.Sprintf(changeIdUrl, id), &o)
	return &o.Details, res, err
}

// ListChanges will return paginated/filtered Change using ListChangesOptions
func (s *ChangeService) ListChanges(opt *ListChangesOptions) (*Changes, *http.Response, error) {
	o := new(Changes)
	res, err := s.client.List(changesUrl, opt, &o)
	return o, res, err
}

// CreateChange will create and return a new Change based on CreateChangeModel
func (s *ChangeService) CreateChange(newChange *CreateChangeModel) (*Change, *http.Response, error) {
	o := new(changeWrapper)
	res, err := s.client.Post(changesUrl, newChange, &o)
	return &o.Details, res, err
}

// UpdateChange will update and return a Change matching id based on UpdateChangeModel
func (s *ChangeService) UpdateChange(id int, ticket *UpdateChangeModel) (*Change, *http.Response, error) {
	o := new(changeWrapper)
	res, err := s.client.Put(fmt.Sprintf(changeIdUrl, id), ticket, &o)
	return &o.Details, res, err
}

// DeleteChange will trash a Change from FreshService (Can be restored by RestoreChange)
func (s *ChangeService) DeleteChange(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(changeIdUrl, id))
	return success, res, err
}

// RestoreChange will restore a previously trashed (deleted) Change
func (s *ChangeService) RestoreChange(id int) (bool, *http.Response, error) {
	res, err := s.client.Put(fmt.Sprintf(changeRestoreUrl, id), nil, nil)
	success, _ := isSuccessful(res)
	return success, res, err
}
