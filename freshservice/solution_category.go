package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// SolutionCategories contains Collection an array of SolutionCategory
type SolutionCategories struct {
	Collection []SolutionCategory `json:"categories"`
}

// solutionCategoryWrapper contains Details of one SolutionCategory
type solutionCategoryWrapper struct {
	Details SolutionCategory `json:"category"`
}

// SolutionCategory represents a FreshService SolutionCategory
type SolutionCategory struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Position         int       `json:"position"`
	DefaultCategory  bool      `json:"default_category"`
	VisibleInPortals []int     `json:"visible_in_portals"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CreateSolutionCategoryModel is the data structure required to create a new SolutionCategory
type CreateSolutionCategoryModel struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	VisibleInPortals []int  `json:"visible_in_portals"`
}

// UpdateSolutionCategoryModel is the data structure for updating a SolutionCategory
type UpdateSolutionCategoryModel struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	VisibleInPortals []int  `json:"visible_in_portals"`
}

// ListSolutionCategoriesOptions represents filters/pagination for SolutionCategories
type ListSolutionCategoriesOptions struct {
	ListOptions
}

// GetSolutionCategory will return a single SolutionCategory by id
func (s *SolutionService) GetSolutionCategory(id int) (*SolutionCategory, *http.Response, error) {
	o := new(solutionCategoryWrapper)
	res, err := s.client.Get(fmt.Sprintf(solutionCategoryIdUrl, id), &o)
	return &o.Details, res, err
}

// ListSolutionCategories will return paginated/filtered SolutionCategories using ListSolutionCategoriesOptions
func (s *SolutionService) ListSolutionCategories(opt *ListSolutionCategoriesOptions) (*SolutionCategories, *http.Response, error) {
	o := new(SolutionCategories)
	res, err := s.client.List(solutionCategoriesUrl, opt, &o)
	return o, res, err
}

// CreateSolutionCategory will create and return a new SolutionCategory based on CreateSolutionCategoryModel
func (s *SolutionService) CreateSolutionCategory(solutionCategory *CreateSolutionCategoryModel) (*SolutionCategory, *http.Response, error) {
	o := new(solutionCategoryWrapper)
	res, err := s.client.Post(solutionCategoriesUrl, solutionCategory, &o)
	return &o.Details, res, err
}

// UpdateSolutionCategory will update and return a SolutionCategory matching id based on UpdateSolutionCategoryModel
func (s *SolutionService) UpdateSolutionCategory(id int, solutionCategory *UpdateSolutionCategoryModel) (*SolutionCategory, *http.Response, error) {
	o := new(solutionCategoryWrapper)
	res, err := s.client.Put(fmt.Sprintf(solutionCategoryIdUrl, id), solutionCategory, &o)
	return &o.Details, res, err
}

// DeleteSolutionCategory will completely remove a SolutionCategory from FreshService matching id
func (s *SolutionService) DeleteSolutionCategory(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(solutionCategoryIdUrl, id))
	return success, res, err
}
