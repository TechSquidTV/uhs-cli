package overseerr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Overseerr struct {
	service.Service `yaml:",inline"`
	OverseerrOptions `yaml:",inline"`
}

type OverseerrOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Overseerr) Default() service.ServiceInterface {
	http := 5055

	p := &Overseerr{
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
		OverseerrOptions: OverseerrOptions{
			Config: "/opt/overseerr/config",
		},
	}
	return p
}

func (s *Overseerr) Configure() service.ServiceInterface {
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
