package gotify

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Gotify struct {
	services.Service `yaml:",inline"`
	GotifyOptions    `yaml:",inline"`
}

type GotifyOptions struct {
	Data string `yaml:"data"`
}

// Return default values for service
func (s *Gotify) Default() services.ServiceInterface {
	http := 80

	p := &Gotify{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "ghcr.io/gotify/server",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		GotifyOptions: GotifyOptions{
			Data: "/opt/gotify/data",
		},
	}
	return p
}

func (s *Gotify) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Gotify data folder:",
		Default: s.Data,
	}
	err := survey.AskOne(inputConfigPath, &s.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
