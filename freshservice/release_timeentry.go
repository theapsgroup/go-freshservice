package freshservice

import (
    "fmt"
    "net/http"
)

// GetTimeEntry will return a single TimeEntry for the specified Release
func (s *ReleaseService) GetTimeEntry(releaseId int, timeEntryId int) (*TimeEntry, *http.Response, error) {
    o := new(timeEntryWrapper)
    res, err := s.client.Get(fmt.Sprintf(releaseTimeEntryIdUrl, releaseId, timeEntryId), &o)
    return &o.Details, res, err
}

// ListTimeEntries will return TimeEntries for the specified Release
func (s *ReleaseService) ListTimeEntries(releaseId int) (*TimeEntries, *http.Response, error) {
    o := new(TimeEntries)
    res, err := s.client.List(fmt.Sprintf(releaseTimeEntryUrl, releaseId), nil, &o)
    return o, res, err
}

// CreateTimeEntry will create and return a new TimeEntry for the corresponding Release by releaseId based on CreateTimeEntryModel
func (s *ReleaseService) CreateTimeEntry(releaseId int, timeEntry *CreateTimeEntryModel) (*TimeEntry, *http.Response, error) {
    o := new(timeEntryWrapper)
    i := createTimeEntryWrapper{
        Data: *timeEntry,
    }
    res, err := s.client.Post(fmt.Sprintf(releaseTimeEntryUrl, releaseId), &i, &o)
    return &o.Details, res, err
}

// DeleteTimeEntry will completely remove a TimeEntry from a Release
func (s *ReleaseService) DeleteTimeEntry(releaseId int, timeEntryId int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(releaseTimeEntryIdUrl, releaseId, timeEntryId))
    return success, res, err
}
