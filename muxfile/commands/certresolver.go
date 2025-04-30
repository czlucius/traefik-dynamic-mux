package commands

import (
	"fmt"
	"slices"

	"github.com/czlucius/traefik-dynamic-mux/dynamic"
)

type CertResolverCommand struct{}

// Adheres to the CommandImpl interface
func (c *CertResolverCommand) Execute(args []string, config *dynamic.Configuration) error {
	if len(args) < 2 {
		return fmt.Errorf("CertResolver command requires a resolver name, an entryPoint name and optional override flag")
	}
	// 2nd argument is optional, which is a boolean
	override := false
	if len(args) == 3 {
		if args[2] == "override" {
			override = true
		}
	}

	// Applies the resolver to all HTTP routers
	for key, value := range config.HTTP.Routers {
		if value.GetTLS() != nil && !override {
			// We do not want to override existing TLS configurations
			continue
		}
		if slices.Contains(value.EntryPoints, args[1]) {
			config.HTTP.Routers[key].SetTLS(&dynamic.RouterTLSConfig{
				CertResolver: args[0],
			})
		}
	}

	return nil
}
