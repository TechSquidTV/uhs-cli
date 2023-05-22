package playwright

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Playwright struct {
	services.Service  `yaml:",inline"`
	PlaywrightOptions `yaml:",inline"`
}

type PlaywrightOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func (s *Playwright) Default() services.ServiceInterface {
	http := 3000

	p := &Playwright{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "browserless/chrome",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		PlaywrightOptions: PlaywrightOptions{
			Config: "/opt/playwright/config",
		},
	}
	return p
}

func (s *Playwright) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Playwright config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
