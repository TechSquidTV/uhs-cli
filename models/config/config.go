package config

import (
	"fmt"

	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service/servicemap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Global   manager.Configurer `yaml:"global,omitempty"`
	Services ServicesConfig     `yaml:"services"`
}

type ServicesConfig map[string]manager.Configurer

func (sc ServicesConfig) UnmarshalYAML(unmarshal func(any) error) error {
	services := make(map[string]any)
	if err := unmarshal(&services); err != nil {
		fmt.Println("could not unmarshal:", err)
		return err
	}
	for k, v := range services {
		b, err := yaml.Marshal(v) // remarshal this specific service so we can unmarshal it properly.
		if err != nil {
			fmt.Println("Error marshalling:", err)
			continue
		}
		service, err := unmarshalService(k, b)
		if err != nil {
			fmt.Println("Failed to unmarshal service", err)
			continue
		}
		sc[k] = service
	}
	return nil
}

func unmarshalService(key string, data []byte) (manager.Configurer, error) {
	v, ok := servicemap.Registered[key]
	if !ok {
		return nil, fmt.Errorf("unregistered service name: %q", key)
	}
	sc := v()
	if err := sc.Fill(data); err != nil {
		return nil, err
	}
	return sc, nil
}
