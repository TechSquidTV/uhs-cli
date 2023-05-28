package radarr

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Radarr struct {
	services.Service `yaml:",inline"`
	RadarrOptions    `yaml:",inline"`
}

type RadarrOptions struct {
	Config    string `yaml:"config"`
	Downloads string `yaml:"downloads"`
	Movies     string `yaml:"movies"`
}

// Return default values for service
func (s *Radarr) Default() services.ServiceInterface {
	http := 7878

	p := &Radarr{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "lscr.io/linuxserver/radarr",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		RadarrOptions: RadarrOptions{
			Config:    "/opt/radarr/config",
			Downloads: "~/downloads",
			Movies:     "/opt/radarr/movies",
		},
	}
	return p
}

func (s *Radarr) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your radarr config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDownloadsPath := &survey.Input{
		Message: "Enter the path to your radarr downloads folder:",
		Default: s.Downloads,
	}
	err = survey.AskOne(inputDownloadsPath, &s.Downloads)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputMoviePath := &survey.Input{
		Message: "Enter the path to your radarr Movies folder:",
		Default: s.Movies,
	}
	err = survey.AskOne(inputMoviePath, &s.Movies)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
