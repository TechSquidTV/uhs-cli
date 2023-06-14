package changedetectionio

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Changedetectionio struct {
	service.Service         `yaml:",inline"`
	ChangedetectionioOptions `yaml:",inline"`
}

type ChangedetectionioOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Changedetectionio) Default() service.ServiceInterface {
	http := 5000

	p := &Changedetectionio{
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
		ChangedetectionioOptions: ChangedetectionioOptions{
			Config: "/opt/changedetectionio/config",
		},
	}
	return p
}

func (s *Changedetectionio) Configure() service.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Changedetectionio config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
