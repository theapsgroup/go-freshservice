package freshservice

import (
	"fmt"
	"net/http"
)

// GetChangeNote will return a single Note by id
func (s *ChangeService) GetChangeNote(changeId int, changeNoteId int) (*Note, *http.Response, error) {
	o := new(noteWrapper)
	res, err := s.client.Get(fmt.Sprintf(changeNoteIdUrl, changeId, changeNoteId), &o)
	return &o.Details, res, err
}

// ListChangeNotes will return  Notes for a specific Change
func (s *ChangeService) ListChangeNotes(changeId int) (*Notes, *http.Response, error) {
	o := new(Notes)
	res, err := s.client.List(fmt.Sprintf(changeNotesUrl, changeId), nil, &o)
	return o, res, err
}

// CreateChangeNote will create and return a new Note based on UpsertNoteModel
func (s *ChangeService) CreateChangeNote(changeId int, note *UpsertNoteModel) (*Note, *http.Response, error) {
	o := new(noteWrapper)
	res, err := s.client.Post(fmt.Sprintf(changeNotesUrl, changeId), note, &o)
	return &o.Details, res, err
}

// UpdateChangeNote will update and return a Note matching id based on UpsertNoteModel
func (s *ChangeService) UpdateChangeNote(changeId int, changeNoteId int, note *UpsertNoteModel) (*Note, *http.Response, error) {
	o := new(noteWrapper)
	res, err := s.client.Put(fmt.Sprintf(changeNoteIdUrl, changeId, changeNoteId), note, &o)
	return &o.Details, res, err
}

// DeleteChangeNote will completely remove a Note from a Change
func (s *ChangeService) DeleteChangeNote(changeId int, changeNoteId int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(changeNoteIdUrl, changeId, changeNoteId))
	return success, res, err
}
