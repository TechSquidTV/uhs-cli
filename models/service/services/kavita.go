package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type kavita struct {
	service.Service `yaml:",inline"`
	kavitaOptions   `yaml:",inline"`
}

type kavitaOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func NewKavita() manager.Configurer {
	http := 5000

	return &kavita{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "kizaing/kavita",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		kavitaOptions: kavitaOptions{
			Config:  "/opt/kavita/config",
			Library: "/data/library/books",
		},
	}
}

func (s *kavita) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your Kavita config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	inputLibraryPath := &survey.Input{
		Message: "Enter the path to your Kavita library folder:",
		Default: s.Library,
	}
	err = survey.AskOne(inputLibraryPath, &s.Library)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *kavita) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
