package jellyfin

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Jellyfin struct {
	service.Service `yaml:",inline"`
	JellyfinOptions  `yaml:",inline"`
}

type JellyfinOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func (s *Jellyfin) Default() service.ServiceInterface {
	http := 8096
	udp := 7359

	p := &Jellyfin{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "ghcr.io/onedr0p/jellyfin",
				Tag:        "rolling",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
				Udp:  &udp,
			},
		},
		JellyfinOptions: JellyfinOptions{
			Config:  "/opt/jellyfin/config",
			Library: "/data/library",
		},
	}
	return p
}

func (s *Jellyfin) Configure() service.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your jellyfin config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputLibraryPath := &survey.Input{
		Message: "Enter the path to your jellyfin library folder:",
		Default: s.Library,
	}

	err = survey.AskOne(inputLibraryPath, &s.Library)
	if err != nil {
		fmt.Println(err.Error())
	}

	return s
}
