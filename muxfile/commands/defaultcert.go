package commands

import (
	"fmt"

	"czlucius.dev/tmap/dynamic"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
)

type DefaultCertCommand struct{}

// Adheres to the CommandImpl interface
func (c *DefaultCertCommand) Execute(args []string, config *dynamic.Configuration) error {
	if len(args) < 2 {
		return fmt.Errorf("DefaultCert command requires a certificate file and key file")
	}

	defaultStore, ok := config.TLS.Stores["default"]
	if !ok {
		// Create a default store if it doesn't exist
		config.TLS.Stores["default"] = tls.Store{
			DefaultCertificate: &tls.Certificate{},
		}
		// defaultStore is actually assigned but undefined (because of the error)
		defaultStore = config.TLS.Stores["default"]
	}

	defaultStore.DefaultCertificate.CertFile = types.FileOrContent(args[0])
	defaultStore.DefaultCertificate.KeyFile = types.FileOrContent(args[1])
	// Set the default store in the configuration
	config.TLS.Stores["default"] = defaultStore
	return nil
}
