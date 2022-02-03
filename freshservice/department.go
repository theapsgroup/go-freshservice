package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	departmentsUrl  = "departments"
	departmentIdUrl = "departments/%d"
)

// DepartmentService API Docs: https://api.freshservice.com/#departments
type DepartmentService struct {
	client *Client
}

// Departments contains Collection an array of Department
type Departments struct {
	Collection []Department `json:"departments"`
}

// departmentWrapper contains Details of one Department
type departmentWrapper struct {
	Details Department `json:"department"`
}

// Department represents a FreshService Department (Company/Team)
type Department struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	HeadUserId  int       `json:"head_user_id"`
	PrimeUserId int       `json:"prime_user_id"`
	Domains     []string  `json:"domains"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateDepartmentModel is the data structure required to create a new Department
type CreateDepartmentModel struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	HeadUserId  int      `json:"head_user_id"`
	PrimeUserId int      `json:"prime_user_id"`
	Domains     []string `json:"domains"`
}

// UpdateDepartmentModel is the data structure for updating a Department
type UpdateDepartmentModel struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	HeadUserId  int      `json:"head_user_id"`
	PrimeUserId int      `json:"prime_user_id"`
	Domains     []string `json:"domains"`
}

// ListDepartmentsOptions represents filters/pagination for Departments
type ListDepartmentsOptions struct {
	ListOptions
}

// GetDepartment will return a single Department by id
func (s *DepartmentService) GetDepartment(id int) (*Department, *http.Response, error) {
	o := new(departmentWrapper)
	res, err := s.client.Get(fmt.Sprintf(departmentIdUrl, id), &o)
	return &o.Details, res, err
}

// ListDepartments will return paginated/filtered Departments using ListDepartmentsOptions
func (s *DepartmentService) ListDepartments(opt *ListDepartmentsOptions) (*Departments, *http.Response, error) {
	o := new(Departments)
	res, err := s.client.List(departmentsUrl, opt, &o)
	return o, res, err
}

// CreateDepartment will create and return a new Department based on CreateDepartmentModel
func (s *DepartmentService) CreateDepartment(newDepartment *CreateDepartmentModel) (*Department, *http.Response, error) {
	o := new(departmentWrapper)
	res, err := s.client.Post(departmentsUrl, newDepartment, &o)
	return &o.Details, res, err
}

// UpdateDepartment will update and return a Department matching id based on UpdateDepartmentModel
func (s *DepartmentService) UpdateDepartment(id int, department *UpdateDepartmentModel) (*Department, *http.Response, error) {
	o := new(departmentWrapper)
	res, err := s.client.Put(fmt.Sprintf(departmentIdUrl, id), department, &o)
	return &o.Details, res, err
}

// DeleteDepartment will completely remove a Department from FreshService matching id
func (s *DepartmentService) DeleteDepartment(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(departmentIdUrl, id))
	return success, res, err
}
