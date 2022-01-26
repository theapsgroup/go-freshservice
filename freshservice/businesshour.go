package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

// BusinessHoursService API Docs: https://api.freshservice.com/#business-hours
type BusinessHoursService struct {
	client *Client
}

// BusinessHoursCollection contains collection of BusinessHours
type BusinessHoursCollection struct {
	Collection []BusinessHours `json:"business_hours"`
}

// SpecificBusinessHours contains Details of one specific BusinessHours
type SpecificBusinessHours struct {
	Details BusinessHours `json:"business_hours"`
}

// BusinessHours represents the BusinessHours configuration in the FreshService instance
type BusinessHours struct {
	ID               int                `json:"id"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	IsDefault        bool               `json:"is_default"`
	TimeZone         string             `json:"time_zone"`
	ListOfHolidays   []Holiday          `json:"list_of_holidays"`
	ServiceDeskHours map[string]Workday `json:"service_desk_hours"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

// Workday is a data struct representing the start and end of a working day
type Workday struct {
	BeginningOfWorkday string `json:"beginning_of_workday"`
	EndOfWorkday       string `json:"end_of_workday"`
}

// Holiday is a data struct representing a configured holiday
type Holiday struct {
	HolidayDate string `json:"holiday_date"`
	HolidayName string `json:"holiday_name"`
}

// ListBusinessHoursOptions represents filters/pagination for BusinessHours
type ListBusinessHoursOptions struct {
	ListOptions
}

// GetBusinessHours will return a single BusinessHours configuration by id
func (s *BusinessHoursService) GetBusinessHours(id int) (*BusinessHours, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("business_hours/%v", id), nil)
	if err != nil {
		return nil, nil, err
	}

	bh := new(SpecificBusinessHours)
	res, err := s.client.SendRequest(req, &bh)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &bh.Details, res, nil
}

// ListBusinessHours will return BusinessHoursCollection
func (s *BusinessHoursService) ListBusinessHours(opt *ListBusinessHoursOptions) (*BusinessHours, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "business_hours", opt)
	if err != nil {
		return nil, nil, err
	}

	bhs := new(BusinessHours)
	res, err := s.client.SendRequest(req, &bhs)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return bhs, res, nil
}
