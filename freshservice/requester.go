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
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(requesterIdUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(requesterWrapper)
	res, err := s.client.SendRequest(req, &r)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &r.Details, res, nil
}

// ListRequesters will return paginated/filtered Requesters using ListRequestersOptions
func (s *RequesterService) ListRequesters(opt *ListRequestersOptions) (*Requesters, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, requestersUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	rs := new(Requesters)
	res, err := s.client.SendRequest(req, &rs)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return rs, res, nil
}

// CreateRequester will create and return a new Requester based on CreateRequesterModel
func (s *RequesterService) CreateRequester(newRequester *CreateRequesterModel) (*Requester, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, requestersUrl, newRequester)
	if err != nil {
		return nil, nil, err
	}

	r := new(requesterWrapper)
	res, err := s.client.SendRequest(req, &r)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &r.Details, res, nil
}

// UpdateRequester will update and return an Requester matching id based on UpdateRequesterModel
func (s *RequesterService) UpdateRequester(id int, requester *UpdateRequesterModel) (*Requester, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(requesterIdUrl, id), requester)
	if err != nil {
		return nil, nil, err
	}

	r := new(requesterWrapper)
	res, err := s.client.SendRequest(req, &r)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &r.Details, res, nil
}

// DeleteRequester will completely remove a Requester from FreshService matching id (along with their requested Tickets)
func (s *RequesterService) DeleteRequester(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(requesterForgetUrl, id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}

// DeactivateRequester will deactivate the Requester matching the id
func (s *RequesterService) DeactivateRequester(id int) (*Requester, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(requesterIdUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(requesterWrapper)
	res, err := s.client.SendRequest(req, &r)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &r.Details, res, nil
}

// ReactivateRequester will reactivate a deactivated Requester matching the id
func (s *RequesterService) ReactivateRequester(id int) (*Requester, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(requesterReactivateUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(requesterWrapper)
	res, err := s.client.SendRequest(req, &r)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &r.Details, res, nil
}
