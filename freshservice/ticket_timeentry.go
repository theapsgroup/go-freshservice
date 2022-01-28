package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// TimeEntry represents a TimeEntry associated to a Task / Ticket
type TimeEntry struct {
	ID           int       `json:"id"`
	StartTime    time.Time `json:"start_time"`
	ExecutedAt   time.Time `json:"executed_at"`
	TimerRunning bool      `json:"timer_running"`
	Billable     bool      `json:"billable"`
	TimeSpent    string    `json:"time_spent"`
	TaskID       int       `json:"task_id"`
	AgentID      int       `json:"agent_id"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// timeEntryWrapper contains Details of one TimeEntry
type timeEntryWrapper struct {
	Details TimeEntry `json:"time_entry"`
}

// TimeEntries contains Collection an array of TimeEntry
type TimeEntries struct {
	Collection []TimeEntry `json:"time_entries"`
}

// GetTimeEntry will return a single TimeEntry for the specified Ticket
func (s *TicketService) GetTimeEntry(ticketId int, timeEntryId int) (*TimeEntry, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(ticketTimeEntryIdUrl, ticketId, timeEntryId), nil)
	if err != nil {
		return nil, nil, err
	}

	te := new(timeEntryWrapper)
	res, err := s.client.SendRequest(req, &te)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &te.Details, res, nil
}

// ListTimeEntries will return TimeEntries for the specified Ticket
func (s *TicketService) ListTimeEntries(ticketId int) (*TimeEntries, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(ticketTimeEntryUrl, ticketId), nil)
	if err != nil {
		return nil, nil, err
	}

	tes := new(TimeEntries)
	res, err := s.client.SendRequest(req, &tes)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return tes, res, nil
}
