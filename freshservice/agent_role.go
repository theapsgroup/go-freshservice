package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AgentRoles contains a Collection of AgentRole
type AgentRoles struct {
	Collection []AgentRole `json:"roles"`
}

// SpecificAgentRole contains Details of an AgentRole
type SpecificAgentRole struct {
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

// ListAgentRoleOptions represents pagination/filtering for AgentRoles
type ListAgentRoleOptions struct {
	ListOptions
}

// GetAgentRole will return a single AgentRole by id
func (s *AssetService) GetAgentRole(id int) (*AgentRole, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("roles/%v", id), nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(SpecificAgentRole)
	res, err := s.client.SendRequest(req, &ar)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &ar.Details, res, nil
}

// GetAgentRoles will return AssetRoles collection
func (s *AssetService) GetAgentRoles(opt ListAgentRoleOptions) (*AgentRoles, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "roles", opt)
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
