package sonarr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Sonarr struct {
	service.Service `yaml:",inline"`
	SonarrOptions   `yaml:",inline"`
}

type SonarrOptions struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
}

// Return default values for service
func (s *Sonarr) Default() service.ServiceInterface {
	http := 8989

	p := &Sonarr{
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
		SonarrOptions: SonarrOptions{
			Config: "/opt/sonarr/config",
			Data:   "/data",
		},
	}
	return p
}

func (s *Sonarr) Configure() service.ServiceInterface {
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

	return s
}
