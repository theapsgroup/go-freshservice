package freshservice

import (
    "fmt"
    "net/http"
)

// GetTimeEntry will return a single TimeEntry for the specified Problem
func (s *ProblemService) GetTimeEntry(problemId int, timeEntryId int) (*TimeEntry, *http.Response, error) {
    o := new(timeEntryWrapper)
    res, err := s.client.Get(fmt.Sprintf(problemTimeEntryIdUrl, problemId, timeEntryId), &o)
    return &o.Details, res, err
}

// ListTimeEntries will return TimeEntries for the specified Problem
func (s *ProblemService) ListTimeEntries(problemId int) (*TimeEntries, *http.Response, error) {
    o := new(TimeEntries)
    res, err := s.client.List(fmt.Sprintf(problemTimeEntryUrl, problemId), nil, &o)
    return o, res, err
}

// CreateTimeEntry will create and return a new TimeEntry for the corresponding Problem by problemId based on CreateTimeEntryModel
func (s *ProblemService) CreateTimeEntry(problemId int, timeEntry *CreateTimeEntryModel) (*TimeEntry, *http.Response, error) {
    o := new(timeEntryWrapper)
    i := createTimeEntryWrapper{
        Data: *timeEntry,
    }
    res, err := s.client.Post(fmt.Sprintf(problemTimeEntryUrl, problemId), &i, &o)
    return &o.Details, res, err
}

// DeleteTimeEntry will completely remove a TimeEntry from a Problem
func (s *ProblemService) DeleteTimeEntry(problemId int, timeEntryId int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(problemTimeEntryIdUrl, problemId, timeEntryId))
    return success, res, err
}
