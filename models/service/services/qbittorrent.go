package services

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type qbittorrent struct {
	service.Service    `yaml:",inline"`
	qbittorrentOptions `yaml:",inline"`
}

type qbittorrentOptions struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
}

// Return default values for service
func NewQbittorrent() manager.Configurer {
	http := 8080
	p2p := 6881
	p2pudp := 6881

	// Create a new Qbittorrent instance
	return &qbittorrent{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "lscr.io/linuxserver/qbittorrent",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http:   &http,
				P2p:    &p2p,
				P2pudp: &p2pudp,
			},
		},
		qbittorrentOptions: qbittorrentOptions{
			Config: "/opt/qbittorrent/config",
			Data:   "/data/torrents",
		},
	}
}

func (s *qbittorrent) Configure() {
	s.Enabled = true
	inputConfigPath := &survey.Input{
		Message: "Enter the path to your qbittorrent config folder:",
		Default: s.Config,
	}
	err := survey.AskOne(inputConfigPath, &s.Config)
	if err != nil {
		fmt.Println(err.Error())
	}

	inputDataPath := &survey.Input{
		Message: "Enter the path to your torrent data folder:",
		Default: s.Data,
	}
	err = survey.AskOne(inputDataPath, &s.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *qbittorrent) Fill(data []byte) error { return yaml.Unmarshal(data, s) }
