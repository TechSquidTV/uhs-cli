package cloudflared

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/service"
)

type Cloudflared struct {
	service.Service    `yaml:",inline"`
	CloudflaredOptions `yaml:",inline"`
}

type CloudflaredOptions struct {
	Tunnel string `yaml:"tunnel"`
	Domain string `yaml:"domain"`
	URL    string `yaml:"url"`
	Target string `yaml:"target"`
}

// Return default values for service
func (s *Cloudflared) Default() service.ServiceInterface {

	p := &Cloudflared{
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
		CloudflaredOptions: CloudflaredOptions{
			Tunnel: "example-tunnel",
			Domain: "example.com",
			URL:    "http://foo.example.com",
			Target: "localhost:8080",
		},
	}
	return p
}

func (s *Cloudflared) Configure() service.ServiceInterface {
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

	return s
}
