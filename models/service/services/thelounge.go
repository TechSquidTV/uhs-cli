package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type thelounge struct {
	service.Service  `yaml:",inline"`
	theloungeOptions `yaml:",inline"`
}

type theloungeOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewThelounge() manager.Configurer {
	http := 9000

	return &thelounge{
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
		theloungeOptions: theloungeOptions{
			Config: "/opt/thelounge/config",
		},
	}
}

func (s *thelounge) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Thelounge config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *thelounge) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
