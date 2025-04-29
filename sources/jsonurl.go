package sources

import (
	"encoding/json"
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

	var remoteConfig dynamic.Configuration

	// Parse the JSON response into the dynamic.Configuration struct
	if err := json.NewDecoder(resp.Body).Decode(&remoteConfig); err != nil {
		// Handle the error (e.g., log it or return the original config)
		return nil, err
	}

	// Return the parsed configuration
	return &remoteConfig, nil
}
