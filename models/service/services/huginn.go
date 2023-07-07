package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type huginn struct {
	service.Service `yaml:",inline"`
	huginnOptions   `yaml:",inline"`
}

type huginnOptions struct {
	Data             string `yaml:"data"`
	Invitation_code  string `yaml:"invitation_code"`
	App_secret_token string `yaml:"app_secret_token"`
}

// Return default values for service
func NewHuginn() manager.Configurer {
	http := 3000

	return &huginn{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "ghcr.io/huginn/huginn",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		huginnOptions: huginnOptions{
			Data:             "/opt/huginn/data",
			Invitation_code:  "invite-me",
			App_secret_token: "<super secret token>",
		},
	}
}

func (s *huginn) Configure() {
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
}

func (s *huginn) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
