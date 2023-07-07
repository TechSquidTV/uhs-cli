package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type sonarr struct {
	service.Service `yaml:",inline"`
	sonarrOptions   `yaml:",inline"`
}

type sonarrOptions struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
}

// Return default values for service
func NewSonarr() manager.Configurer {
	http := 8989

	return &sonarr{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/sonarr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		sonarrOptions: sonarrOptions{
			Config: "/opt/sonarr/config",
			Data:   "/data",
		},
	}
}

func (s *sonarr) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Sonarr config folder:",
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

func (s *sonarr) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
