package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// Task represents a Task on a FreshService Ticket
type Task struct {
	ID           int       `json:"id"`
	AgentID      int       `json:"agent_id"`
	Status       int       `json:"status"`
	DueDate      time.Time `json:"due_date"`
	NotifyBefore int       `json:"notify_before"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	GroupID      int       `json:"group_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
}

// Tasks contains Collection an array of Task
type Tasks struct {
	Collection []Task `json:"tasks"`
}

// taskWrapper contains Details of one Task
type taskWrapper struct {
	Details Task `json:"task"`
}

// CreateTaskModel is the data structure required to create a new Task
type CreateTaskModel struct {
	DueDate      time.Time `json:"due_date"`
	NotifyBefore int       `json:"notify_before"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
}

// UpdateTaskModel is the data structure for updating an existing Task
type UpdateTaskModel struct {
	AgentID      int       `json:"agent_id"`
	Status       int       `json:"status"`
	DueDate      time.Time `json:"due_date"`
	NotifyBefore int       `json:"notify_before"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	GroupID      int       `json:"group_id"`
}

// ListTasksOptions represents filters/pagination for Tasks
type ListTasksOptions struct {
	ListOptions
}

// GetTask will return a single Task from a Ticket by the id
func (s *TicketService) GetTask(ticketId int, taskId int) (*Task, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(ticketTaskIdUrl, ticketId, taskId), nil)
	if err != nil {
		return nil, nil, err
	}

	t := new(taskWrapper)
	res, err := s.client.SendRequest(req, &t)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &t.Details, res, nil
}

// ListTasks will return paginated/filtered Tasks using ListTasksOptions
func (s *TicketService) ListTasks(opt *ListTasksOptions) (*Tasks, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, ticketTasksUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	ts := new(Tasks)
	res, err := s.client.SendRequest(req, &ts)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return ts, res, nil
}

// CreateTask will create and return a new Task based on CreateTaskModel
func (s *TicketService) CreateTask(newTask *CreateTaskModel) (*Task, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, ticketTasksUrl, newTask)
	if err != nil {
		return nil, nil, err
	}

	t := new(taskWrapper)
	res, err := s.client.SendRequest(req, &t)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &t.Details, res, nil
}

// UpdateTask will update and return a Task matching id based on UpdateTaskModel
func (s *TicketService) UpdateTask(ticketId int, taskId int, task *UpdateTaskModel) (*Task, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(ticketTaskIdUrl, ticketId, taskId), task)
	if err != nil {
		return nil, nil, err
	}

	t := new(taskWrapper)
	res, err := s.client.SendRequest(req, &t)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &t.Details, res, nil
}

// DeleteTask deletes the Task on a Ticket with the given ID
func (s *TicketService) DeleteTask(ticketId int, taskId int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(ticketTaskIdUrl, ticketId, taskId), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
