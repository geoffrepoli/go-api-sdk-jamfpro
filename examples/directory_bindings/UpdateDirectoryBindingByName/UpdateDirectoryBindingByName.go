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
	// Directory binding data to update
	updateBinding := &jamfpro.ResponseDirectoryBinding{
		Name:       "Updated Binding",
		Priority:   1,
		Domain:     "updated.example.com",
		Username:   "user@updated.com",
		Password:   "newpassword",
		ComputerOU: "CN=UpdatedComputers,DC=updated,DC=example,DC=com",
		Type:       "Active Directory",
	}

	// Update directory binding by Name
	bindingName := "New Binding" // Assuming an existing binding name
	updatedBindingByName, err := client.UpdateDirectoryBindingByName(bindingName, updateBinding)
	if err != nil {
		fmt.Println("Error updating directory binding by Name:", err)
		return
	}
	updatedBindingByNameXML, _ := xml.MarshalIndent(updatedBindingByName, "", "    ")
	fmt.Printf("Updated Directory Binding by Name:\n%s\n", string(updatedBindingByNameXML))
}
