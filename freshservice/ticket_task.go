package freshservice

import (
	"fmt"
	"net/http"
)

// GetTask will return a single Task from a Ticket by the id
func (s *TicketService) GetTask(ticketId int, taskId int) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Get(fmt.Sprintf(ticketTaskIdUrl, ticketId, taskId), &o)
    return &o.Details, res, err
}

// ListTasks will return paginated/filtered Tasks using ListTasksOptions
func (s *TicketService) ListTasks(ticketId int, opt *ListTasksOptions) (*Tasks, *http.Response, error) {
    o := new(Tasks)
    res, err := s.client.List(fmt.Sprintf(ticketTasksUrl, ticketId), opt, &o)
    return o, res, err
}

// CreateTask will create and return a new Task based on CreateTaskModel
func (s *TicketService) CreateTask(ticketId int, newTask *CreateTaskModel) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Post(fmt.Sprintf(ticketTasksUrl, ticketId), newTask, &o)
    return &o.Details, res, err
}

// UpdateTask will update and return a Task matching id based on UpdateTaskModel
func (s *TicketService) UpdateTask(ticketId int, taskId int, task *UpdateTaskModel) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Put(fmt.Sprintf(ticketTaskIdUrl, ticketId, taskId), task, &o)
    return &o.Details, res, err
}

// DeleteTask deletes the Task on a Ticket with the given ID
func (s *TicketService) DeleteTask(ticketId int, taskId int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(ticketTaskIdUrl, ticketId, taskId))
    return success, res, err
}
