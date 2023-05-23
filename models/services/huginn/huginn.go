package huginn

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Huginn struct {
	services.Service `yaml:",inline"`
	HuginnOptions    `yaml:",inline"`
}

type HuginnOptions struct {
	Data             string `yaml:"data"`
	Invitation_code  string `yaml:"invitation_code"`
	App_secret_token string `yaml:"app_secret_token"`
}

// Return default values for service
func (s *Huginn) Default() services.ServiceInterface {
	http := 3000

	p := &Huginn{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "ghcr.io/huginn/huginn",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		HuginnOptions: HuginnOptions{
			Data:             "/opt/huginn/data",
			Invitation_code:  "invite-me",
			App_secret_token: "<super secret token>",
		},
	}
	return p
}

func (s *Huginn) Configure() services.ServiceInterface {
	s.Enabled = true
	inputDataPath := &survey.Input{
		Message: "Enter the path to your Huginn data folder:",
		Default: s.Data,
	}
	err := survey.AskOne(inputDataPath, &s.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	inputInvitationCode := &survey.Input{
		Message: "Enter the invitation code for your Huginn instance:",
		Default: s.Invitation_code,
	}
	err = survey.AskOne(inputInvitationCode, &s.Invitation_code)
	if err != nil {
		fmt.Println(err.Error())
	}
	inputAppSecretToken := &survey.Input{
		Message: "Enter the app secret token for your Huginn instance:",
		Default: s.App_secret_token,
	}
	err = survey.AskOne(inputAppSecretToken, &s.App_secret_token)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
