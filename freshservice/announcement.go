package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// AnnouncementService API Docs: https://api.freshservice.com/#announcements
type AnnouncementService struct {
	client *Client
}

// Announcements contains collection of Announcement
type Announcements struct {
	Collection []Announcement `json:"announcements"`
}

// SpecificAnnouncement contains Details of one specific Announcement
type SpecificAnnouncement struct {
	Details Announcement `json:"announcement"`
}

// Announcement represents a FreshService Announcement
type Announcement struct {
	ID               int       `json:"id"`
	CreatedBy        int       `json:"created_by"`
	State            string    `json:"state"`
	Title            string    `json:"title"`
	Body             string    `json:"body"`
	BodyHtml         string    `json:"body_html"`
	VisibleFrom      time.Time `json:"visible_from"`
	VisibleTo        time.Time `json:"visible_to"`
	Visibility       string    `json:"visibility"`
	Departments      []int     `json:"departments"`
	Groups           []int     `json:"groups"`
	IsRead           bool      `json:"is_read"`
	SendEmail        bool      `json:"send_email"`
	AdditionalEmails []string  `json:"additional_emails"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// NewAnnouncement is the data structure required to create an Announcement
type NewAnnouncement struct {
	Title            string    `json:"title"`
	BodyHtml         string    `json:"body_html"`
	VisibleFrom      time.Time `json:"visible_from"`
	VisibleTo        time.Time `json:"visible_to"`
	Visibility       string    `json:"visibility"`
	Departments      []int     `json:"departments"`
	Groups           []int     `json:"groups"`
	SendEmail        bool      `json:"send_email"`
	AdditionalEmails []string  `json:"additional_emails"`
}

// UpdateAnnouncement is the data structure required to update an Announcement
type UpdateAnnouncement struct {
	Title            string    `json:"title"`
	BodyHtml         string    `json:"body_html"`
	VisibleFrom      time.Time `json:"visible_from"`
	VisibleTo        time.Time `json:"visible_to"`
	Visibility       string    `json:"visibility"`
	Departments      []int     `json:"departments"`
	Groups           []int     `json:"groups"`
	SendEmail        bool      `json:"send_email"`
	AdditionalEmails []string  `json:"additional_emails"`
}

// ListAnnouncementOptions represents filters for Announcements
type ListAnnouncementOptions struct {
	ListOptions
	State string `json:"state,omitempty" url:"state,omitempty"`
}

// GetAnnouncement will return a single Announcement by id
func (s *AnnouncementService) GetAnnouncement(id int) (*Announcement, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("announcements/%v", id), nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAnnouncement)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// GetAnnouncements will return Announcements collection
func (s *AnnouncementService) GetAnnouncements(opt ListAnnouncementOptions) (*Announcements, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "announcements", opt)
	if err != nil {
		return nil, nil, err
	}

	as := new(Announcements)
	res, err := s.client.SendRequest(req, &as)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return as, res, nil
}

// CreateAnnouncement will create a new Announcement in FreshService
func (s *AnnouncementService) CreateAnnouncement(newAnnouncement *NewAnnouncement) (*Announcement, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "announcements", newAnnouncement)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAnnouncement)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// UpdateAnnouncement will update the Announcement matching the id and return the updated Announcement
func (s *AnnouncementService) UpdateAnnouncement(id int, announcement *UpdateAnnouncement) (*Announcement, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("announcements/%d", id), announcement)
	if err != nil {
		return nil, nil, err
	}

	a := new(SpecificAnnouncement)
	res, err := s.client.SendRequest(req, &a)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &a.Details, res, nil
}

// DeleteAnnouncement irrecoverably removes an Announcement from FreshService matching the id
func (s *AnnouncementService) DeleteAnnouncement(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("announcements/%d", id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
