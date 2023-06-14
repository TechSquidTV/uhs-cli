package apprise

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Apprise struct {
	service.Service `yaml:",inline"`
	AppriseOptions   `yaml:",inline"`
}

type AppriseOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Apprise) Default() service.ServiceInterface {
	http := 8000

	p := &Apprise{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/apprise-api",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		AppriseOptions: AppriseOptions{
			Config: "/opt/apprise/config",
		},
	}
	return p
}

func (s *Apprise) Configure() service.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Apprise config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
