package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type gotify struct {
	service.Service `yaml:",inline"`
	gotifyOptions   `yaml:",inline"`
}

type gotifyOptions struct {
	Data string `yaml:"data"`
}

// Return default values for service
func NewGotify() manager.Configurer {
	http := 80

	return &gotify{
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
		gotifyOptions: gotifyOptions{
			Data: "/opt/gotify/data",
		},
	}
}

func (s *gotify) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Gotify data folder:",
		Default: s.Data,
	}
	err := survey.AskOne(inputConfigPath, &s.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *gotify) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
