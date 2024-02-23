package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:            loadedConfig.ClientOptions.LogLevel,
			LogOutputFormat:     loadedConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator: loadedConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:   loadedConfig.ClientOptions.HideSensitiveData,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Set up new settings
	newSettings := jamfpro.ResourceComputerCheckin{
		CheckInFrequency:                 15, // Values can be 60, 30, 15 and  5
		CreateStartupScript:              true,
		LogStartupEvent:                  true,
		CheckForPoliciesAtStartup:        true,
		ApplyComputerLevelManagedPrefs:   true,
		EnsureSSHIsEnabled:               false,
		CreateLoginLogoutHooks:           true,
		LogUsername:                      true,
		CheckForPoliciesAtLoginLogout:    true,
		ApplyUserLevelManagedPreferences: true,
		HideRestorePartition:             false,
		PerformLoginActionsInBackground:  true,
		DisplayStatusToUser:              false,
	}

	// Update computer check-in settings
	err = client.UpdateComputerCheckinInformation(&newSettings)
	if err != nil {
		fmt.Printf("Error updating computer check-in settings: %s\n", err)
		return
	}

	fmt.Println("Computer check-in settings updated successfully.")
}
