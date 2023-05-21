package models

import (
	"github.com/techsquidtv/uhs-cli/models/services"
	"github.com/techsquidtv/uhs-cli/models/services/plex"
	"github.com/techsquidtv/uhs-cli/models/services/qbittorrent"
)

type UHSConfig struct {
	Services services.ServicesConfig `yaml:"services"`
}

func DefaultServiceConfig() services.ServicesConfig {
	config := make(services.ServicesConfig)

	// Qbittorrent
	qbt := &qbittorrent.Qbittorrent{}
	plex := &plex.Plex{}
	config["qbittorrent"] = qbt.Default()
	config["plex"] = plex.Default()

	return config
}
