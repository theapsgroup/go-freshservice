package freshservice

import (
	"fmt"
	"net/http"
)

// GetTimeEntry will return a single TimeEntry for the specified Ticket
func (s *TicketService) GetTimeEntry(ticketId int, timeEntryId int) (*TimeEntry, *http.Response, error) {
	o := new(timeEntryWrapper)
	res, err := s.client.Get(fmt.Sprintf(ticketTimeEntryIdUrl, ticketId, timeEntryId), &o)
	return &o.Details, res, err
}

// ListTimeEntries will return TimeEntries for the specified Ticket
func (s *TicketService) ListTimeEntries(ticketId int) (*TimeEntries, *http.Response, error) {
	o := new(TimeEntries)
	res, err := s.client.List(fmt.Sprintf(ticketTimeEntryUrl, ticketId), nil, &o)
	return o, res, err
}

// CreateTimeEntry will create and return a new TimeEntry for the corresponding Ticket by ticketId based on CreateTimeEntryModel
func (s *TicketService) CreateTimeEntry(ticketId int, timeEntry *CreateTimeEntryModel) (*TimeEntry, *http.Response, error) {
	o := new(timeEntryWrapper)
	i := createTimeEntryWrapper{
		Data: *timeEntry,
	}
	res, err := s.client.Post(fmt.Sprintf(ticketTimeEntryUrl, ticketId), &i, &o)
	return &o.Details, res, err
}

// DeleteTimeEntry will completely remove a TimeEntry from a Ticket
func (s *TicketService) DeleteTimeEntry(ticketId int, timeEntryId int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(ticketTimeEntryIdUrl, ticketId, timeEntryId))
	return success, res, err
}
