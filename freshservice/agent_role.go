package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	agentRolesUrl  = "roles"
	agentRoleIdUrl = "roles/%d"
)

// AgentRoles contains Collection an array of AgentRole
type AgentRoles struct {
	Collection []AgentRole `json:"roles"`
}

// agentRoleWrapper contains Details of an AgentRole
type agentRoleWrapper struct {
	Details AgentRole `json:"role"`
}

// AgentRole represents a FreshService AgentRole
type AgentRole struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Default     bool      `json:"default"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ListAgentRolesOptions represents pagination/filtering for AgentRoles
type ListAgentRolesOptions struct {
	ListOptions
}

// GetAgentRole will return a single AgentRole by id
func (s *AgentService) GetAgentRole(id int) (*AgentRole, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(agentRoleIdUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(agentRoleWrapper)
	res, err := s.client.SendRequest(req, &ar)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &ar.Details, res, nil
}

// ListAgentRoles will return paginated/filtered AgentRoles using ListAgentRolesOptions
func (s *AgentService) ListAgentRoles(opt ListAgentRolesOptions) (*AgentRoles, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, agentRolesUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	ars := new(AgentRoles)
	res, err := s.client.SendRequest(req, &ars)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return ars, res, nil
}
