package plex

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Plex struct {
	service.Service `yaml:",inline"`
	PlexOptions      `yaml:",inline"`
}

type PlexOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func (s *Plex) Default() service.ServiceInterface {
	http := 32400

	p := &Plex{
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
		PlexOptions: PlexOptions{
			Config:  "/opt/plex/config",
			Library: "/data/library",
		},
	}
	return p
}

func (s *Plex) Configure() service.ServiceInterface {
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
	return s
}
