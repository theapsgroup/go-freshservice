package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// Conversations contains Collection an array of Conversation
type Conversations struct {
	Collection []Conversation `json:"conversations"`
}

// Conversation represents a Conversation / Discussion on a Ticket
type Conversation struct {
	ID           int                `json:"id"`
	Attachments  []TicketAttachment `json:"attachments"`
	Body         string             `json:"body"`
	BodyText     string             `json:"body_text"`
	Incoming     bool               `json:"incoming"`
	ToEmails     []string           `json:"to_emails"`
	Private      bool               `json:"private"`
	Source       int                `json:"source"`
	SupportEmail string             `json:"support_email"`
	TicketID     int                `json:"ticket_id"`
	UserID       int                `json:"user_id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

// ListConversationsOptions represents filters/pagination for Conversations
type ListConversationsOptions struct {
	ListOptions
}

// ListConversations will return paginated/filtered Conversation using ListConversationsOptions
func (s *TicketService) ListConversations(ticketId int, opt *ListConversationsOptions) (*Conversations, *http.Response, error) {
	o := new(Conversations)
	res, err := s.client.List(fmt.Sprintf(ticketConversationsUrl, ticketId), opt, &o)
	return o, res, err
}
