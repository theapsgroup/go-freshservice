package freshservice

import (
    "fmt"
    "net/http"
)

// GetTask will return a single Task from a Release by the id
func (s *ReleaseService) GetTask(releaseId int, taskId int) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Get(fmt.Sprintf(releaseTaskIdUrl, releaseId, taskId), &o)
    return &o.Details, res, err
}

// ListTasks will return paginated/filtered Tasks using ListTasksOptions
func (s *ReleaseService) ListTasks(releaseId int, opt *ListTasksOptions) (*Tasks, *http.Response, error) {
    o := new(Tasks)
    res, err := s.client.List(fmt.Sprintf(releaseTasksUrl, releaseId), opt, &o)
    return o, res, err
}

// CreateTask will create and return a new Task based on CreateTaskModel
func (s *ReleaseService) CreateTask(releaseId int, newTask *CreateTaskModel) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Post(fmt.Sprintf(releaseTasksUrl, releaseId), newTask, &o)
    return &o.Details, res, err
}

// UpdateTask will update and return a Task matching id based on UpdateTaskModel
func (s *ReleaseService) UpdateTask(releaseId int, taskId int, task *UpdateTaskModel) (*Task, *http.Response, error) {
    o := new(taskWrapper)
    res, err := s.client.Put(fmt.Sprintf(releaseTaskIdUrl, releaseId, taskId), task, &o)
    return &o.Details, res, err
}

// DeleteTask deletes the Task on a Release with the given ID
func (s *ReleaseService) DeleteTask(releaseId int, taskId int) (bool, *http.Response, error) {
    success, res, err := s.client.Delete(fmt.Sprintf(releaseTaskIdUrl, releaseId, taskId))
    return success, res, err
}
