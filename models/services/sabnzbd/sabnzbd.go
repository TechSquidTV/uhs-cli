package sabnzbd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Sabnzbd struct {
	services.Service `yaml:",inline"`
	SabnzbdOptions   `yaml:",inline"`
}

type SabnzbdOptions struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
}

// Return default values for service
func (s *Sabnzbd) Default() services.ServiceInterface {
	http := 8080

	// Create a new Sabnzbd instance
	q := &Sabnzbd{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "lscr.io/linuxserver/sabnzbd",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		SabnzbdOptions: SabnzbdOptions{
			Config: "/opt/sabnzbd/config",
			Data:   "/data/usenet",
		},
	}
	return q
}

func (s *Sabnzbd) Configure() services.ServiceInterface {
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

	return s
}
