package thelounge

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Thelounge struct {
	service.Service  `yaml:",inline"`
	TheloungeOptions `yaml:",inline"`
}

type TheloungeOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Thelounge) Default() service.ServiceInterface {
	http := 9000

	p := &Thelounge{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/thelounge",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		TheloungeOptions: TheloungeOptions{
			Config: "/opt/thelounge/config",
		},
	}
	return p
}

func (s *Thelounge) Configure() service.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Thelounge config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
