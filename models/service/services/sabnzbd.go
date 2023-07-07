package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type sabnzbd struct {
	service.Service `yaml:",inline"`
	sabnzbdOptions  `yaml:",inline"`
}

type sabnzbdOptions struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
}

// Return default values for service
func NewSabnzbd() manager.Configurer {
	http := 8080

	// Create a new Sabnzbd instance
	return &sabnzbd{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/sabnzbd",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		sabnzbdOptions: sabnzbdOptions{
			Config: "/opt/sabnzbd/config",
			Data:   "/data/usenet",
		},
	}
}

func (s *sabnzbd) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your sabnzbd config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDataPath := &survey.Input{
		Message: "Enter the path to your usenet data folder:",
		Default: s.Data,
	}
	err = survey.AskOne(inputDataPath, &s.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *sabnzbd) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
