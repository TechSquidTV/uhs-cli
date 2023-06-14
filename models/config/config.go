package config

import "github.com/techsquidtv/uhs-cli/models/service"
import "github.com/techsquidtv/uhs-cli/models/service/servicemap"

type Config struct {
	Common   service.ServiceInterface `yaml:"common,omitempty"`
	Services service.ServicesConfig   `yaml:"services"`
}

func DefaultServiceConfig() service.ServicesConfig {
	uhsConfig := make(service.ServicesConfig)

	for serviceType, service := range servicemap.ServiceMap {
		uhsConfig[serviceType] = service.Default()
	}

	return uhsConfig
}
