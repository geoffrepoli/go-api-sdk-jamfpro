package main

import (
	"encoding/xml"
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
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the updated advanced user search details
	updatedAdvancedUserSearch := &jamfpro.ResourceAdvancedUserSearch{
		Name: "Advanced User Search Name by jamf pro sdk",
		Criteria: jamfpro.SharedContainerCriteria{
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Email Address",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "like",
					Value:        "company.com",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.SharedAdvancedSearchContainerDisplayField{
			{
				DisplayField: []jamfpro.SharedAdvancedSearchSubsetDisplayField{
					{
						Name: "Computers",
					},
					{
						Name: "Content Name",
					},
					{
						Name: "Roster Course Source",
					},
					// Additional display fields can be added here
				},
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	searchID := 17 // Replace 123 with the actual ID

	// Update by ID
	updatedByID, err := client.UpdateAdvancedUserSearchByID(searchID, updatedAdvancedUserSearch)
	if err != nil {
		log.Fatalf("Error updating advanced user search by ID: %v", err)
	}
	// Print updated search by ID
	updatedByIDXML, err := xml.MarshalIndent(updatedByID, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated advanced user search by ID: %v", err)
	}
	fmt.Println("Updated Advanced User Search by ID:\n", string(updatedByIDXML))
}
