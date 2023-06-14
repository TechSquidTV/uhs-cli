package autobrr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Autobrr struct {
	service.Service `yaml:",inline"`
	AutobrrOptions   `yaml:",inline"`
}

type AutobrrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Autobrr) Default() service.ServiceInterface {
	http := 7474

	p := &Autobrr{
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
		AutobrrOptions: AutobrrOptions{
			Config: "/opt/autobrr/config",
		},
	}
	return p
}

func (s *Autobrr) Configure() service.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Autobrr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
