package freshservice

import (
    "fmt"
    "net/http"
)

// GetTask will return a single Task from a Problem by the id
func (s *ProblemService) GetTask(problemId int, taskId int) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Get(fmt.Sprintf(problemTaskIdUrl, problemId, taskId), &o)
    return &o.Details, res, err
}

// ListTasks will return paginated/filtered Tasks using ListTasksOptions
func (s *ProblemService) ListTasks(problemId int, opt *ListTasksOptions) (*Tasks, *http.Response, error) {
    o := new(Tasks)
    res, err := s.client.List(fmt.Sprintf(problemTasksUrl, problemId), opt, &o)
    return o, res, err
}

// CreateTask will create and return a new Task based on CreateTaskModel
func (s *ProblemService) CreateTask(problemId int, newTask *CreateTaskModel) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Post(fmt.Sprintf(problemTasksUrl, problemId), newTask, &o)
    return &o.Details, res, err
}

// UpdateTask will update and return a Task matching id based on UpdateTaskModel
func (s *ProblemService) UpdateTask(problemId int, taskId int, task *UpdateTaskModel) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Put(fmt.Sprintf(problemTaskIdUrl, problemId, taskId), task, &o)
    return &o.Details, res, err
}

// DeleteTask deletes the Task on a Problem with the given ID
func (s *ProblemService) DeleteTask(problemId int, taskId int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(problemTaskIdUrl, problemId, taskId))
    return success, res, err
}
