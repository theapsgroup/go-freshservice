package freshservice

import (
    "fmt"
    "net/http"
)

// GetReleaseNote will return a single Note by id
func (s *ReleaseService) GetReleaseNote(releaseId int, releaseNoteId int) (*Note, *http.Response, error) {
    o := new(noteWrapper)
    res, err := s.client.Get(fmt.Sprintf(releaseNoteIdUrl, releaseId, releaseNoteId), &o)
    return &o.Details, res, err
}

// ListReleaseNotes will return  Notes for a specific Release
func (s *ReleaseService) ListReleaseNotes(releaseId int) (*Notes, *http.Response, error) {
    o := new(Notes)
    res, err := s.client.List(fmt.Sprintf(releaseNotesUrl, releaseId), nil, &o)
    return o, res, err
}

// CreateReleaseNote will create and return a new Note based on UpsertNoteModel
func (s *ReleaseService) CreateReleaseNote(releaseId int, note *UpsertNoteModel) (*Note, *http.Response, error) {
    o := new(noteWrapper)
    res, err := s.client.Post(fmt.Sprintf(releaseNotesUrl, releaseId), note, &o)
    return &o.Details, res, err
}

// UpdateReleaseNote will update and return a Note matching id based on UpsertNoteModel
func (s *ReleaseService) UpdateReleaseNote(releaseId int, releaseNoteId int, note *UpsertNoteModel) (*Note, *http.Response, error) {
    o := new(noteWrapper)
    res, err := s.client.Put(fmt.Sprintf(releaseNoteIdUrl, releaseId, releaseNoteId), note, &o)
    return &o.Details, res, err
}

// DeleteReleaseNote will completely remove a Note from a Release
func (s *ReleaseService) DeleteReleaseNote(releaseId int, releaseNoteId int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(releaseNoteIdUrl, releaseId, releaseNoteId))
    return success, res, err
}
