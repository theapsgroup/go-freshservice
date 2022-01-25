package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AgentService API Docs: https://api.freshservice.com/#agents https://api.freshservice.com/#agent-roles https://api.freshservice.com/#agent-groups
type AgentService struct {
	client *Client
}

// Agents contains collection of Agent
type Agents struct {
	Collection []Agent `json:"agents"`
}

// SpecificAgent contains Details of one specific Agent
type SpecificAgent struct {
	Details Agent `json:"agent"`
}

// Agent represents a FreshService Agent
type Agent struct {
	ID                    int                   `json:"id"`
	FirstName             string                `json:"first_name"`
	LastName              string                `json:"last_name"`
	Occasional            bool                  `json:"occasional"`
	JobTitle              string                `json:"job_title"`
	Email                 string                `json:"email"`
	WorkPhoneNumber       string                `json:"work_phone_number"`
	MobilePhoneNumber     string                `json:"mobile_phone_number"`
	DepartmentIDs         []int                 `json:"department_ids"`
	Active                bool                  `json:"active"`
	Address               string                `json:"address"`
	ReportingManagerID    int                   `json:"reporting_manager_id"`
	TimeZone              string                `json:"time_zone"`
	TimeFormat            string                `json:"time_format"`
	Language              string                `json:"language"`
	LocationID            int                   `json:"location_id"`
	BackgroundInformation string                `json:"background_information"`
	ScoreboardLevelID     int                   `json:"scoreboard_level_id"`
	MemberOf              []int                 `json:"member_of"`
	ObserverOf            []int                 `json:"observer_of"`
	Roles                 []AgentRoleAssignment `json:"roles"`
	LastLoginAt           time.Time             `json:"last_login_at"`
	LastActiveAt          time.Time             `json:"last_active_at"`
	HasLoggedIn           bool                  `json:"has_logged_in"`
	CreatedAt             time.Time             `json:"created_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
}

// NewAgent is a data struct for creating a new Agent
type NewAgent struct {
	FirstName             string                `json:"first_name"`
	LastName              string                `json:"last_name"`
	Occasional            bool                  `json:"occasional"`
	JobTitle              string                `json:"job_title"`
	Email                 string                `json:"email"`
	WorkPhoneNumber       string                `json:"work_phone_number"`
	MobilePhoneNumber     string                `json:"mobile_phone_number"`
	DepartmentIDs         []int                 `json:"department_ids"`
	Address               string                `json:"address"`
	ReportingManagerID    int                   `json:"reporting_manager_id"`
	TimeZone              string                `json:"time_zone"`
	TimeFormat            string                `json:"time_format"`
	Language              string                `json:"language"`
	LocationID            int                   `json:"location_id"`
	BackgroundInformation string                `json:"background_information"`
	ScoreboardLevelID     int                   `json:"scoreboard_level_id"`
	MemberOf              []int                 `json:"member_of"`
	ObserverOf            []int                 `json:"observer_of"`
	Roles                 []AgentRoleAssignment `json:"roles"`
}

// UpdateAgent ris the data struct required to update an Agent
type UpdateAgent struct {
	Occasional            bool                  `json:"occasional"`
	Email                 string                `json:"email"`
	DepartmentIDs         []int                 `json:"department_ids"`
	Address               string                `json:"address"`
	ReportingManagerID    int                   `json:"reporting_manager_id"`
	TimeZone              string                `json:"time_zone"`
	TimeFormat            string                `json:"time_format"`
	Language              string                `json:"language"`
	LocationID            int                   `json:"location_id"`
	BackgroundInformation string                `json:"background_information"`
	ScoreboardLevelID     int                   `json:"scoreboard_level_id"`
	MemberOf              []int                 `json:"member_of"`
	ObserverOf            []int                 `json:"observer_of"`
	Roles                 []AgentRoleAssignment `json:"roles"`
}

// AgentRoleAssignment represents a Role Assignment on an Agent
type AgentRoleAssignment struct {
	RoleID          int    `json:"role_id"`
	AssignmentScope string `json:"assignment_scope"`
	Groups          []int  `json:"groups"`
}

// ListAgentOptions represents query filters for Agents
type ListAgentOptions struct {
	ListOptions
	Email  *string `json:"email,omitempty" url:"email,omitempty"`
	Active *bool   `json:"active,omitempty" url:"active,omitempty"`
	State  *string `json:"state,omitempty" url:"state,omitempty"`
}

// GetAgent will return a single Agent by id, assuming a record is found.
func (s *AgentService) GetAgent(id int) (*Agent, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("agents/%v", id), nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAgent)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// GetAgents will return Agents collection, filter with ListAgentOptions
func (s *AgentService) GetAgents(opt *ListAgentOptions) (*Agents, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "agents", opt)
	if err != nil {
		return nil, nil, err
	}

	as := new(Agents)
	res, err := s.client.SendRequest(req, &as)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return as, res, nil
}

// CreateAgent will create a new Agent in FreshService
func (s *AgentService) CreateAgent(newAgent *NewAgent) (*Agent, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "agents", newAgent)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAgent)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// UpdateAgent will update the Agent matching the id and return the updated Agent
func (s *AgentService) UpdateAgent(id int, agent *UpdateAgent) (*Agent, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("agents/%d", id), agent)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAgent)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// DeleteAgent will completely remove an Agent from FreshService along with their requested Tickets, returns true if successful
func (s *AgentService) DeleteAgent(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("agents/%d/forget", id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}

// DeactivateAgent will deactivate the FreshService Agent
func (s *AgentService) DeactivateAgent(id int) (*Agent, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("agents/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAgent)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// ReactivateAgent will reactivate a deactivated FreshService Agent
func (s *AgentService) ReactivateAgent(id int) (*Agent, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("agents/%d/reactivate", id), nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAgent)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}
