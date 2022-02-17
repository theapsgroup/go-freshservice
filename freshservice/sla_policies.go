package freshservice

import (
	"net/http"
	"time"
)

const (
	slaUrl = "sla_policies"
)

// SLAPoliciesService API Docs: https://api.freshservice.com/#sla-policies
type SLAPoliciesService struct {
	client *Client
}

// Policies contains Collection an array of Policy
type Policies struct {
	Collection []Policy `json:"sla_policies"`
}

// Policy represents an SLA Policy in FreshService
type Policy struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Position    int         `json:"position"`
	IsDefault   bool        `json:"is_default"`
	Active      bool        `json:"active"`
	Deleted     bool        `json:"deleted"`
	Description string      `json:"description"`
	Targets     []SLATarget `json:"sla_targets"`
	Applicable  Applicable  `json:"applicable_to"`
	Escalation  Escalation  `json:"escalation"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type SLATarget struct {
	Priority          int  `json:"priority"`
	EscalationEnabled bool `json:"escalation_enabled"`
	RespondWithin     int  `json:"respond_within"`
	ResolveWithin     int  `json:"resolve_within"`
	BusinessHours     bool `json:"business_hours"`
}

type Applicable struct {
	TicketType        []string `json:"ticket_type"`
	ServiceItems      []int    `json:"service_items"`
	ServiceCategories []int    `json:"service_categories"`
	DepartmentIDs     []int    `json:"department_id"`
	GroupIDs          []int    `json:"group_id"`
	Category          string   `json:"category"`
	SubCategory       string   `json:"sub_category"`
	ItemCategory      string   `json:"item_category"`
	Source            []int    `json:"source"`
}

type Escalation struct {
	Response   EscalationDetails   `json:"response"`
	Resolution []EscalationDetails `json:"resolution"`
}

type EscalationDetails struct {
	Level          string `json:"level"`
	EscalationWhen string `json:"escalation_when"`
	EscalationTime int    `json:"escalation_time"`
	AgentIDs       []int  `json:"agent_ids"`
	GroupIDs       []int  `json:"group_ids"`
}

// ListPolicies will return SLA Policies
func (s *SLAPoliciesService) ListPolicies() (*Policies, *http.Response, error) {
	o := new(Policies)
	res, err := s.client.List(slaUrl, nil, &o)
	return o, res, err
}
