package freshservice

import "time"

// Notes contains Collection an array of Note
type Notes struct {
	Collection []Note `json:"notes"`
}

// noteWrapper contains Details of one Note
type noteWrapper struct {
	Details Note `json:"note"`
}

// Note represents a Note attached to a Change
type Note struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Body         string    `json:"body"`
	BodyText     string    `json:"body_text"`
	NotifyEmails []string  `json:"notify_emails"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// UpsertNoteModel is a data struct for creating/updating Note
type UpsertNoteModel struct {
	Body string `json:"body"`
}
