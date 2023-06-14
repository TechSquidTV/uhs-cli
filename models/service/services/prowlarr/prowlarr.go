package prowlarr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Prowlarr struct {
	service.Service `yaml:",inline"`
	ProwlarrOptions  `yaml:",inline"`
}

type ProwlarrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Prowlarr) Default() service.ServiceInterface {
	http := 9696

	p := &Prowlarr{
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
		ProwlarrOptions: ProwlarrOptions{
			Config: "/opt/prowlarr/config",
		},
	}
	return p
}

func (s *Prowlarr) Configure() service.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your prowlarr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
