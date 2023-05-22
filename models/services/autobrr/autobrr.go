package autobrr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Autobrr struct {
	services.Service `yaml:",inline"`
	AutobrrOptions   `yaml:",inline"`
}

type AutobrrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Autobrr) Default() services.ServiceInterface {
	http := 7474

	p := &Autobrr{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "ghcr.io/autobrr/autobrr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		AutobrrOptions: AutobrrOptions{
			Config: "/opt/autobrr/config",
		},
	}
	return p
}

func (s *Autobrr) Configure() services.ServiceInterface {
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
