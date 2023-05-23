package thelounge

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Thelounge struct {
	services.Service `yaml:",inline"`
	TheloungeOptions `yaml:",inline"`
}

type TheloungeOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Thelounge) Default() services.ServiceInterface {
	http := 9000

	p := &Thelounge{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "lscr.io/linuxserver/thelounge",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		TheloungeOptions: TheloungeOptions{
			Config: "/opt/thelounge/config",
		},
	}
	return p
}

func (s *Thelounge) Configure() services.ServiceInterface {
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
