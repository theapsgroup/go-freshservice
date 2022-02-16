package freshservice

import "time"

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
