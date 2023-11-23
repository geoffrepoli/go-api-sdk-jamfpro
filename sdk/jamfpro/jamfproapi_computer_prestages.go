// jamfproapi_computer_prestages.go
// Jamf Pro Api - Computer Prestages
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-prestages-scope
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const uriComputerPrestagesV2 = "/api/v2/computer-prestages"
const uriComputerPrestagesV3 = "/api/v3/computer-prestages"

// ResponseDeviceScope represents the structure of the response for a specific computer prestage scope.
type ResponseDeviceScope struct {
	PrestageId  string           `json:"prestageId"`
	Assignments []AssignmentItem `json:"assignments"`
	VersionLock int              `json:"versionLock"`
}

// AssignmentItem represents the structure of each assignment within the prestage scope.
type AssignmentItem struct {
	SerialNumber   string `json:"serialNumber"`
	AssignmentDate string `json:"assignmentDate"`
	UserAssigned   string `json:"userAssigned"`
}

type ResponseComputerPrestagesV3 struct {
	TotalCount *int                    `json:"totalCount"`
	Results    []ComputerPrestagesItem `json:"results"`
}

type ComputerPrestagesItem struct {
	DisplayName                       string                                 `json:"displayName"`
	Mandatory                         bool                                   `json:"mandatory"`
	MDMRemovable                      bool                                   `json:"mdmRemovable"`
	SupportPhoneNumber                string                                 `json:"supportPhoneNumber"`
	SupportEmailAddress               string                                 `json:"supportEmailAddress"`
	Department                        string                                 `json:"department"`
	DefaultPrestage                   bool                                   `json:"defaultPrestage"`
	EnrollmentSiteId                  string                                 `json:"enrollmentSiteId"`
	KeepExistingSiteMembership        bool                                   `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation   bool                                   `json:"keepExistingLocationInformation"`
	RequireAuthentication             bool                                   `json:"requireAuthentication"`
	AuthenticationPrompt              string                                 `json:"authenticationPrompt"`
	PreventActivationLock             bool                                   `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock   bool                                   `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string                                 `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                    map[string]bool                        `json:"skipSetupItems"`
	LocationInformation               ComputerPrestagesLocationInformation   `json:"locationInformation"`
	PurchasingInformation             ComputerPrestagesPurchasingInformation `json:"purchasingInformation"`
	AnchorCertificates                []string                               `json:"anchorCertificates"`
	EnrollmentCustomizationId         string                                 `json:"enrollmentCustomizationId"`
	Language                          string                                 `json:"language"`
	Region                            string                                 `json:"region"`
	AutoAdvanceSetup                  bool                                   `json:"autoAdvanceSetup"`
	InstallProfilesDuringSetup        bool                                   `json:"installProfilesDuringSetup"`
	PrestageInstalledProfileIds       []string                               `json:"prestageInstalledProfileIds"`
	CustomPackageIds                  []string                               `json:"customPackageIds"`
	CustomPackageDistributionPointId  string                                 `json:"customPackageDistributionPointId"`
	EnableRecoveryLock                bool                                   `json:"enableRecoveryLock"`
	RecoveryLockPasswordType          string                                 `json:"recoveryLockPasswordType"`
	RecoveryLockPassword              string                                 `json:"recoveryLockPassword"`
	RotateRecoveryLockPassword        bool                                   `json:"rotateRecoveryLockPassword"`
	ID                                string                                 `json:"id"`
	ProfileUuid                       string                                 `json:"profileUuid"`
	SiteId                            string                                 `json:"siteId"`
	VersionLock                       int                                    `json:"versionLock"`
	AccountSettings                   ComputerPrestagesAccountSettings       `json:"accountSettings"`
}

