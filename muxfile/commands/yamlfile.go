package commands

import (
	"fmt"

	"github.com/czlucius/traefik-dynamic-mux/dynamic"
	"github.com/czlucius/traefik-dynamic-mux/sources"
)

type YAMLFileCommand struct{}

// Adheres to the CommandImpl interface

func (c *YAMLFileCommand) Execute(args []string, config *dynamic.Configuration) error {

	if len(args) < 1 {
		return fmt.Errorf("YAMLFile command requires a file path argument")
	}

	var yamlFileSource sources.YAMLFileSource = sources.YAMLFileSource{
		FilePath: args[0],
	}
	remoteConfig, err := yamlFileSource.PassConfig()
	if err != nil {
		return fmt.Errorf("failed to pass config: %w", err)
	}

	// Merge the remote configuration into the existing configuration
	sources.MergeConfig(config, remoteConfig)

	return nil
}
