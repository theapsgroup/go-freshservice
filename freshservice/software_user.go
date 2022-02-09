package freshservice

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// SoftwareUsers contains Collection an array of SoftwareUser
type SoftwareUsers struct {
	Collection []SoftwareUser `json:"application_users"`
}

// SoftwareUserBindings contains Collection an array of SoftwareUserBindingModel
type SoftwareUserBindings struct {
	Collection []SoftwareUserBindingModel `json:"application_users"`
}

// softwareUserWrapper contains Details of an SoftwareUser
type softwareUserWrapper struct {
	Details SoftwareUser `json:"application_user"`
}

// SoftwareUser represents a binding between a User (Requester or Agent) and an Application
type SoftwareUser struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	LicenseID     int       `json:"license_id"`
	AllocatedDate time.Time `json:"allocated_date"`
	FirstUsed     time.Time `json:"first_used"`
	LastUsed      time.Time `json:"last_used"`
	Sources       []string  `json:"sources"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type SoftwareUserBindingModel struct {
	UserID        int       `json:"user_id"`
	LicenseID     int       `json:"license_id"`
	AllocatedDate time.Time `json:"allocated_date"`
	FirstUsed     time.Time `json:"first_used"`
	LastUsed      time.Time `json:"last_used"`
	Sources       []string  `json:"sources"`
}

// ListSoftwareUsersOptions represents filters/pagination for SoftwareUsers
type ListSoftwareUsersOptions struct {
	ListOptions
}

// GetSoftwareUser will return an SoftwareUser by id
func (s *SoftwareService) GetSoftwareUser(applicationId int, id int) (*SoftwareUser, *http.Response, error) {
	o := new(softwareUserWrapper)
	res, err := s.client.Get(fmt.Sprintf(applicationUsersIdUrl, applicationId, id), &o)
	return &o.Details, res, err
}

// ListSoftwareUsers will return paginated/filtered SoftwareUsers using ListSoftwareUsersOptions
func (s *SoftwareService) ListSoftwareUsers(applicationId int, opt *ListSoftwareUsersOptions) (*SoftwareUsers, *http.Response, error) {
	o := new(SoftwareUsers)
	res, err := s.client.List(fmt.Sprintf(applicationUsersUrl, applicationId), opt, &o)
	return o, res, err
}

// BulkAddUsers allows for adding many SoftwareUser records to an Application as a bulk operation, returns SoftwareUsers
func (s *SoftwareService) BulkAddUsers(applicationId int, userBindings *SoftwareUserBindings) (*SoftwareUsers, *http.Response, error) {
	o := new(SoftwareUsers)
	res, err := s.client.Post(fmt.Sprintf(applicationUsersUrl, applicationId), userBindings, &o)
	return o, res, err
}

// BulkUpdateUsers allows for updating many SoftwareUser records of an Application as a bulk operation, returns SoftwareUsers
func (s *SoftwareService) BulkUpdateUsers(applicationId int, userBindings *SoftwareUserBindings) (*SoftwareUsers, *http.Response, error) {
	o := new(SoftwareUsers)
	res, err := s.client.Put(fmt.Sprintf(applicationUsersUrl, applicationId), userBindings, &o)
	return o, res, err
}

// DeleteUsers allows for bulk removal of Users (Requesters or Agent)
func (s *SoftwareService) DeleteUsers(applicationId int, userIds []string) (bool, *http.Response, error) {
	path := fmt.Sprintf(applicationIdUrl, applicationId)
	q := strings.Join(userIds, ",")
	success, res, err := s.client.Delete(fmt.Sprintf("%s?user_ids=%s", path, q))
	return success, res, err
}
