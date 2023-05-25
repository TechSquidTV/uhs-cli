package qbittorrent

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/services"
)

type Qbittorrent struct {
	services.Service   `yaml:",inline"`
	QbittorrentOptions `yaml:",inline"`
}

type QbittorrentOptions struct {
	Config    string `yaml:"config"`
	Downloads string `yaml:"downloads"`
}

// Return default values for service
func (s *Qbittorrent) Default() services.ServiceInterface {
	http := 8080
	p2p := 6881
	p2pudp := 6881

	// Create a new Qbittorrent instance
	q := &Qbittorrent{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "lscr.io/linuxserver/qbittorrent",
				Tag:        "latest",
				PullPolicy: "IfNotPresent",
			},
			Ports: services.Ports{
				Http:   &http,
				P2p:    &p2p,
				P2pudp: &p2pudp,
			},
		},
		QbittorrentOptions: QbittorrentOptions{
			Config:    "/opt/qbittorrent/config",
			Downloads: "~/downloads",
		},
	}
	return q
}

func (s *Qbittorrent) Configure() services.ServiceInterface {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your qbittorrent config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDownloadsPath := &survey.Input{
		Message: "Enter the path to your qbittorrent downloads folder:",
		Default: s.Downloads,
	}
	err = survey.AskOne(inputDownloadsPath, &s.Downloads)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}
