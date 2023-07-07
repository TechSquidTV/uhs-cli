package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type prowlarr struct {
	service.Service `yaml:",inline"`
	prowlarrOptions `yaml:",inline"`
}

type prowlarrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewProwlarr() manager.Configurer {
	http := 9696

	return &prowlarr{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/prowlarr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		prowlarrOptions: prowlarrOptions{
			Config: "/opt/prowlarr/config",
		},
	}
}

func (s *prowlarr) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your prowlarr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *prowlarr) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
