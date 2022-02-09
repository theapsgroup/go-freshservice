package freshservice

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// SoftwareInstallations contains Collection an array of SoftwareInstallation
type SoftwareInstallations struct {
	Collection []SoftwareInstallation `json:"installations"`
}

// SoftwareInstallation represents a binding between an Application and a Device in FreshService
type SoftwareInstallation struct {
	ID                    int       `json:"id"`
	InstallationMachineID int       `json:"installation_machine_id"`
	InstallationPath      string    `json:"installation_path"`
	Version               string    `json:"version"`
	UserID                int       `json:"user_id"`
	DepartmentID          int       `json:"department_id"`
	InstallationDate      time.Time `json:"installation_date"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type CreateInstallationModel struct {
	InstallationMachineID int       `json:"installation_machine_id"`
	InstallationPath      string    `json:"installation_path"`
	Version               string    `json:"version"`
	InstallationDate      time.Time `json:"installation_date"`
}

// AddInstallation allows for adding a Device to an Application as a SoftwareInstallation
func (s *SoftwareService) AddInstallation(applicationId int, installation *CreateInstallationModel) (*SoftwareInstallation, *http.Response, error) {
	o := new(SoftwareInstallation)
	res, err := s.client.Post(fmt.Sprintf(applicationInstallationsUrl, applicationId), installation, &o)
	return o, res, err
}

// ListInstallations will return SoftwareInstallations for a specific Application
func (s *SoftwareService) ListInstallations(applicationId int) (*SoftwareInstallations, *http.Response, error) {
	o := new(SoftwareInstallations)
	res, err := s.client.List(fmt.Sprintf(applicationInstallationsUrl, applicationId), nil, &o)
	return o, res, err
}

// DeleteInstallations allows for bulk removal of Devices from Application
func (s *SoftwareService) DeleteInstallations(applicationId int, deviceIds []string) (bool, *http.Response, error) {
	path := fmt.Sprintf(applicationInstallationsUrl, applicationId)
	q := strings.Join(deviceIds, ",")
	success, res, err := s.client.Delete(fmt.Sprintf("%s?device_ids=%s", path, q))
	return success, res, err
}
