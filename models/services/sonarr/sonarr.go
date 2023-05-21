package sonarr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Sonarr struct {
	services.Service `yaml:",inline"`
	SonarrOptions    `yaml:",inline"`
}

type SonarrOptions struct {
	Config string `yaml:"config"`
	Downloads string `yaml:"downloads"`
	Tv string `yaml:"tv"`
}

// Return default values for service
func (s *Sonarr) Default() services.ServiceInterface {
	http := 8989

	p := &Sonarr{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "linuxserver/sonarr",
				Tag:        "latest",
				PullPolicy: "IfNotPresent",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		SonarrOptions: SonarrOptions{
			Config:  "/opt/sonarr/config",
			Downloads: "~/downloads",
			Tv: "/opt/sonarr/tv",
		},
	}
	return p
}

func (s *Sonarr) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Sonarr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDownloadsPath := &survey.Input{
		Message: "Enter the path to your Sonarr downloads folder:",
		Default: s.Downloads,
	}
	err = survey.AskOne(inputDownloadsPath, &s.Downloads)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputTvPath := &survey.Input{
		Message: "Enter the path to your Sonarr TV folder:",
		Default: s.Tv,
	}
	err = survey.AskOne(inputTvPath, &s.Tv)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}