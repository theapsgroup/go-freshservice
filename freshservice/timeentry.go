package freshservice

import "time"

// TimeEntry represents a TimeEntry associated to a Task / Ticket
type TimeEntry struct {
	ID           int       `json:"id"`
	StartTime    time.Time `json:"start_time"`
	ExecutedAt   time.Time `json:"executed_at"`
	TimerRunning bool      `json:"timer_running"`
	Billable     bool      `json:"billable"`
	TimeSpent    string    `json:"time_spent"`
	TaskID       int       `json:"task_id"`
	AgentID      int       `json:"agent_id"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// timeEntryWrapper contains Details of one TimeEntry
type timeEntryWrapper struct {
	Details TimeEntry `json:"time_entry"`
}

// TimeEntries contains Collection an array of TimeEntry
type TimeEntries struct {
	Collection []TimeEntry `json:"time_entries"`
}

// CreateTimeEntryModel is a data structure for creating TimeEntry
type CreateTimeEntryModel struct {
	AgentID   int    `json:"agent_id"`
	Note      string `json:"note"`
	TimeSpent string `json:"time_spent"`
}

type createTimeEntryWrapper struct {
	Data CreateTimeEntryModel `json:"time_entry"`
}
