package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	agentsUrl          = "agents"
	agentIdUrl         = "agents/%d"
	agentForgetUrl     = "agents/%d/forget"
	agentReactivateUrl = "agents/%d/reactivate"
)

// AgentService API Docs: https://api.freshservice.com/#agents https://api.freshservice.com/#agent-roles https://api.freshservice.com/#agent-groups
type AgentService struct {
	client *Client
}

// Agents contains Collection an array of Agent
type Agents struct {
	Collection []Agent `json:"agents"`
}

// agentWrapper contains Details of one Agent
type agentWrapper struct {
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

// CreateAgentModel is a data struct for creating a new Agent
type CreateAgentModel struct {
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

// UpdateAgentModel ris the data struct required to update an Agent
type UpdateAgentModel struct {
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

// ListAgentsOptions represents filters/pagination for Agents
type ListAgentsOptions struct {
	ListOptions
	Email  *string `json:"email,omitempty" url:"email,omitempty"`
	Active *bool   `json:"active,omitempty" url:"active,omitempty"`
	State  *string `json:"state,omitempty" url:"state,omitempty"`
}

// GetAgent will return a single Agent by id
func (s *AgentService) GetAgent(id int) (*Agent, *http.Response, error) {
	o := new(agentWrapper)
	res, err := s.client.Get(fmt.Sprintf(agentIdUrl, id), &o)
	return &o.Details, res, err
}

// ListAgents will return paginated/filtered Agents using ListAgentsOptions
func (s *AgentService) ListAgents(opt *ListAgentsOptions) (*Agents, *http.Response, error) {
	o := new(Agents)
	res, err := s.client.List(agentsUrl, opt, &o)
	return o, res, err
}

// CreateAgent will create and return a new Agent based on CreateAgentModel
func (s *AgentService) CreateAgent(newAgent *CreateAgentModel) (*Agent, *http.Response, error) {
	o := new(agentWrapper)
	res, err := s.client.Post(agentsUrl, newAgent, &o)
	return &o.Details, res, err
}

// UpdateAgent will update and return an Agent matching id based on UpdateAgentModel
func (s *AgentService) UpdateAgent(id int, agent *UpdateAgentModel) (*Agent, *http.Response, error) {
	o := new(agentWrapper)
	res, err := s.client.Put(fmt.Sprintf(agentIdUrl, id), agent, &o)
	return &o.Details, res, err
}

// DeleteAgent will completely remove an Agent from FreshService matching id (along with their requested Tickets)
func (s *AgentService) DeleteAgent(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(agentForgetUrl, id))
	return success, res, err
}

// DeactivateAgent will deactivate the Agent matching the id
func (s *AgentService) DeactivateAgent(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(agentIdUrl, id))
	return success, res, err
}

// ReactivateAgent will reactivate a deactivated Agent matching the id
func (s *AgentService) ReactivateAgent(id int) (*Agent, *http.Response, error) {
	o := new(agentWrapper)
	res, err := s.client.Put(fmt.Sprintf(agentReactivateUrl, id), nil, &o)
	return &o.Details, res, err
}
