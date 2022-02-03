package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// SolutionFolders contains Collection an array of SolutionFolder
type SolutionFolders struct {
	Collection []SolutionFolder `json:"folders"`
}

// solutionFolderWrapper contains Details of one SolutionFolder
type solutionFolderWrapper struct {
	Details SolutionFolder `json:"folder"`
}

// SolutionFolder represents a FreshService SolutionFolder
type SolutionFolder struct {
	ID                int              `json:"id"`
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	Position          int              `json:"position"`
	DefaultFolder     bool             `json:"default_folder"`
	CategoryID        int              `json:"category_id"`
	Visibility        int              `json:"visibility"`
	DepartmentIDs     []int            `json:"department_ids"`
	GroupIDs          []int            `json:"group_ids"`
	RequesterGroupIDs []int            `json:"requester_group_ids"`
	ManageByGroupIDs  []int            `json:"manage_by_group_ids"`
	ApprovalSettings  ApprovalSettings `json:"approval_settings"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

type ApprovalSettings struct {
	ApprovalType int   `json:"approval_type"`
	ApproverIDs  []int `json:"approver_ids"`
}

// CreateSolutionFolderModel is a data struct for creating a new SolutionFolder
type CreateSolutionFolderModel struct {
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	CategoryID        int              `json:"category_id"`
	Visibility        int              `json:"visibility"`
	DepartmentIDs     []int            `json:"department_ids,omitempty"`
	GroupIDs          []int            `json:"group_ids,omitempty"`
	RequesterGroupIDs []int            `json:"requester_group_ids,omitempty"`
	ManageByGroupIDs  []int            `json:"manage_by_group_ids"`
	ApprovalSettings  ApprovalSettings `json:"approval_settings,omitempty"`
}

// UpdateSolutionFolderModel is a data struct for updating a SolutionFolder
type UpdateSolutionFolderModel struct {
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	Visibility        int              `json:"visibility"`
	DepartmentIDs     []int            `json:"department_ids,omitempty"`
	GroupIDs          []int            `json:"group_ids,omitempty"`
	RequesterGroupIDs []int            `json:"requester_group_ids,omitempty"`
	ManageByGroupIDs  []int            `json:"manage_by_group_ids"`
	ApprovalSettings  ApprovalSettings `json:"approval_settings,omitempty"`
}

// ListSolutionFoldersOptions represents filters/pagination for SolutionFolders
type ListSolutionFoldersOptions struct {
	ListOptions
	CategoryID int `json:"category_id"`
}

// GetSolutionFolder will return a SolutionFolder by id
func (s *SolutionService) GetSolutionFolder(id int) (*SolutionFolder, *http.Response, error) {
	o := new(solutionFolderWrapper)
	res, err := s.client.Get(fmt.Sprintf(solutionFolderIdUrl, id), &o)
	return &o.Details, res, err
}

// ListSolutionFolders will return paginated/filtered SolutionFolders using ListSolutionFoldersOptions
func (s *SolutionService) ListSolutionFolders(opt *ListSolutionFoldersOptions) (*SolutionFolders, *http.Response, error) {
	o := new(SolutionFolders)
	res, err := s.client.List(solutionFoldersUrl, opt, &o)
	return o, res, err
}

// CreateSolutionFolder will create and return a new SolutionFolder based on CreateSolutionFolderModel
func (s *SolutionService) CreateSolutionFolder(solutionFolder *CreateSolutionFolderModel) (*SolutionFolder, *http.Response, error) {
	o := new(solutionFolderWrapper)
	res, err := s.client.Post(solutionFoldersUrl, solutionFolder, &o)
	return &o.Details, res, err
}

// UpdateSolutionFolder will update and return a SolutionFolder matching id based UpdateSolutionFolderModel
func (s *SolutionService) UpdateSolutionFolder(id int, solutionFolder *UpdateSolutionFolderModel) (*SolutionFolder, *http.Response, error) {
	o := new(solutionFolderWrapper)
	res, err := s.client.Put(fmt.Sprintf(solutionFolderIdUrl, id), solutionFolder, &o)
	return &o.Details, res, err
}

// DeleteSolutionFolder will completely remove a SolutionFolder from FreshService matching id
func (s *SolutionService) DeleteSolutionFolder(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(solutionFolderIdUrl, id))
	return success, res, err
}
