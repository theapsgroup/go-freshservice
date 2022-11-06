package freshservice

// ListOptions defines basic options for pagination, etc
type ListOptions struct {
	Page    int `json:"page,omitempty" url:"page,omitempty"`
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
}

// Actor represents a simple id/name object
type Actor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ActorEmail represents a simple id/name/email object
type ActorEmail struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Description represents a Description/DescriptionText pairing.
type Description struct {
	Description     string `json:"description"`
	DescriptionText string `json:"description_text"`
}
