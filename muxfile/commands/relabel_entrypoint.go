package commands

import (
	"fmt"

	"github.com/czlucius/traefik-dynamic-mux/dynamic"
)

type RelabelEntryPointCommand struct{}

// Adheres to the CommandImpl interface
func (c *RelabelEntryPointCommand) Execute(args []string, config *dynamic.Configuration) error {
	if len(args) < 2 {
		return fmt.Errorf("RelabelEntryPoint command requires 2 arguments (old and new entry point names)")
	}

	// Replace args[0] with args[1] for all entry points
	// HTTP
	for key, value := range config.HTTP.Routers {
		if value.EntryPoints == nil {
			continue
		}
		for i, entryPoint := range value.EntryPoints {
			if entryPoint == args[0] {
				value.EntryPoints[i] = args[1]
			}
		}
		config.HTTP.Routers[key] = value
	}
	// TCP
	for key, value := range config.TCP.Routers {
		if value.EntryPoints == nil {
			continue
		}
		for i, entryPoint := range value.EntryPoints {
			if entryPoint == args[0] {
				value.EntryPoints[i] = args[1]
			}
		}
		config.TCP.Routers[key] = value
	}

	// UDP
	for key, value := range config.UDP.Routers {
		if value.EntryPoints == nil {
			continue
		}
		for i, entryPoint := range value.EntryPoints {
			if entryPoint == args[0] {
				value.EntryPoints[i] = args[1]
			}
		}
		config.UDP.Routers[key] = value
	}

	return nil
}
