package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	ticketsUrl                = "tickets"
	ticketIdUrl               = "tickets/%d"
	ticketRestoreUrl          = "tickets/%d/restore"
	ticketRemoveAttachmentUrl = "tickets/%d/attachments/%d"
	ticketActivitiesUrl       = "tickets/%d/activities"
	ticketTimeEntryUrl        = "tickets/%d/time_entries"
	ticketTimeEntryIdUrl      = "tickets/%d/time_entries/%d"
	ticketTasksUrl            = "tickets/%d/tasks"
	ticketTaskIdUrl           = "tickets/%d/tasks/%d"
	ticketConversationsUrl    = "tickets/%d/conversations"
)

const (
	TicketOpen          = 2
	TicketPending       = 3
	TicketResolved      = 4
	TicketClosed        = 5
	SourceEmail         = 1
	SourcePortal        = 2
	SourcePhone         = 3
	SourceChat          = 4
	SourceWidget        = 5
	SourceYammer        = 6
	SourceAwsCloudWatch = 7
	SourcePagerDuty     = 8
	SourceWalkUp        = 9
	SourceSlack         = 10
	PriorityLow         = 1
	PriorityMedium      = 2
	PriorityHigh        = 3
	PriorityUrgent      = 4
)

// TicketService API Docs: https://api.freshservice.com/#tickets
type TicketService struct {
	client *Client
}

// Tickets contains Collection an array of Ticket
type Tickets struct {
	Collection []Ticket `json:"tickets"`
}

// ticketWrapper contains Details of one Ticket
type ticketWrapper struct {
	Details Ticket `json:"ticket,omitempty"`
}

