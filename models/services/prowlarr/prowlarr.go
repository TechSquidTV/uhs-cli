package prowlarr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Prowlarr struct {
	services.Service `yaml:",inline"`
	ProwlarrOptions      `yaml:",inline"`
}

type ProwlarrOptions struct {
	Config  string `yaml:"config"`
}

// Return default values for service
func (s *Prowlarr) Default() services.ServiceInterface {
	http := 9696

	p := &Prowlarr{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "linuxserver/prowlarr",
				Tag:        "latest",
				PullPolicy: "IfNotPresent",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		ProwlarrOptions: ProwlarrOptions{
			Config:  "/opt/prowlarr/config",
		},
	}
	return p
}

func (s *Prowlarr) Configure() services.ServiceInterface {
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
