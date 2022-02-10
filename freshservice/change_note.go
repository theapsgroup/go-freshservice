package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// ChangeNotes contains Collection an array of ChangeNote
type ChangeNotes struct {
	Collection []ChangeNote `json:"notes"`
}

// changeNoteWrapper contains Details of one ChangeNote
type changeNoteWrapper struct {
	Details ChangeNote `json:"note"`
}

// ChangeNote represents a Note attached to a Change
type ChangeNote struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Body         string    `json:"body"`
	BodyText     string    `json:"body_text"`
	NotifyEmails []string  `json:"notify_emails"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ChangeNoteModel is a data struct for creating/updating ChangeNote
type ChangeNoteModel struct {
	Body string `json:"body"`
}

// GetChangeNote will return a single ChangeNote by id
func (s *ChangeService) GetChangeNote(changeId int, changeNoteId int) (*ChangeNote, *http.Response, error) {
	o := new(changeNoteWrapper)
	res, err := s.client.Get(fmt.Sprintf(changeNoteIdUrl, changeId, changeNoteId), &o)
	return &o.Details, res, err
}

// ListChangeNotes will return  ChangeNotes for a specific Change
func (s *ChangeService) ListChangeNotes(changeId int) (*ChangeNotes, *http.Response, error) {
	o := new(ChangeNotes)
	res, err := s.client.List(fmt.Sprintf(changeNotesUrl, changeId), nil, &o)
	return o, res, err
}

// CreateChangeNote will create and return a new ChangeNote based on CreateChangeNoteModel
func (s *ChangeService) CreateChangeNote(changeId int, note *ChangeNoteModel) (*ChangeNote, *http.Response, error) {
	o := new(changeNoteWrapper)
	res, err := s.client.Post(fmt.Sprintf(changeNotesUrl, changeId), note, &o)
	return &o.Details, res, err
}

// UpdateChangeNote will update and return an ChangeNote matching id based on UpdateChangeNoteModel
func (s *ChangeService) UpdateChangeNote(changeId int, changeNoteId int, note *ChangeNoteModel) (*ChangeNote, *http.Response, error) {
	o := new(changeNoteWrapper)
	res, err := s.client.Put(fmt.Sprintf(changeNoteIdUrl, changeId, changeNoteId), note, &o)
	return &o.Details, res, err
}

// DeleteChangeNote will completely remove an ChangeNote from FreshService
func (s *ChangeService) DeleteChangeNote(changeId int, changeNoteId int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(changeNoteIdUrl, changeId, changeNoteId))
	return success, res, err
}
