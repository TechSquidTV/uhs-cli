package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type jellyfin struct {
	service.Service `yaml:",inline"`
	jellyfinOptions `yaml:",inline"`
}

type jellyfinOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func NewJellyfin() manager.Configurer {
	http := 8096
	udp := 7359

	return &jellyfin{
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
		jellyfinOptions: jellyfinOptions{
			Config:  "/opt/jellyfin/config",
			Library: "/data/library",
		},
	}
}

func (s *jellyfin) Configure() {
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
}

func (s *jellyfin) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
