package freshservice

import (
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
	ID                    int         `json:"id"`
	FirstName             string      `json:"first_name"`
	LastName              string      `json:"last_name"`
	Occasional            bool        `json:"occasional"`
	JobTitle              string      `json:"job_title"`
	Email                 string      `json:"email"`
	WorkPhoneNumber       string      `json:"work_phone_number"`
	MobilePhoneNumber     string      `json:"mobile_phone_number"`
	DepartmentIDs         []int       `json:"department_ids"`
	Active                bool        `json:"active"`
	Address               string      `json:"address"`
	ReportingManagerID    int         `json:"reporting_manager_id"`
	TimeZone              string      `json:"time_zone"`
	TimeFormat            string      `json:"time_format"`
	Language              string      `json:"language"`
	LocationID            int         `json:"location_id"`
	BackgroundInformation string      `json:"background_information"`
	ScoreboardLevelID     int         `json:"scoreboard_level_id"`
	MemberOf              []int       `json:"member_of"`
	ObserverOf            []int       `json:"observer_of"`
	Roles                 []AgentRole `json:"roles"`
	LastLoginAt           time.Time   `json:"last_login_at"`
	LastActiveAt          time.Time   `json:"last_active_at"`
	HasLoggedIn           bool        `json:"has_logged_in"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
}

type AgentRole struct {
	RoleID          int    `json:"role_id"`
	AssignmentScope string `json:"assignment_scope"`
	Groups          []int  `json:"groups"`
}
