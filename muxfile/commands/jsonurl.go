package commands

import (
	"fmt"

	"github.com/czlucius/traefik-dynamic-mux/dynamic"
	"github.com/czlucius/traefik-dynamic-mux/sources"
)

type JSONURLCommand struct{}

// Adheres to the CommandImpl interface
func (c *JSONURLCommand) Execute(args []string, config *dynamic.Configuration) error {
	if len(args) < 1 {
		return fmt.Errorf("JSONURL command requires a URL argument")
	}

	var jsonURLSource sources.JSONUrlSource = sources.JSONUrlSource{
		Url: args[0],
	}
	remoteConfig, err := jsonURLSource.PassConfig()
	if err != nil {
		return fmt.Errorf("failed to pass config: %w", err)
	}

	// Merge the remote configuration into the existing configuration
	sources.MergeConfig(config, remoteConfig)

	return nil
}
