package main

import (
	"encoding/xml"
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
	logLevel := http_client.LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Let's assume you want to get the disk encryption configuration with name "Corporate Encryption"
	configName := "Corporate Encryption"
	configuration, err := client.GetDiskEncryptionConfigurationByName(configName)
	if err != nil {
		log.Fatalf("Error fetching disk encryption configuration by name: %v", err)
	}

	// Print the configuration in a pretty XML format (assuming the response is XML)
	configXML, err := xml.MarshalIndent(configuration, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling configuration data: %v", err)
	}
	fmt.Printf("Fetched Disk Encryption Configuration by Name:\n%s\n", configXML)
}