// Ticket represents a FreshService Ticket
type Ticket struct {
	ID                     int                `json:"id"`
	Attachments            []TicketAttachment `json:"attachments"`
	CcEmails               []string           `json:"cc_emails"`
	DepartmentID           int                `json:"department_id"`
	Deleted                bool               `json:"deleted"`
	Description            string             `json:"description"`
	DescriptionText        string             `json:"description_text"`
	DueBy                  time.Time          `json:"due_by"`
	Email                  string             `json:"email"`
	EmailConfigID          int                `json:"email_config_id"`
	FirstResponseDueBy     time.Time          `json:"fr_due_by"`
	FirstResponseEscalated bool               `json:"fr_escalated"`
	ForwardEmails          []string           `json:"fwd_emails"`
	GroupID                int                `json:"group_id"`
	IsEscalated            bool               `json:"is_escalated"`
	Name                   string             `json:"name"`
	Phone                  string             `json:"phone"`
	Priority               int                `json:"priority"`
	Category               int                `json:"category"`
	SubCategory            string             `json:"sub_category"`
	ItemCategory           string             `json:"item_category"`
	ReplyCcEmails          []string           `json:"reply_cc_emails"`
	RequesterID            int                `json:"requester_id"`
	ResponderID            int                `json:"responder_id"`
	Source                 int                `json:"source"`
	Spam                   bool               `json:"spam"`
	Status                 int                `json:"status"`
	Subject                string             `json:"subject"`
	Tags                   []string           `json:"tags"`
	ToEmails               []string           `json:"to_emails"`
	Type                   string             `json:"type"`
	Urgency                int                `json:"urgency"`
	Impact                 int                `json:"impact"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
}

// CreateTicketModel is a data struct for creating a new Ticket
type CreateTicketModel struct {
	Attachments        []TicketAttachment `json:"attachments"`
	CcEmails           []string           `json:"cc_emails"`
	DepartmentID       int                `json:"department_id"`
	Description        string             `json:"description"`
	DueBy              time.Time          `json:"due_by"`
	Email              string             `json:"email"`
	EmailConfigID      int                `json:"email_config_id"`
	FirstResponseDueBy time.Time          `json:"fr_due_by"`
	GroupID            int                `json:"group_id"`
	Name               string             `json:"name"`
	Phone              string             `json:"phone"`
	Priority           int                `json:"priority"`
	Category           int                `json:"category"`
	SubCategory        string             `json:"sub_category"`
	ItemCategory       string             `json:"item_category"`
	RequesterID        int                `json:"requester_id"`
	ResponderID        int                `json:"responder_id"`
	Source             int                `json:"source"`
	Status             int                `json:"status"`
	Subject            string             `json:"subject"`
	Tags               []string           `json:"tags"`
	Type               string             `json:"type"`
	Urgency            int                `json:"urgency"`
	Impact             int                `json:"impact"`
}

// UpdateTicketModel is a data struct for updating a Ticket
type UpdateTicketModel struct {
	Attachments        []TicketAttachment `json:"attachments"`
	DepartmentID       int                `json:"department_id"`
	Description        string             `json:"description"`
	DueBy              time.Time          `json:"due_by"`
	Email              string             `json:"email"`
	EmailConfigID      int                `json:"email_config_id"`
	FirstResponseDueBy time.Time          `json:"fr_due_by"`
	GroupID            int                `json:"group_id"`
	Name               string             `json:"name"`
	Phone              string             `json:"phone"`
	Priority           int                `json:"priority"`
	Category           int                `json:"category"`
	SubCategory        string             `json:"sub_category"`
	ItemCategory       string             `json:"item_category"`
	RequesterID        int                `json:"requester_id"`
	ResponderID        int                `json:"responder_id"`
	Source             int                `json:"source"`
	Status             int                `json:"status"`
	Subject            string             `json:"subject"`
	Tags               []string           `json:"tags"`
	Type               string             `json:"type"`
	Urgency            int                `json:"urgency"`
	Impact             int                `json:"impact"`
}

// TicketAttachment represents an Attachment on a Ticket
type TicketAttachment struct {
	Name          string    `json:"name"`
	Size          int       `json:"size"`
	ContentType   string    `json:"content_type"`
	AttachmentUrl string    `json:"attachment_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// TicketActivities contains Collection an array of TicketActivity
type TicketActivities struct {
	Collection []TicketActivity `json:"activities"`
}

// TicketActivity represents an Audit Item on a Ticket
type TicketActivity struct {
	Actor       Actor     `json:"actor"`
	Content     string    `json:"content"`
	SubContents string    `json:"sub_contents"`
	CreatedAt   time.Time `json:"created_at"`
}

// ListTicketsOptions represents filters/pagination for Tickets
type ListTicketsOptions struct {
	ListOptions
	Email        *string    `json:"email,omitempty" url:"email,omitempty"`
	RequesterID  *int       `json:"requester_id,omitempty" url:"requester_id,omitempty"`
	UpdatedSince *time.Time `json:"updated_since,omitempty" url:"updated_since,omitempty"`
	Type         *string    `json:"type,omitempty" url:"type,omitempty"`
}

// GetTicket will return a single Ticket by id
func (s *TicketService) GetTicket(id int) (*Ticket, *http.Response, error) {
	o := new(ticketWrapper)
	res, err := s.client.Get(fmt.Sprintf(ticketIdUrl, id), &o)
	return &o.Details, res, err
}

// ListTickets will return paginated/filtered Ticket using ListTicketsOptions
func (s *TicketService) ListTickets(opt *ListTicketsOptions) (*Tickets, *http.Response, error) {
	o := new(Tickets)
	res, err := s.client.List(ticketsUrl, opt, &o)
	return o, res, err
}

// CreateTicket will create and return a new Ticket based on CreateTicketModel
func (s *TicketService) CreateTicket(newTicket *CreateAgentModel) (*Ticket, *http.Response, error) {
	o := new(ticketWrapper)
	res, err := s.client.Post(ticketsUrl, newTicket, &o)
	return &o.Details, res, err
}

// UpdateTicket will update and return a Ticket matching id based on UpdateTicketModel
func (s *TicketService) UpdateTicket(id int, ticket *UpdateTicketModel) (*Ticket, *http.Response, error) {
	o := new(ticketWrapper)
	res, err := s.client.Put(fmt.Sprintf(ticketIdUrl, id), ticket, &o)
	return &o.Details, res, err
}

// DeleteTicket will trash a Ticket from FreshService (Can be restored by RestoreTicket)
func (s *TicketService) DeleteTicket(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(ticketIdUrl, id))
	return success, res, err
}

// RestoreTicket will restore a previously trashed (deleted) Ticket
func (s *TicketService) RestoreTicket(id int) (bool, *http.Response, error) {
	res, err := s.client.Put(fmt.Sprintf(ticketRestoreUrl, id), nil, nil)
	success, _ := isSuccessful(res)
	return success, res, err
}

// DeleteAttachment will remove a TicketAttachment from a Ticket
func (s *TicketService) DeleteAttachment(ticketId int, attachmentId int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(ticketRemoveAttachmentUrl, ticketId, attachmentId))
	return success, res, err
}

// GetAudit returns TicketActivities for a specific Ticket
func (s *TicketService) GetAudit(ticketId int) (*TicketActivities, *http.Response, error) {
	o := new(TicketActivities)
	res, err := s.client.List(fmt.Sprintf(ticketActivitiesUrl, ticketId), nil, &o)
	return o, res, err
}
