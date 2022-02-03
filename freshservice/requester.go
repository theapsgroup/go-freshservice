package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	requestersUrl          = "requesters"
	requesterIdUrl         = "requesters/%d"
	requesterForgetUrl     = "requesters/%d/forget"
	requesterReactivateUrl = "requesters/%d/reactivate"
)

// RequesterService API Docs: https://api.freshservice.com/#requesters https://api.freshservice.com/#requester-groups
type RequesterService struct {
	client *Client
}

// Requesters contains Collection an array of Requester
type Requesters struct {
	Collection []Requester `json:"requesters"`
}

// requesterWrapper contains Details of one Requester
type requesterWrapper struct {
	Details Requester `json:"requester"`
}

// Requester represents a FreshService Requester (User)
type Requester struct {
	ID                    int       `json:"id"`
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	JobTitle              string    `json:"job_title"`
	Email                 string    `json:"primary_email"`
	AdditionalEmails      []string  `json:"secondary_emails"`
	WorkPhoneNumber       string    `json:"work_phone_number"`
	MobilePhoneNumber     string    `json:"mobile_phone_number"`
	DepartmentIDs         []int     `json:"department_ids"`
	Active                bool      `json:"active"`
	Address               string    `json:"address"`
	ReportingManagerID    int       `json:"reporting_manager_id"`
	TimeZone              string    `json:"time_zone"`
	TimeFormat            string    `json:"time_format"`
	Language              string    `json:"language"`
	LocationID            int       `json:"location_id"`
	BackgroundInformation string    `json:"background_information"`
	HasLoggedIn           bool      `json:"has_logged_in"`
	IsAgent               bool      `json:"is_agent"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// CreateRequesterModel is a data struct for creating a new Requester
type CreateRequesterModel struct {
	FirstName             string   `json:"first_name"`
	LastName              string   `json:"last_name"`
	JobTitle              string   `json:"job_title"`
	Email                 string   `json:"primary_email"`
	AdditionalEmails      []string `json:"secondary_emails"`
	WorkPhoneNumber       string   `json:"work_phone_number"`
	MobilePhoneNumber     string   `json:"mobile_phone_number"`
	DepartmentIDs         []int    `json:"department_ids"`
	Address               string   `json:"address"`
	ReportingManagerID    int      `json:"reporting_manager_id"`
	TimeZone              string   `json:"time_zone"`
	TimeFormat            string   `json:"time_format"`
	Language              string   `json:"language"`
	LocationID            int      `json:"location_id"`
	BackgroundInformation string   `json:"background_information"`
}

// UpdateRequesterModel is a data struct for updating a Requester
type UpdateRequesterModel struct {
	FirstName             string   `json:"first_name"`
	LastName              string   `json:"last_name"`
	JobTitle              string   `json:"job_title"`
	Email                 string   `json:"primary_email"`
	AdditionalEmails      []string `json:"secondary_emails"`
	WorkPhoneNumber       string   `json:"work_phone_number"`
	MobilePhoneNumber     string   `json:"mobile_phone_number"`
	DepartmentIDs         []int    `json:"department_ids"`
	Address               string   `json:"address"`
	ReportingManagerID    int      `json:"reporting_manager_id"`
	TimeZone              string   `json:"time_zone"`
	TimeFormat            string   `json:"time_format"`
	Language              string   `json:"language"`
	LocationID            int      `json:"location_id"`
	BackgroundInformation string   `json:"background_information"`
}

// ListRequestersOptions represents filters/pagination for Requesters
type ListRequestersOptions struct {
	ListOptions
	Email         *string `json:"email,omitempty" url:"email,omitempty"`
	IncludeAgents *bool   `json:"include_agents,omitempty" url:"include_agents,omitempty"`
}

// GetRequester will return a single Requester by id
func (s *RequesterService) GetRequester(id int) (*Requester, *http.Response, error) {
	o := new(requesterWrapper)
	res, err := s.client.Get(fmt.Sprintf(requesterIdUrl, id), &o)
	return &o.Details, res, err
}

// ListRequesters will return paginated/filtered Requesters using ListRequestersOptions
func (s *RequesterService) ListRequesters(opt *ListRequestersOptions) (*Requesters, *http.Response, error) {
	o := new(Requesters)
	res, err := s.client.List(requestersUrl, opt, &o)
	return o, res, err
}

// CreateRequester will create and return a new Requester based on CreateRequesterModel
func (s *RequesterService) CreateRequester(newRequester *CreateRequesterModel) (*Requester, *http.Response, error) {
	o := new(requesterWrapper)
	res, err := s.client.Post(requestersUrl, newRequester, &o)
	return &o.Details, res, err
}

// UpdateRequester will update and return an Requester matching id based on UpdateRequesterModel
func (s *RequesterService) UpdateRequester(id int, requester *UpdateRequesterModel) (*Requester, *http.Response, error) {
	o := new(requesterWrapper)
	res, err := s.client.Put(fmt.Sprintf(requesterIdUrl, id), requester, &o)
	return &o.Details, res, err
}

// DeleteRequester will completely remove a Requester from FreshService matching id (along with their requested Tickets)
func (s *RequesterService) DeleteRequester(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(requesterForgetUrl, id))
	return success, res, err
}

// DeactivateRequester will deactivate the Requester matching the id
func (s *RequesterService) DeactivateRequester(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(requesterIdUrl, id))
	return success, res, err
}

// ReactivateRequester will reactivate a deactivated Requester matching the id
func (s *RequesterService) ReactivateRequester(id int) (*Requester, *http.Response, error) {
	o := new(requesterWrapper)
	res, err := s.client.Put(fmt.Sprintf(requesterReactivateUrl, id), nil, &o)
	return &o.Details, res, err
}
