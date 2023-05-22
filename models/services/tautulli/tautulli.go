package tautulli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Tautulli struct {
	services.Service `yaml:",inline"`
	TautulliOptions  `yaml:",inline"`
}

type TautulliOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Tautulli) Default() services.ServiceInterface {
	http := 8181

	p := &Tautulli{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "lscr.io/linuxserver/tautulli",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		TautulliOptions: TautulliOptions{
			Config: "/opt/tautulli/config",
		},
	}
	return p
}

func (s *Tautulli) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Tautulli config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
