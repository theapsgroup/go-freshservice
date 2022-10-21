package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	customObjectsUrl            = "objects"               //GET all custom objects
	customObjectRecordUrl       = "objects/%d"            //GET specific custom object
	customObjectRecordCreateUrl = "objects/%d/records"    //POST create records for a specific custom object
	customObjectRecordListUrl   = "objects/%d/records"    //GET list records for a specific custom object
	customObjectRecordEditUrl   = "objects/%d/records/%d" //PUT or DELETE for a specific custom object and record
)

type CustomObjectService struct {
	client *Client
}

// CustomObjects contains Collection an array of CustomObject
type CustomObjects struct {
	Collection []CustomObject `json:"custom_objects"`
}

// CustomObjectWrapper contains Details of a CustomObject
type CustomObjectWrapper struct {
	Details CustomObject `json:"custom_object"`
}

type CustomObject struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	//LastUpdatedBy
	//Meta
}

type LastUpdatedBy struct {
	Email  string `json:"email,omitempty"`
	ID     int    `json:"id,omitempty"`
	System string `json:"name,omitempty"`
}
type Meta struct {
	Count   int `json:"count,omitempty"`
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// ListCustomObjectsOptions represents filters/pagination for Products
type ListCustomObjectsOptions struct {
	ListOptions
	Email        *string    `json:"email,omitempty" url:"email,omitempty"`
	RequesterID  *int       `json:"requester_id,omitempty" url:"requester_id,omitempty"`
	UpdatedSince *time.Time `json:"updated_since,omitempty" url:"updated_since,omitempty"`
	Type         *string    `json:"type,omitempty" url:"type,omitempty"`
}

// GetCustomObjectsList will return a list of all custom objects
func (s *CustomObjectService) GetCustomObjectsList() (*CustomObject, *http.Response, error) {
	o := new(CustomObjectWrapper)
	res, err := s.client.List(customObjectsUrl, nil, &o)
	return &o.Details, res, err
}

// ListCustomObjects will return a list of all records for a specific custom object.
func (s *CustomObjectService) ListCustomObjects(opt *ListCustomObjectsOptions) (*CustomObjects, *http.Response, error) {
	o := new(CustomObjects)
	res, err := s.client.List(customObjectRecordUrl, opt, &o)
	return o, res, err
}

// GetCustomObjectRecords returns a list of records for a specific custom object id
func (s *CustomObjectService) GetCustomObjectRecords(id int) (*CustomObject, *http.Response, error) {
	o := new(CustomObjectWrapper)
	res, err := s.client.Get(fmt.Sprintf(customObjectRecordListUrl, id), &o)
	return &o.Details, res, err
}

// CreateCustomObjectRecord will create a custom object based on the CustomObject struct
func (s *CustomObjectService) CreateCustomObjectRecord(newCustomObject *CustomObject) (*CustomObject, *http.Response, error) {
	o := new(CustomObjectWrapper)
	res, err := s.client.Post(customObjectRecordCreateUrl, newCustomObject, &o)
	return &o.Details, res, err
}

// UpdateCustomObjectRecord will update the record(s) based on customObject id and record id
func (s *CustomObjectService) UpdateCustomObjectRecord(co, rec int, object *CustomObject) (*CustomObject, *http.Response, error) {
	o := new(CustomObjectWrapper)
	res, err := s.client.Put(fmt.Sprintf(customObjectRecordEditUrl, co, rec), object, &o)
	return &o.Details, res, err
}

// DeleteCustomObjectRecord will delete a custom object record based on custom object id and record id
func (s *CustomObjectService) DeleteCustomObjectRecord(object *CustomObject, id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(customObjectRecordEditUrl, object, id))
	return success, res, err
}
