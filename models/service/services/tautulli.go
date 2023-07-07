package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type tautulli struct {
	service.Service `yaml:",inline"`
	tautulliOptions `yaml:",inline"`
}

type tautulliOptions struct {
	Config string `yaml:"config"`
}

// Return default values for service
func NewTautulli() manager.Configurer {
	http := 8181

	return &tautulli{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/tautulli",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		tautulliOptions: tautulliOptions{
			Config: "/opt/tautulli/config",
		},
	}
}

func (s *tautulli) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Tautulli config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *tautulli) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
