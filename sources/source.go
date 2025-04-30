package sources

import "github.com/czlucius/traefik-dynamic-mux/dynamic" // Replace with the correct path to the config package

type Source interface {
	PassConfig() *dynamic.Configuration
}
