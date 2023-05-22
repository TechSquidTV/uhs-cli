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
	Config    string `yaml:"config"`
	Downloads string `yaml:"downloads"`
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
				Repository: "linuxserver/sabnzbd",
				Tag:        "latest",
				PullPolicy: "IfNotPresent",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		SabnzbdOptions: SabnzbdOptions{
			Config:    "/opt/sabnzbd/config",
			Downloads: "~/downloads",
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

	inputDownloadsPath := &survey.Input{
		Message: "Enter the path to your sabnzbd downloads folder:",
		Default: s.Downloads,
	}
	err = survey.AskOne(inputDownloadsPath, &s.Downloads)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
