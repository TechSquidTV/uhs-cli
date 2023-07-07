package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type playwright struct {
	service.Service   `yaml:",inline"`
	playwrightOptions `yaml:",inline"`
}

type playwrightOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewPlaywright() manager.Configurer {
	http := 3000

	return &playwright{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "browserless/chrome",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		playwrightOptions: playwrightOptions{
			Config: "/opt/playwright/config",
		},
	}
}

func (s *playwright) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Playwright config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *playwright) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
