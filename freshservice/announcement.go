package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	announcementsUrl  = "announcements"
	announcementIdUrl = "announcements/%d"
)

// AnnouncementService API Docs: https://api.freshservice.com/#announcements
type AnnouncementService struct {
	client *Client
}

// Announcements contains collection of Announcement
type Announcements struct {
	Collection []Announcement `json:"announcements"`
}

// announcementWrapper contains Details of one Announcement
type announcementWrapper struct {
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

// CreateAnnouncementModel is the data structure required to create an Announcement
type CreateAnnouncementModel struct {
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

// UpdateAnnouncementModel is the data structure required to update an Announcement
type UpdateAnnouncementModel struct {
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

// ListAnnouncementsOptions represents filters/pagination for Announcements
type ListAnnouncementsOptions struct {
	ListOptions
	State string `json:"state,omitempty" url:"state,omitempty"`
}

// GetAnnouncement will return a single Announcement by id
func (s *AnnouncementService) GetAnnouncement(id int) (*Announcement, *http.Response, error) {
	o := new(announcementWrapper)
	res, err := s.client.Get(fmt.Sprintf(announcementIdUrl, id), &o)
	return &o.Details, res, err
}

// ListAnnouncements will return paginated/filtered Announcements using ListAnnouncementsOptions
func (s *AnnouncementService) ListAnnouncements(opt ListAnnouncementsOptions) (*Announcements, *http.Response, error) {
	o := new(Announcements)
	res, err := s.client.List(announcementsUrl, opt, &o)
	return o, res, err
}

// CreateAnnouncement will create and return a new Announcement based on CreateAnnouncementModel
func (s *AnnouncementService) CreateAnnouncement(newAnnouncement *CreateAnnouncementModel) (*Announcement, *http.Response, error) {
	o := new(announcementWrapper)
	res, err := s.client.Post(announcementsUrl, newAnnouncement, &o)
	return &o.Details, res, err
}

// UpdateAnnouncement will update and return the Announcement matching the id based on UpdateAnnouncementModel
func (s *AnnouncementService) UpdateAnnouncement(id int, announcement *UpdateAnnouncementModel) (*Announcement, *http.Response, error) {
	o := new(announcementWrapper)
	res, err := s.client.Put(fmt.Sprintf(announcementIdUrl, id), announcement, &o)
	return &o.Details, res, err
}

// DeleteAnnouncement irrecoverably removes an Announcement from FreshService matching the id
func (s *AnnouncementService) DeleteAnnouncement(id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(announcementIdUrl, id))
	return success, res, err
}
