package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type plex struct {
	service.Service `yaml:",inline"`
	plexOptions     `yaml:",inline"`
}

type plexOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func NewPlex() manager.Configurer {
	http := 32400

	return &plex{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/plex",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		plexOptions: plexOptions{
			Config:  "/opt/plex/config",
			Library: "/data/library",
		},
	}
}

func (s *plex) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your plex config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputLibraryPath := &survey.Input{
		Message: "Enter the path to your plex library folder:",
		Default: s.Library,
	}
	err = survey.AskOne(inputLibraryPath, &s.Library)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *plex) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
