package kavita

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Kavita struct {
	services.Service `yaml:",inline"`
	KavitaOptions    `yaml:",inline"`
}

type KavitaOptions struct {
	Config  string `yaml:"config"`
	Library string `yaml:"library"`
}

// Return default values for service
func (s *Kavita) Default() services.ServiceInterface {
	http := 5000

	p := &Kavita{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "kizaing/kavita",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		KavitaOptions: KavitaOptions{
			Config:  "/opt/kavita/config",
			Library: "/data/library/books",
		},
	}
	return p
}

func (s *Kavita) Configure() services.ServiceInterface {
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
	return s
}
