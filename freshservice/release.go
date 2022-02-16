package freshservice

import (
    "fmt"
    "net/http"
    "time"
)

const (
    releasesUrl       = "releases"
    releaseIdUrl      = "releases/%d"
    releaseRestoreUrl = "releases/%d/restore"
)

const (
    ReleasePriorityLow      = 1
    ReleasePriorityMedium   = 2
    ReleasePriorityHigh     = 3
    ReleasePriorityUrgent   = 4
    ReleaseStatusOpen       = 1
    ReleaseStatusOnHold     = 2
    ReleaseStatusInProgress = 3
    ReleaseStatusIncomplete = 4
    ReleaseStatusCompleted  = 5
    ReleaseTypeMinor        = 1
    ReleaseTypeStandard     = 2
    ReleaseTypeMajor        = 3
    ReleaseTypeEmergency    = 4
)

// ReleaseService API Docs: https://api.freshservice.com/#releases
type ReleaseService struct {
    client *Client
}

// Releases contains Collection an array of Release
type Releases struct {
    Collection []Release `json:"releases"`
}

// releaseWrapper contains Details of a Release
type releaseWrapper struct {
    Details Release `json:"release"`
}

// Release represents a Release in FreshService
type Release struct {
    ID                int       `json:"id"`
    AgentID           int       `json:"agent_id"`
    GroupID           int       `json:"group_id"`
    Priority          int       `json:"priority"`
    Status            int       `json:"status"`
    ReleaseType       int       `json:"release_type"`
    Subject           string    `json:"subject"`
    Description       string    `json:"description"`
    PlannedStartDate  time.Time `json:"planned_start_date"`
    PlannedEndDate    time.Time `json:"planned_end_date"`
    WorkStartDate     time.Time `json:"work_start_date"`
    WorkEndDate       time.Time `json:"work_end_date"`
    DepartmentID      int       `json:"department_id"`
    Category          string    `json:"category"`
    SubCategory       string    `json:"sub_category"`
    ItemCategory      string    `json:"item_category"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
    AssociatedAssets  []int     `json:"associated_assets"`
    AssociatedChanges []int     `json:"associated_changes"`
}

// CreateReleaseModel is a data struct for creating a new Release
type CreateReleaseModel struct {
    AgentID          int       `json:"agent_id"`
    GroupID          int       `json:"group_id"`
    Priority         int       `json:"priority"`
    Status           int       `json:"status"`
    ReleaseType      int       `json:"release_type"`
    Subject          string    `json:"subject"`
    Description      string    `json:"description"`
    PlannedStartDate time.Time `json:"planned_start_date"`
    PlannedEndDate   time.Time `json:"planned_end_date"`
    DepartmentID     int       `json:"department_id"`
    Category         string    `json:"category"`
    SubCategory      string    `json:"sub_category"`
    ItemCategory     string    `json:"item_category"`
}

// UpdateReleaseModel is a data struct for updating a Release
type UpdateReleaseModel struct {
    AgentID          int       `json:"agent_id"`
    GroupID          int       `json:"group_id"`
    Priority         int       `json:"priority"`
    Status           int       `json:"status"`
    ReleaseType      int       `json:"release_type"`
    Subject          string    `json:"subject"`
    Description      string    `json:"description"`
    PlannedStartDate time.Time `json:"planned_start_date"`
    PlannedEndDate   time.Time `json:"planned_end_date"`
    WorkStartDate    time.Time `json:"work_start_date"`
    WorkEndDate      time.Time `json:"work_end_date"`
    DepartmentID     int       `json:"department_id"`
    Category         string    `json:"category"`
    SubCategory      string    `json:"sub_category"`
    ItemCategory     string    `json:"item_category"`
}

// ListReleasesOptions represents filters/pagination for Releases
type ListReleasesOptions struct {
    ListOptions
    FilterName *string `json:"filter_name,omitempty" url:"filter_name,omitempty"`
}

// GetRelease will return a single Release by id
func (s *ReleaseService) GetRelease(id int) (*Release, *http.Response, error) {
    o := new(releaseWrapper)
    res, err := s.client.Get(fmt.Sprintf(releaseIdUrl, id), &o)
    return &o.Details, res, err
}

// ListReleases will return paginated/filtered Release using ListReleasesOptions
func (s *ReleaseService) ListReleases(opt *ListReleasesOptions) (*Releases, *http.Response, error) {
    o := new(Releases)
    res, err := s.client.List(releasesUrl, opt, &o)
    return o, res, err
}

// CreateRelease will create and return a new Release based on CreateReleaseModel
func (s *ReleaseService) CreateRelease(newRelease *CreateReleaseModel) (*Release, *http.Response, error) {
    o := new(releaseWrapper)
    res, err := s.client.Post(releasesUrl, newRelease, &o)
    return &o.Details, res, err
}

// UpdateRelease will update and return a Release matching id based on UpdateReleaseModel
func (s *ReleaseService) UpdateRelease(id int, ticket *UpdateReleaseModel) (*Release, *http.Response, error) {
    o := new(releaseWrapper)
    res, err := s.client.Put(fmt.Sprintf(releaseIdUrl, id), ticket, &o)
    return &o.Details, res, err
}

// DeleteRelease will delete a Release from FreshService (Can be restored by RestoreRelease)
func (s *ReleaseService) DeleteRelease(id int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(releaseIdUrl, id))
    return success, res, err
}

// RestoreRelease will restore a previously trashed (deleted) Release
func (s *ReleaseService) RestoreRelease(id int) (bool, *http.Response, error) {
    res, err := s.client.Put(fmt.Sprintf(releaseRestoreUrl, id), nil, nil)
    success, _ := isSuccessful(res)
    return success, res, err
}
