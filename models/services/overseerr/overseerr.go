package overseerr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Overseerr struct {
	services.Service `yaml:",inline"`
	OverseerrOptions `yaml:",inline"`
}

type OverseerrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Overseerr) Default() services.ServiceInterface {
	http := 5055

	p := &Overseerr{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "linuxserver/overseerr",
				Tag:        "latest",
				PullPolicy: "IfNotPresent",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		OverseerrOptions: OverseerrOptions{
			Config: "/opt/overseerr/config",
		},
	}
	return p
}

func (s *Overseerr) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your overseerr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}