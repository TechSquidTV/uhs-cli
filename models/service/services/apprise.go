package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type apprise struct {
	service.Service `yaml:",inline"`
	appriseOptions  `yaml:",inline"`
}

type appriseOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewApprise() manager.Configurer {
	http := 8000

	return &apprise{
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
		appriseOptions: appriseOptions{
			Config: "/opt/apprise/config",
		},
	}
}

func (s *apprise) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Apprise config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *apprise) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
