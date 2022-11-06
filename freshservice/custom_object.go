package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	customObjectsUrl        = "objects"
	customObjectIdUrl       = "objects/%d"
	customObjectRecordsUrl  = "objects/%d/records"
	customObjectRecordIdUrl = "objects/%d/records/%d"
)

// CustomObjectService API Docs: https://api.freshservice.com/#custom-objects
type CustomObjectService struct {
	client *Client
}

// CustomObjects contains Collection an array of CustomObject
type CustomObjects struct {
	Collection []CustomObject `json:"custom_objects"`
}

// customObjectWrapper contains Details of a CustomObject
type customObjectWrapper struct {
	Details CustomObject `json:"custom_object"`
}

// CustomObject represents a FreshService CustomObject
type CustomObject struct {
	ID            int        `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LastUpdatedBy ActorEmail `json:"last_updated_by"`
	Fields        []Field    `json:"fields,omitempty"`
}

// CustomObjectMeta represents the metadata for a CustomObject
type CustomObjectMeta struct {
	Count             int `json:"count,omitempty"`
	Page              int `json:"page,omitempty"`
	PerPage           int `json:"per_page,omitempty"`
	TotalRecordsCount int `json:"total_records_count,omitempty"`
}

// Field represents a single field within the CustomObject
type Field struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Type     string    `json:"type"`
	Required bool      `json:"required"`
	Choices  []string  `json:"choices,omitempty"`
	Meta     FieldMeta `json:"meta,omitempty"`
}

// FieldMeta represents the metadata associated with a Field on a CustomObject
type FieldMeta struct {
	SourceName string `json:"source_name,omitempty"`
}

// customObjectRecordWrapper wraps CustomObjectRecord
type customObjectRecordWrapper struct {
	CustomObject CustomObjectRecord `json:"custom_object"`
}

// CustomObjectRecord represents a record of a CustomObject
// TODO: Figure out how to build the data out dynamically since all custom objects can vary
type CustomObjectRecord struct {
	Data interface{} `json:"data"`
}

// ListCustomObjectsOptions represents filters/pagination for CustomObjects
type ListCustomObjectsOptions struct {
	ListOptions
}

// GetCustomObject will return a CustomObject by id
func (s *CustomObjectService) GetCustomObject(id int) (*CustomObject, *http.Response, error) {
	o := new(customObjectWrapper)
	res, err := s.client.Get(fmt.Sprintf(customObjectIdUrl, id), &o)
	return &o.Details, res, err
}

// ListCustomObjects will return paginated/filtered CustomObjectSummary using ListCustomObjectsOptions
func (s *CustomObjectService) ListCustomObjects(opt *ListCustomObjectsOptions) (*CustomObjects, *http.Response, error) {
	o := new(CustomObjects)
	res, err := s.client.List(customObjectsUrl, opt, &o)
	return o, res, err
}

// TODO: func (s *CustomObjectService) ListCustomObjectRecords()

// CreateCustomObjectRecord should create and return a new CustomObjectRecord
func (s *CustomObjectService) CreateCustomObjectRecord(customObjectId int, newRecord CustomObjectRecord) (*CustomObjectRecord, *http.Response, error) {
	o := new(customObjectRecordWrapper)
	res, err := s.client.Post(fmt.Sprintf(customObjectRecordsUrl, customObjectId), newRecord, &o)
	return &o.CustomObject, res, err
}

// UpdateCustomObjectRecord should update and return a CustomObjectRecord
func (s *CustomObjectService) UpdateCustomObjectRecord(customObjectId int, recordId int, record CustomObjectRecord) (*CustomObjectRecord, *http.Response, error) {
	o := new(customObjectRecordWrapper)
	res, err := s.client.Put(fmt.Sprintf(customObjectRecordIdUrl, customObjectId, recordId), record, &o)
	return &o.CustomObject, res, err
}

// DeleteCustomObjectRecord will completely remove a CustomObjectRecord from FreshService with a matching id
func (s *CustomObjectService) DeleteCustomObjectRecord(customObjectId int, id int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(customObjectRecordIdUrl, customObjectId, id))
	return success, res, err
}
