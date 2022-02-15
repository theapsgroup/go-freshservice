package freshservice

import (
	"fmt"
	"net/http"
)

// GetProblemNote will return a single Note by id
func (s *ProblemService) GetProblemNote(problemId int, problemNoteId int) (*Note, *http.Response, error) {
	o := new(noteWrapper)
	res, err := s.client.Get(fmt.Sprintf(problemNoteIdUrl, problemId, problemNoteId), &o)
	return &o.Details, res, err
}

// ListProblemNotes will return  Notes for a specific Problem
func (s *ProblemService) ListProblemNotes(problemId int) (*Notes, *http.Response, error) {
	o := new(Notes)
	res, err := s.client.List(fmt.Sprintf(problemNotesUrl, problemId), nil, &o)
	return o, res, err
}

// CreateProblemNote will create and return a new Note based on UpsertNoteModel
func (s *ProblemService) CreateProblemNote(problemId int, note *UpsertNoteModel) (*Note, *http.Response, error) {
	o := new(noteWrapper)
	res, err := s.client.Post(fmt.Sprintf(problemNotesUrl, problemId), note, &o)
	return &o.Details, res, err
}

// UpdateProblemNote will update and return a Note matching id based on UpsertNoteModel
func (s *ProblemService) UpdateProblemNote(problemId int, problemNoteId int, note *UpsertNoteModel) (*Note, *http.Response, error) {
	o := new(noteWrapper)
	res, err := s.client.Put(fmt.Sprintf(problemNoteIdUrl, problemId, problemNoteId), note, &o)
	return &o.Details, res, err
}

// DeleteProblemNote will completely remove a Note from a Problem
func (s *ProblemService) DeleteProblemNote(problemId int, problemNoteId int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(problemNoteIdUrl, problemId, problemNoteId))
	return success, res, err
}