type ComputerPrestagesLocationInformation struct {
	Username     string `json:"username"`
	Realname     string `json:"realname"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Room         string `json:"room"`
	Position     string `json:"position"`
	DepartmentId string `json:"departmentId"`
	BuildingId   string `json:"buildingId"`
	ID           string `json:"id"`
	VersionLock  int    `json:"versionLock"`
}

type ComputerPrestagesPurchasingInformation struct {
	ID                string `json:"id"`
	Leased            bool   `json:"leased"`
	Purchased         bool   `json:"purchased"`
	AppleCareId       string `json:"appleCareId"`
	PONumber          string `json:"poNumber"`
	Vendor            string `json:"vendor"`
	PurchasePrice     string `json:"purchasePrice"`
	LifeExpectancy    int    `json:"lifeExpectancy"`
	PurchasingAccount string `json:"purchasingAccount"`
	PurchasingContact string `json:"purchasingContact"`
	LeaseDate         string `json:"leaseDate"`
	PODate            string `json:"poDate"`
	WarrantyDate      string `json:"warrantyDate"`
	VersionLock       int    `json:"versionLock"`
}

type ComputerPrestagesAccountSettings struct {
	ID                                      string `json:"id"`
	PayloadConfigured                       bool   `json:"payloadConfigured"`
	LocalAdminAccountEnabled                bool   `json:"localAdminAccountEnabled"`
	AdminUsername                           string `json:"adminUsername"`
	AdminPassword                           string `json:"adminPassword"`
	HiddenAdminAccount                      bool   `json:"hiddenAdminAccount"`
	LocalUserManaged                        bool   `json:"localUserManaged"`
	UserAccountType                         string `json:"userAccountType"`
	VersionLock                             int    `json:"versionLock"`
	PrefillPrimaryAccountInfoFeatureEnabled bool   `json:"prefillPrimaryAccountInfoFeatureEnabled"`
	PrefillType                             string `json:"prefillType"`
	PrefillAccountFullName                  string `json:"prefillAccountFullName"`
	PrefillAccountUserName                  string `json:"prefillAccountUserName"`
	PreventPrefillInfoFromModification      bool   `json:"preventPrefillInfoFromModification"`
}

// GetComputerPrestagesV3 retrieves all computer prestage information with optional sorting.
func (c *Client) GetComputerPrestagesV3(sort []string) (*ResponseComputerPrestagesV3, error) {
	const maxPageSize = 2000 // Assuming 2000 is a suitable limit for this API
	var allPrestages []ComputerPrestagesItem

	page := 0
	for {
		// Construct the endpoint with query parameters for the current page
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{strconv.Itoa(maxPageSize)},
		}
		if len(sort) > 0 {
			params.Add("sort", url.QueryEscape(strings.Join(sort, ",")))
		}
		endpointWithParams := fmt.Sprintf("%s?%s", uriComputerPrestagesV3, params.Encode())

		// Fetch the computer prestages for the current page
		var responsePrestagesV3 ResponseComputerPrestagesV3
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responsePrestagesV3)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch computer prestages v3: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched prestages to the total list
		allPrestages = append(allPrestages, responsePrestagesV3.Results...)

		// Check if all prestages have been fetched
		if responsePrestagesV3.TotalCount == nil || len(allPrestages) >= *responsePrestagesV3.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Create an int variable for the total count and assign its address to TotalCount
	totalCount := len(allPrestages)

	// Return the combined list of all computer prestages
	return &ResponseComputerPrestagesV3{
		TotalCount: &totalCount,
		Results:    allPrestages,
	}, nil
}

// GetDeviceScopeForComputerPrestage retrieves the device scope for a specific computer prestage by its ID.
func (c *Client) GetDeviceScopeForComputerPrestage(id string) (*ResponseDeviceScope, error) {
	endpoint := fmt.Sprintf("%s/%s/scope", uriComputerPrestagesV2, id)

	var deviceScope ResponseDeviceScope
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceScope)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch device scope for computer prestage with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceScope, nil
}

// GetComputerPrestageByID retrieves a specific computer prestage by its ID.
func (c *Client) GetComputerPrestageByID(id string) (*ComputerPrestagesItem, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputerPrestagesV3, id)

	var prestage ComputerPrestagesItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &prestage)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch computer prestage with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &prestage, nil
}

// GetComputerPrestageByName retrieves a specific computer prestage by its name.
func (c *Client) GetComputerPrestageByName(name string) (*ComputerPrestagesItem, error) {
	// Fetch all prestages
	response, err := c.GetComputerPrestagesV3([]string{}) // Adjust sort as needed
	if err != nil {
		return nil, fmt.Errorf("failed to fetch computer prestages: %v", err)
	}

	// Search for the prestage with the given name
	var prestageID string
	for _, prestage := range response.Results {
		if prestage.DisplayName == name {
			prestageID = prestage.ID
			break
		}
	}

	if prestageID == "" {
		return nil, fmt.Errorf("no computer prestage found with the name %s", name)
	}

	// Use the ID to get the full details of the prestage
	return c.GetComputerPrestageByID(prestageID)
}

// CreateComputerPrestage creates a new computer prestage with the given details.
func (c *Client) CreateComputerPrestage(prestage *ComputerPrestagesItem) (*ComputerPrestagesItem, error) {
	endpoint := uriComputerPrestagesV3

	var response ComputerPrestagesItem
	resp, err := c.HTTP.DoRequest("POST", endpoint, prestage, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create computer prestage: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateComputerPrestageByID updates a computer prestage by its ID.
func (c *Client) UpdateComputerPrestageByID(id string, prestageUpdate *ComputerPrestagesItem) (*ComputerPrestagesItem, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf("%s/%s", uriComputerPrestagesV3, id)

	var updatedPrestage ComputerPrestagesItem
	resp, err := c.HTTP.DoRequest("PUT", endpoint, prestageUpdate, &updatedPrestage)
	if err != nil {
		return nil, fmt.Errorf("failed to update computer prestage with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedPrestage, nil
}

// UpdateComputerPrestageByNameByID updates a computer prestage based on its display name.
func (c *Client) UpdateComputerPrestageByNameByID(name string, updatedPrestage *ComputerPrestagesItem) (*ComputerPrestagesItem, error) {
	// Fetch all prestages
	prestagesList, err := c.GetComputerPrestagesV3([]string{}) // Adjust sort as needed
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all computer prestages: %v", err)
	}

	// Search for the prestage with the given name
	for _, prestage := range prestagesList.Results {
		if prestage.DisplayName == name {
			// Update the prestage using its ID
			return c.UpdateComputerPrestageByID(prestage.ID, updatedPrestage)
		}
	}

	return nil, fmt.Errorf("no computer prestage found with the name %s", name)
}

// DeleteComputerPrestageByID deletes a computer prestage by its ID
func (c *Client) DeleteComputerPrestageByID(id string) error {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf("%s/%s", uriComputerPrestagesV3, id)

	// Perform the DELETE request
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete computer prestage with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerPrestageByNameByID deletes a computer prestage by its name.
func (c *Client) DeleteComputerPrestageByNameByID(name string) error {
	// Fetch all prestages to find the one with the given name.
	response, err := c.GetComputerPrestagesV3([]string{}) // Adjust sort as needed.
	if err != nil {
		return fmt.Errorf("failed to fetch computer prestages: %v", err)
	}

	// Find the prestage with the given name and delete it using its ID.
	for _, prestage := range response.Results {
		if prestage.DisplayName == name {
			return c.DeleteComputerPrestageByID(prestage.ID)
		}
	}

	return fmt.Errorf("no computer prestage found with the name %s", name)
}
