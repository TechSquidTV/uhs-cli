package common

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Common struct {
	TZ      string `yaml:"tz"`
	Network Network `yaml:"network"`
	Certs   Certs   `yaml:"certs"`
}

type Network struct {
	Gateway string `yaml:"gateway"`
}

type Certs struct {
	SSLCertificateKey string `yaml:"ssl_certificate_key"`
	SSLCertificate    string `yaml:"ssl_certificate"`
	SSLDHParam        string `yaml:"ssl_dhparam"`
}

func (c *Common) Default() services.ServiceInterface {
	r := &Common{
		TZ: "America/New_York",
		Network: Network{
			Gateway: "192.168.1.1",
		},
		Certs: Certs{
			SSLCertificateKey: "/etc/letsencrypt/live/${DOMAIN}/privkey.pem",
			SSLCertificate:    "/etc/letsencrypt/live/${DOMAIN}/fullchain.pem",
			SSLDHParam:        "/etc/letsencrypt/certs/dhparam.pem",
		},
	}
	return r
}

func (c *Common) Configure() services.ServiceInterface {
	inputTz := &survey.Input{
		Message: "Enter your timezone:",
		Default: c.TZ,
	}
	err := survey.AskOne(inputTz, &c.TZ)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputGateway := &survey.Input{
		Message: "Enter your network gateway:",
		Default: c.Network.Gateway,
	}
	err = survey.AskOne(inputGateway, &c.Network.Gateway)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputSSLCertificateKey := &survey.Input{
		Message: "Enter your SSL certificate key path:",
		Default: c.Certs.SSLCertificateKey,
	}
	err = survey.AskOne(inputSSLCertificateKey, &c.Certs.SSLCertificateKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputSSLCertificate := &survey.Input{
		Message: "Enter your SSL certificate path:",
		Default: c.Certs.SSLCertificate,
	}
	err = survey.AskOne(inputSSLCertificate, &c.Certs.SSLCertificate)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputSSLDHParam := &survey.Input{
		Message: "Enter your SSL DHParam path:",
		Default: c.Certs.SSLDHParam,
	}
	err = survey.AskOne(inputSSLDHParam, &c.Certs.SSLDHParam)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	return c
}