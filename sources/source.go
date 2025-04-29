package sources

import "czlucius.dev/tmap/dynamic" // Replace with the correct path to the config package

type Source interface {
	PassConfig() *dynamic.Configuration
}
