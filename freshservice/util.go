package freshservice

// ListOptions defines basic options for pagination, etc
type ListOptions struct {
    Page    int `json:"page,omitempty" url:"page,omitempty"`
    PerPage int `json:"per_page,omitempty" url:"page,omitempty"`
}
