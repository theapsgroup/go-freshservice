package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	applicationsUrl             = "applications"
	applicationIdUrl            = "applications/%d"
	applicationUsersUrl         = "applications/%d/users"
	applicationUsersIdUrl       = "applications/%d/users/%d"
	applicationInstallationsUrl = "applications/%d/installations"
)

// SoftwareService API Docs: https://api.freshservice.com/#software
type SoftwareService struct {
	client *Client
}

// Applications contains Collection an array of Application
type Applications struct {
	Collection []Application `json:"applications"`
}

// applicationWrapper contains Details of an Application
type applicationWrapper struct {
	Details Application `json:"application"`
}

// Application represents an Application (Software) registered in FreshService
type Application struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	ApplicationType   string    `json:"application_type"`
	Status            string    `json:"status"`
	PublisherID       int       `json:"publisher_id"`
	ManagedByID       int       `json:"managed_by_id"`
	Notes             string    `json:"notes"`
	Category          string    `json:"category"`
	Sources           []string  `json:"sources"`
	UserCount         int       `json:"user_count"`
	InstallationCount int       `json:"installation_count"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// CreateApplicationModel is a data struct for creating a new Application
type CreateApplicationModel struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	ApplicationType string   `json:"application_type"`
	Status          string   `json:"status"`
	PublisherID     int      `json:"publisher_id"`
	ManagedByID     int      `json:"managed_by_id"`
	Notes           string   `json:"notes"`
	Category        string   `json:"category"`
	Sources         []string `json:"sources"`
}

// UpdateApplicationModel is a data struct for updating an Application
type UpdateApplicationModel struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	ApplicationType string   `json:"application_type"`
	Status          string   `json:"status"`
	PublisherID     int      `json:"publisher_id"`
	ManagedByID     int      `json:"managed_by_id"`
	Notes           string   `json:"notes"`
	Category        string   `json:"category"`
	Sources         []string `json:"sources"`
}

// ListApplicationsOptions represents filters/pagination for Applications
type ListApplicationsOptions struct {
	ListOptions
}

// GetApplication will return an Application by id
func (s *SoftwareService) GetApplication(id int) (*Application, *http.Response, error) {
	o := new(applicationWrapper)
	res, err := s.client.Get(fmt.Sprintf(applicationIdUrl, id), &o)
	return &o.Details, res, err
}

// ListApplications will return paginated/filtered Applications using ListApplicationsOptions
func (s *SoftwareService) ListApplications(opt *ListApplicationsOptions) (*Applications, *http.Response, error) {
	o := new(Applications)
	res, err := s.client.List(applicationsUrl, opt, &o)
	return o, res, err
}

// CreateApplication will create and return a new Application based on CreateApplicationModel
func (s *SoftwareService) CreateApplication(newApplication *CreateApplicationModel) (*Application, *http.Response, error) {
	o := new(applicationWrapper)
	res, err := s.client.Post(applicationsUrl, newApplication, &o)
	return &o.Details, res, err
}

// UpdateApplication will update and return a Application matching id based UpdateApplicationModel
func (s *SoftwareService) UpdateApplication(id int, application *UpdateLocationModel) (*Application, *http.Response, error) {
	o := new(applicationWrapper)
	res, err := s.client.Put(fmt.Sprintf(applicationIdUrl, id), application, &o)
	return &o.Details, res, err
}

// DeleteApplication will completely remove an Application from FreshService matching id
func (s *SoftwareService) DeleteApplication(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(applicationIdUrl, id))
	return success, res, err
}
