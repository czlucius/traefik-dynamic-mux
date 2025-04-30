package sources

import (
	"github.com/czlucius/traefik-dynamic-mux/dynamic"
	"github.com/traefik/traefik/v3/pkg/tls"
)

type Merge interface {
	MergeConfig(config1 *dynamic.Configuration, config2 *dynamic.Configuration) *dynamic.Configuration
}

func MergeConfig(config1 *dynamic.Configuration, config2 *dynamic.Configuration) {
	if config2 == nil {
		return
	}

	// Merge the HTTP configuration
	if config1.HTTP == nil {
		config1.HTTP = &dynamic.HTTPConfiguration{}
	}
	if config2.HTTP != nil {
		// Merge the routers configuration
		if config1.HTTP.Routers == nil {
			config1.HTTP.Routers = make(map[string]*dynamic.Router)
		}
		if config2.HTTP.Routers != nil {
			for key, value := range config2.HTTP.Routers {
				config1.HTTP.Routers[key] = value
			}
		}
		// Merge the services configuration
		if config1.HTTP.Services == nil {
			config1.HTTP.Services = make(map[string]*dynamic.Service)
		}
		if config2.HTTP.Services != nil {
			for key, value := range config2.HTTP.Services {
				config1.HTTP.Services[key] = value
			}
		}
		// Merge the middlewares configuration
		if config1.HTTP.Middlewares == nil {
			config1.HTTP.Middlewares = make(map[string]*dynamic.Middleware)
		}
		if config2.HTTP.Middlewares != nil {
			for key, value := range config2.HTTP.Middlewares {
				config1.HTTP.Middlewares[key] = value
			}
		}
	}

	// Merge the TCP configuration
	if config1.TCP == nil {
		config1.TCP = &dynamic.TCPConfiguration{}
	}
	if config2.TCP != nil {
		// Merge the routers configuration for TCP
		if config1.TCP.Routers == nil {
			config1.TCP.Routers = make(map[string]*dynamic.TCPRouter)
		}
		if config2.TCP.Routers != nil {
			for key, value := range config2.TCP.Routers {
				config1.TCP.Routers[key] = value
			}
		}
		// Merge the services configuration for TCP
		if config1.TCP.Services == nil {
			config1.TCP.Services = make(map[string]*dynamic.TCPService)
		}
		if config2.TCP.Services != nil {
			for key, value := range config2.TCP.Services {
				config1.TCP.Services[key] = value
			}
		}
	}

	// Merge the UDP configuration
	if config1.UDP == nil {
		config1.UDP = &dynamic.UDPConfiguration{}
	}
	if config2.UDP != nil {
		// Merge the routers configuration for UDP
		if config1.UDP.Routers == nil {
			config1.UDP.Routers = make(map[string]*dynamic.UDPRouter)
		}
		if config2.UDP.Routers != nil {
			for key, value := range config2.UDP.Routers {
				config1.UDP.Routers[key] = value
			}
		}
		// Merge the services configuration for UDP
		if config1.UDP.Services == nil {
			config1.UDP.Services = make(map[string]*dynamic.UDPService)
		}
		if config2.UDP.Services != nil {
			for key, value := range config2.UDP.Services {
				config1.UDP.Services[key] = value
			}
		}
	}

	// Merge the TLS configuration
	if config1.TLS == nil {
		config1.TLS = &dynamic.TLSConfiguration{}
	}
	if config2.TLS != nil {
		// Merge the certificates configuration
		if config1.TLS.Certificates == nil {
			config1.TLS.Certificates = make([]*tls.CertAndStores, 0)
		}
		if config2.TLS.Certificates != nil {
			config1.TLS.Certificates = append(config1.TLS.Certificates, config2.TLS.Certificates...)
		}

		// Merge the options configuration
		if config1.TLS.Options == nil {
			config1.TLS.Options = make(map[string]tls.Options)
		}
		if config2.TLS.Options != nil {
			for key, value := range config2.TLS.Options {
				config1.TLS.Options[key] = value
			}
		}
		// Merge the stores configuration
		if config1.TLS.Stores == nil {
			config1.TLS.Stores = make(map[string]tls.Store)
		}
		if config2.TLS.Stores != nil {
			for key, value := range config2.TLS.Stores {
				config1.TLS.Stores[key] = value
			}
		}
	}
}
