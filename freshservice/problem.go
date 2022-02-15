package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	problemsUrl       = "problems"
	problemIdUrl      = "problems/%d"
	problemRestoreUrl = "problems/%d/restore"
)

// ProblemService API Docs: https://api.freshservice.com/#problems
type ProblemService struct {
	client *Client
}

// Problems contains Collection an array of Problem
type Problems struct {
	Collection []Problem `json:"problems"`
}

// problemWrapper contains Details of one Problem
type problemWrapper struct {
	Details Problem `json:"problem"`
}

// Problem represents a FreshService Problem
type Problem struct {
	ID               int             `json:"id"`
	AgentID          int             `json:"agent_id"`
	RequesterID      int             `json:"requester_id"`
	GroupID          int             `json:"group_id"`
	Description      string          `json:"description"`
	DescriptionText  string          `json:"description_text"`
	Priority         int             `json:"priority"`
	Status           int             `json:"status"`
	Impact           int             `json:"impact"`
	KnownError       bool            `json:"known_error"`
	Subject          string          `json:"subject"`
	DueBy            time.Time       `json:"due_by"`
	DepartmentID     int             `json:"department_id"`
	Category         string          `json:"category"`
	SubCategory      string          `json:"sub_category"`
	ItemCategory     string          `json:"item_category"`
	AssociatedChange int             `json:"associated_change"`
	AnalysisFields   ProblemAnalysis `json:"analysis_fields,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

// ProblemAnalysis is a data structure
type ProblemAnalysis struct {
	ProblemCause   Description `json:"problem_cause,omitempty"`
	ProblemSymptom Description `json:"problem_symptom,omitempty"`
	ProblemImpact  Description `json:"problem_impact,omitempty"`
}

// CreateProblemModel is the data structure required to create a new Problem
type CreateProblemModel struct {
	AgentID        int             `json:"agent_id"`
	GroupID        int             `json:"group_id"`
	Description    string          `json:"description"`
	Priority       int             `json:"priority"`
	Status         int             `json:"status"`
	Impact         int             `json:"impact"`
	Subject        string          `json:"subject"`
	DueBy          time.Time       `json:"due_by"`
	DepartmentID   int             `json:"department_id"`
	Category       string          `json:"category"`
	SubCategory    string          `json:"sub_category"`
	ItemCategory   string          `json:"item_category"`
	AnalysisFields ProblemAnalysis `json:"analysis_fields,omitempty"`
}

// UpdateProblemModel is the data structure required for updating a Problem
type UpdateProblemModel struct {
	AgentID        int             `json:"agent_id"`
	GroupID        int             `json:"group_id"`
	Description    string          `json:"description"`
	Priority       int             `json:"priority"`
	Status         int             `json:"status"`
	Impact         int             `json:"impact"`
	Subject        string          `json:"subject"`
	DueBy          time.Time       `json:"due_by"`
	DepartmentID   int             `json:"department_id"`
	Category       string          `json:"category"`
	SubCategory    string          `json:"sub_category"`
	ItemCategory   string          `json:"item_category"`
	AnalysisFields ProblemAnalysis `json:"analysis_fields,omitempty"`
}

// ListProblemsOptions represents filters/pagination for Problems
type ListProblemsOptions struct {
	ListOptions
}

// GetProblem will return a single Problem by id
func (s *ProblemService) GetProblem(id int) (*Problem, *http.Response, error) {
	o := new(problemWrapper)
	res, err := s.client.Get(fmt.Sprintf(problemIdUrl, id), &o)
	return &o.Details, res, err
}

// ListProblems will return paginated/filtered Problems using ListProblemsOptions
func (s *ProblemService) ListProblems(opt *ListProblemsOptions) (*Problems, *http.Response, error) {
	o := new(Problems)
	res, err := s.client.List(problemsUrl, opt, &o)
	return o, res, err
}

// CreateProblem will create and return a new Problem based on CreateProblemModel
func (s *ProblemService) CreateProblem(problem *CreateProblemModel) (*Problem, *http.Response, error) {
	o := new(problemWrapper)
	res, err := s.client.Post(assetsUrl, problem, &o)
	return &o.Details, res, err
}

// UpdateProblem will update and return a Problem matching id based on UpdateProblemModel
func (s *ProblemService) UpdateProblem(id int, problem *UpdateProblemModel) (*Problem, *http.Response, error) {
	o := new(problemWrapper)
	res, err := s.client.Put(fmt.Sprintf(problemIdUrl, id), problem, &o)
	return &o.Details, res, err
}

// DeleteProblem will delete a Problem matching the id (non-permanent delete)
func (s *ProblemService) DeleteProblem(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(problemIdUrl, id))
	return success, res, err
}

// RestoreProblem will restore a previously deleted Problem by id
func (s *ProblemService) RestoreProblem(id int) (bool, *http.Response, error) {
	res, err := s.client.Put(fmt.Sprintf(problemRestoreUrl, id), nil, nil)
	success, _ := isSuccessful(res)
	return success, res, err
}
