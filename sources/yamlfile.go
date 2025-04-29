package sources

import (
	"fmt"

	"czlucius.dev/tmap/dynamic"
	"io"
	"os"
	"sigs.k8s.io/yaml"
)

type YAMLFileSource struct {
	FilePath string
}

func (y *YAMLFileSource) PassConfig() (*dynamic.Configuration, error) {
	// Read the YAML file
	yamlFile, err := os.Open(y.FilePath)
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()

	// Read the content of the YAML file
	yamlContent, err := io.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML content into a dynamic.Configuration struct
	var config dynamic.Configuration
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		return nil, err
	}
	fmt.Println("YAML file content:", config.HTTP.Routers["hd-tls"])

	return &config, nil
}
