package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type nginx struct {
	service.Service `yaml:",inline"`
	nginxOptions    `yaml:",inline"`
}

type nginxOptions struct {
	PublicPath string `yaml:"public_path"`
	Domain     string `yaml:"domain"`
}

// Return default values for service
func NewNginx() manager.Configurer {
	http := 80
	https := 443

	return &nginx{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "nginx",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http:  &http,
				Https: &https,
			},
		},
		nginxOptions: nginxOptions{
			PublicPath: "/opt/nginx/public",
			Domain:     "example.com",
		},
	}
}

func (s *nginx) Configure() {
	s.Enabled = true
	inputPublicPath := &survey.Input{
		Message: "Enter the path to your Nginx public folder:",
		Default: s.PublicPath,
	}
	err := survey.AskOne(inputPublicPath, &s.PublicPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDomain := &survey.Input{
		Message: "Enter the domain you want to use:",
		Default: s.Domain,
	}

	err = survey.AskOne(inputDomain, &s.Domain)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *nginx) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
