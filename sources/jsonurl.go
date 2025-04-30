package sources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"czlucius.dev/tmap/dynamic"
)

// Define JSONUrlSource as a struct that implements the Source interface
type JSONUrlSource struct {
	Url string
}

// PassConfig method to pass the configuration to the JSONUrlSource
func (j *JSONUrlSource) PassConfig() (*dynamic.Configuration, error) {
	// Make a request to the URL and get the JSON data
	resp, err := http.Get(j.Url)
	if err != nil {
		// Handle the error (e.g., log it or return the original config)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("JSON URL response status:", resp.Status)
	var remoteConfig dynamic.Configuration
	body := resp.Body

	// Parse the JSON response into the dynamic.Configuration struct
	if err := json.NewDecoder(body).Decode(&remoteConfig); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		// Handle the error (e.g., log it or return the original config)
		return nil, err
	}

	fmt.Println("JSON URL content:", remoteConfig)

	// Return the parsed configuration
	return &remoteConfig, nil
}
