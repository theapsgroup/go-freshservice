package freshservice

import (
	"fmt"
	"net/http"
	"time"
)

const (
	customObjectRecordCreateUrl = "objects/%d/records"    //POST create records for a specific custom object
	customObjectRecordListUrl   = "objects/%d/records"    //GET list records for a specific custom object
	customObjectRecordEditUrl   = "objects/%d/records/%d" //PUT or DELETE for a specific custom object and record
)

type CustomObjectService struct {
	client *Client
}

type CustomObject struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Fields      []Fields `json:"fields,omitempty"`
	Meta        int      `json:"total_records_count,omitempty"`
}

type Fields struct {
	Name     string   `json:"name,omitempty"`
	Label    string   `json:"label,omitempty"`
	Type     string   `json:"type,omitempty"`
	Required bool     `json:"required,omitempty"`
	Choices  []string `json:"choices,omitempty"`
	Meta     string   `json:"meta,omitempty"`
}

// CustomObjectRecord contains a single custom object record
type CustomObjectRecord struct {
	CreatedAt      time.Time `json:"bo_created_at,omitempty"`
	CreatedBy      CreatedBy `json:"bo_created_by,omitempty"`
	DisplayId      int       `json:"bo_display_id,omitempty"`
	UpdatedAt      time.Time `json:"bo_updated_at,omitempty"`
	UpdatedBy      UpdatedBy `json:"bo_updated_by,omitempty"`
	DataSource     string    `json:"data_source_name,omitempty"`
	DataSourceType string    `json:"data_source_id,omitempty"`
}

type CreatedBy struct {
	Id    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type UpdatedBy struct {
	ID    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

// CustomObjectRecords contains an array of CustomObject
type CustomObjectRecords struct {
	Collection   []CustomObjectRecordWrapper `json:"records,omitempty"`
	NextPageLink string                      `json:"next_page_link,omitempty"`
	Meta         map[string]interface{}      `json:"meta,omitempty"`
}

// CustomObjects contains Collection an array of CustomObject
type CustomObjects struct {
	Collection []CustomObject `json:"custom_objects"`
}

//CustomObjectRecordWrapper contains Details of a CustomObjectRecords
type CustomObjectRecordWrapper struct {
	Data CustomObjectRecord `json:"data,omitempty"`
}

type CustomObjectRecordUpdate struct {
	CustomObjectRecord CustomObjectRecord `json:"custom_object,omitempty"`
}

// ListCustomObjectsOptions represents filters/pagination for Products
type ListCustomObjectsOptions struct {
	ListOptions
	Email        *string    `json:"email,omitempty" url:"email,omitempty"`
	RequesterID  *int       `json:"requester_id,omitempty" url:"requester_id,omitempty"`
	UpdatedSince *time.Time `json:"updated_since,omitempty" url:"updated_since,omitempty"`
	Type         *string    `json:"type,omitempty" url:"type,omitempty"`
}

// GetCustomObjectRecords returns a list of records for a specific custom object id
func (s *CustomObjectService) GetCustomObjectRecords(coID int) (*CustomObjectRecords, *http.Response, error) {
	o := new(CustomObjectRecords)
	res, err := s.client.Get(fmt.Sprintf(customObjectRecordListUrl, coID), &o)
	return o, res, err
}

// UpdateCustomObjectRecord will update the record(s) based on customObject id and record id
func (s *CustomObjectService) UpdateCustomObjectRecord(coID int, recID int, object *CustomObjectRecordWrapper) (*CustomObjectRecordUpdate, *http.Response, error) {
	o := new(CustomObjectRecordUpdate)
	res, err := s.client.Put(fmt.Sprintf(customObjectRecordEditUrl, coID, recID), object, o)
	return o, res, err
}

// CreateCustomObjectRecord will create a custom object record based on the custom object id.
func (s *CustomObjectService) CreateCustomObjectRecord(coID int, newCustomObjectRecord *CustomObjectRecordWrapper) (*CustomObjectRecordUpdate, *http.Response, error) {
	o := new(CustomObjectRecordUpdate)
	res, err := s.client.Post(fmt.Sprintf(customObjectRecordCreateUrl, coID), newCustomObjectRecord, o)
	return o, res, err
}

// DeleteCustomObjectRecord will delete a custom object record based on custom object id and record id
func (s *CustomObjectService) DeleteCustomObjectRecord(coID int, recID int) (bool, *http.Response, error) {
	success, res, err := s.client.Delete(fmt.Sprintf(customObjectRecordEditUrl, coID, recID))
	return success, res, err
}
