package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type changedetectionio struct {
	service.Service          `yaml:",inline"`
	changedetectionioOptions `yaml:",inline"`
}

type changedetectionioOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewChangedetectionio() manager.Configurer {
	http := 5000

	return &changedetectionio{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "ghcr.io/dgtlmoon/changedetection.io",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		changedetectionioOptions: changedetectionioOptions{
			Config: "/opt/changedetectionio/config",
		},
	}
}

func (s *changedetectionio) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Changedetectionio config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *changedetectionio) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
