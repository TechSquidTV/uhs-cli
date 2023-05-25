package plex

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Plex struct {
	services.Service `yaml:",inline"`
	PlexOptions      `yaml:",inline"`
}

type PlexOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func (s *Plex) Default() services.ServiceInterface {
	http := 32400

	p := &Plex{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "lscr.io/linuxserver/plex",
				Tag:        "latest",
				PullPolicy: "IfNotPresent",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		PlexOptions: PlexOptions{
			Config:  "/opt/plex/config",
			Library: "/opt/plex/library",
		},
	}
	return p
}

func (s *Plex) Configure() services.ServiceInterface {
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
