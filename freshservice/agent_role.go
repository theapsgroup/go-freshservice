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
	o := new(agentRoleWrapper)
	res, err := s.client.Get(fmt.Sprintf(agentRoleIdUrl, id), &o)
	return &o.Details, res, err
}

// ListAgentRoles will return paginated/filtered AgentRoles using ListAgentRolesOptions
func (s *AgentService) ListAgentRoles(opt ListAgentRolesOptions) (*AgentRoles, *http.Response, error) {
	o := new(AgentRoles)
	res, err := s.client.List(agentRolesUrl, opt, &o)
	return o, res, err
}
