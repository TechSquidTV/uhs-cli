package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type radarr struct {
	service.Service `yaml:",inline"`
	radarrOptions   `yaml:",inline"`
}

type radarrOptions struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
}

// Return default values for service
func NewRadarr() manager.Configurer {
	http := 7878

	return &radarr{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/radarr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		radarrOptions: radarrOptions{
			Config: "/opt/radarr/config",
			Data:   "/data",
		},
	}
}

func (s *radarr) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your radarr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDataPath := &survey.Input{
		Message: "Enter the path to your top level media folder:",
		Default: s.Data,
	}
	err = survey.AskOne(inputDataPath, &s.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *radarr) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
