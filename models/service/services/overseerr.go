package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type overseerr struct {
	service.Service  `yaml:",inline"`
	overseerrOptions `yaml:",inline"`
}

type overseerrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewOverseerr() manager.Configurer {
	http := 5055

	return &overseerr{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/overseerr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		overseerrOptions: overseerrOptions{
			Config: "/opt/overseerr/config",
		},
	}
}

func (s *overseerr) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your overseerr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *overseerr) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
