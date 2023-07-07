package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type autobrr struct {
	service.Service `yaml:",inline"`
	autobrrOptions  `yaml:",inline"`
}

type autobrrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewAutobrr() manager.Configurer {
	http := 7474

	return &autobrr{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "ghcr.io/autobrr/autobrr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		autobrrOptions: autobrrOptions{
			Config: "/opt/autobrr/config",
		},
	}
}

func (s *autobrr) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Autobrr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *autobrr) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
