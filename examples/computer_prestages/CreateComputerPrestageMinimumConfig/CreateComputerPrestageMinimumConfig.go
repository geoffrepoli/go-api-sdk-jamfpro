package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the payload for creating a new computer prestage
	// Manually create a ResourceComputerPrestage struct with mapped values
	prestage := jamfpro.ResourceComputerPrestage{
		DisplayName:                       "jamfpro-sdk-example-computerPrestage-config",
		Mandatory:                         true,
		MDMRemovable:                      false,
		SupportPhoneNumber:                "",
		SupportEmailAddress:               "",
		Department:                        "",
		DefaultPrestage:                   false,
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        false,
		KeepExistingLocationInformation:   false,
		RequireAuthentication:             false,
		AuthenticationPrompt:              "",
		PreventActivationLock:             true,
		EnableDeviceBasedActivationLock:   false,
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: map[string]bool{
			"Accessibility":     true,
			"Appearance":        true,
			"AppleID":           true,
			"Biometric":         true,
			"Diagnostics":       true,
			"DisplayTone":       true,
			"FileVault":         false,
			"Location":          false,
			"Payment":           true,
			"Privacy":           true,
			"Registration":      true,
			"Restore":           true,
			"ScreenTime":        true,
			"Siri":              true,
			"TOS":               true,
			"TermsOfAddress":    true,
			"iCloudDiagnostics": true,
			"iCloudStorage":     true,
		},
		LocationInformation: jamfpro.ComputerPrestageSubsetLocationInformation{
			Username:     "",
			Realname:     "",
			Phone:        "",
			Email:        "",
			Room:         "",
			Position:     "",
			DepartmentId: "-1",
			BuildingId:   "-1",
			ID:           "1",
			VersionLock:  0,
		},
		PurchasingInformation: jamfpro.ComputerPrestageSubsetPurchasingInformation{
			ID:                "1",
			Leased:            false,
			Purchased:         true,
			AppleCareId:       "",
			PONumber:          "",
			Vendor:            "",
			PurchasePrice:     "",
			LifeExpectancy:    0,
			PurchasingAccount: "",
			PurchasingContact: "",
			LeaseDate:         "1970-01-01",
			PODate:            "1970-01-01",
			WarrantyDate:      "1970-01-01",
			VersionLock:       0,
		},
		AnchorCertificates:               []string{},
		EnrollmentCustomizationId:        "0",
		Language:                         "en",
		Region:                           "GB",
		AutoAdvanceSetup:                 true,
		InstallProfilesDuringSetup:       true,
		PrestageInstalledProfileIds:      []string{},
		CustomPackageIds:                 []string{},
		CustomPackageDistributionPointId: "-1",
		EnableRecoveryLock:               false,
		RecoveryLockPasswordType:         "MANUAL",
		RotateRecoveryLockPassword:       false,
		ID:                               "1",
		ProfileUuid:                      "C101330EE870D6082D5D08FA013ADE51",
		SiteId:                           "-1",
		VersionLock:                      3,
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			ID:                                      "1",
			PayloadConfigured:                       true,
			LocalAdminAccountEnabled:                true,
			AdminUsername:                           "localAdmin",
			HiddenAdminAccount:                      true,
			LocalUserManaged:                        false,
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             1,
			PrefillPrimaryAccountInfoFeatureEnabled: false,
			PrefillType:                             "CUSTOM",
			PrefillAccountFullName:                  "",
			PrefillAccountUserName:                  "",
			PreventPrefillInfoFromModification:      false,
		},
	}

	// Call the CreateComputerPrestage function
	createdPrestage, err := client.CreateComputerPrestage(&prestage)
	if err != nil {
		log.Fatalf("Error creating computer prestage: %v", err)
	}

	// Pretty print the computer prestage in JSON
	prestageJSON, err := json.MarshalIndent(createdPrestage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling computer prestage data: %v", err)
	}
	fmt.Println("Created computer prestage:\n", string(prestageJSON))
}
