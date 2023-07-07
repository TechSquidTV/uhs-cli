package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type cloudflared struct {
	service.Service    `yaml:",inline"`
	cloudflaredOptions `yaml:",inline"`
}

type cloudflaredOptions struct {
	Tunnel string `yaml:"tunnel"`
	Domain string `yaml:"domain"`
	URL    string `yaml:"url"`
	Target string `yaml:"target"`
}

// Return default values for service
func NewCloudflared() manager.Configurer {

	return &cloudflared{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "cloudflare/cloudflared",
				Tag:        "latest",
				PullPolicy: "IfNotPresent", // For stability
			},
			Ports: service.Ports{},
		},
		cloudflaredOptions: cloudflaredOptions{
			Tunnel: "example-tunnel",
			Domain: "example.com",
			URL:    "http://foo.example.com",
			Target: "localhost:8080",
		},
	}
}

func (s *cloudflared) Configure() {
	s.Enabled = true
	inputTunnel := &survey.Input{
		Message: "Enter the name of your tunnel:",
		Default: s.Tunnel,
	}
	err := survey.AskOne(inputTunnel, &s.Tunnel)
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

	inputURL := &survey.Input{
		Message: "Enter the URL you want to use:",
		Default: s.URL,
	}
	err = survey.AskOne(inputURL, &s.URL)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputTarget := &survey.Input{
		Message: "Enter the target you want to use:",
		Default: s.Target,
	}
	err = survey.AskOne(inputTarget, &s.Target)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *cloudflared) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
