package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	locationsUrl  = "locations"
	locationIdUrl = "locations/%d"
)

// LocationService API Docs: https://api.freshservice.com/#locations
type LocationService struct {
	client *Client
}

// Locations contains Collection an array of Location
type Locations struct {
	Collection []Location `json:"locations"`
}

// locationWrapper contains Details of a Location
type locationWrapper struct {
	Details Location `json:"location"`
}

// Location represents a FreshService Location (Physical Location)
type Location struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ParentLocationID int       `json:"parent_location_id"`
	PrimaryContactID int       `json:"primary_contact_id"`
	Address          Address   `json:"address"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// CreateLocationModel is a data struct for creating a new Location
type CreateLocationModel struct {
	Name             string  `json:"name"`
	ParentLocationID int     `json:"parent_location_id"`
	PrimaryContactID int     `json:"primary_contact_id"`
	Address          Address `json:"address"`
}

// UpdateLocationModel is a data struct for updating a Location
type UpdateLocationModel struct {
	Name             string  `json:"name"`
	ParentLocationID int     `json:"parent_location_id"`
	PrimaryContactID int     `json:"primary_contact_id"`
	Address          Address `json:"address"`
}

// Address representation of a physical address
type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipcode"`
}

// ListLocationsOptions represents filters/pagination for Locations
type ListLocationsOptions struct {
	ListOptions
}

// GetLocation will return a Location by id
func (s *LocationService) GetLocation(id int) (*Location, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf(locationIdUrl, id), nil)
	if err != nil {
		return nil, nil, err
	}

	l := new(locationWrapper)
	res, err := s.client.SendRequest(req, &l)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &l.Details, res, nil
}

// ListLocations will return paginated/filtered Locations using ListLocationsOptions
func (s *LocationService) ListLocations(opt *ListLocationsOptions) (*Locations, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationsUrl, opt)
	if err != nil {
		return nil, nil, err
	}

	ls := new(Locations)
	res, err := s.client.SendRequest(req, &ls)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return ls, res, nil
}

// CreateLocation will create and return a new Location based on CreateLocationModel
func (s *LocationService) CreateLocation(newLocation *CreateLocationModel) (*Location, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, locationsUrl, newLocation)
	if err != nil {
		return nil, nil, err
	}

	l := new(locationWrapper)
	res, err := s.client.SendRequest(req, &l)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &l.Details, res, nil
}

// UpdateLocation will update and return a Location matching id based UpdateLocationModel
func (s *LocationService) UpdateLocation(id int, location *UpdateLocationModel) (*Location, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf(locationIdUrl, id), location)
	if err != nil {
		return nil, nil, err
	}

	l := new(locationWrapper)
	res, err := s.client.SendRequest(req, &l)
	if b, s := isSuccessful(res); !b {
		return nil, res, fmt.Errorf("%s: %v", s, err)
	}

	return &l.Details, res, nil
}

// DeleteLocation will completely remove a Location from FreshService matching id
func (s *LocationService) DeleteLocation(id int) (bool, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf(locationIdUrl, id), nil)
	if err != nil {
		return false, nil, err
	}

	res, err := s.client.SendRequest(req, nil)
	if b, s := isSuccessful(res); !b {
		return false, res, fmt.Errorf("%s: %v", s, err)
	}

	return true, res, nil
}
