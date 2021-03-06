package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	businessHoursUrl   = "business_hours"
	businessHoursIdUrl = "business_hours/%d"
)

// BusinessHoursService API Docs: https://api.freshservice.com/#business-hours
type BusinessHoursService struct {
	client *Client
}

// BusinessHours contains Collection an array of BusinessHour
type BusinessHours struct {
	Collection []BusinessHour `json:"business_hours"`
}

// businessHourWrapper contains Details of one BusinessHour
type businessHourWrapper struct {
	Details BusinessHour `json:"business_hours"`
}

// BusinessHour represents the BusinessHour configuration in the FreshService instance
type BusinessHour struct {
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

// ListBusinessHoursOptions represents filters/pagination for BusinessHour
type ListBusinessHoursOptions struct {
	ListOptions
}

// GetBusinessHours will return a single BusinessHour configuration by id
func (s *BusinessHoursService) GetBusinessHours(id int) (*BusinessHour, *http.Response, error) {
	o := new(businessHourWrapper)
	res, err := s.client.Get(fmt.Sprintf(businessHoursIdUrl, id), &o)
	return &o.Details, res, err
}

// ListBusinessHours will return paginated/filtered BusinessHours using ListBusinessHoursOptions
func (s *BusinessHoursService) ListBusinessHours(opt *ListBusinessHoursOptions) (*BusinessHours, *http.Response, error) {
	o := new(BusinessHours)
	res, err := s.client.List(businessHoursUrl, opt, &o)
	return o, res, err
}
