package nginx

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Nginx struct {
	service.Service `yaml:",inline"`
	NginxOptions     `yaml:",inline"`
}

type NginxOptions struct {
	Public_path string `yaml:"public_path"`
	Domain      string `yaml:"domain"`
}

// Return default values for service
func (s *Nginx) Default() service.ServiceInterface {
	http := 80
	https := 443

	p := &Nginx{
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
		NginxOptions: NginxOptions{
			Public_path: "/opt/nginx/public",
			Domain:      "example.com",
		},
	}
	return p
}

func (s *Nginx) Configure() service.ServiceInterface {
	s.Enabled = true
	inputPublicPath := &survey.Input{
		Message: "Enter the path to your Nginx public folder:",
		Default: s.Public_path,
	}
	err := survey.AskOne(inputPublicPath, &s.Public_path)
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

	return s
}
