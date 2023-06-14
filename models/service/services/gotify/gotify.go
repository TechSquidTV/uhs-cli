package gotify

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Gotify struct {
	service.Service `yaml:",inline"`
	GotifyOptions    `yaml:",inline"`
}

type GotifyOptions struct {
	Data string `yaml:"data"`
}

// Return default values for service
func (s *Gotify) Default() service.ServiceInterface {
	http := 80

	p := &Gotify{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "ghcr.io/gotify/server",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		GotifyOptions: GotifyOptions{
			Data: "/opt/gotify/data",
		},
	}
	return p
}

func (s *Gotify) Configure() service.ServiceInterface {
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
