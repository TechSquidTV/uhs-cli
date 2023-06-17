package common

import (
	"github.com/stretchr/testify/assert"
	configCommon "github.com/techsquidtv/uhs-cli/models/common"
	"github.com/techsquidtv/uhs-cli/models/config"
	"github.com/techsquidtv/uhs-cli/models/service"
	"github.com/techsquidtv/uhs-cli/models/service/servicemap"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestInput(t *testing.T) {
	// Generate sample config for a service
	uhsConfig := config.Config{
		Common:   new(configCommon.Common).Default(),
		Services: make(service.ServicesConfig),
	}
	// Add a service to the config
	uhsConfig.Services["plex"] = servicemap.ServiceMap["plex"].Default()
	// Marshal the config to yaml
	yamlConfig, err := yaml.Marshal(uhsConfig)
	assert.NoError(t, err)
	// Unmarshal the yaml back into a config
	err = yaml.Unmarshal(yamlConfig, &uhsConfig)
	assert.NoError(t, err)
}
