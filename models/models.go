package models

import (
	"github.com/techsquidtv/uhs-cli/models/services"
	"github.com/techsquidtv/uhs-cli/models/services/plex"
	"github.com/techsquidtv/uhs-cli/models/services/qbittorrent"
	"github.com/techsquidtv/uhs-cli/models/services/radarr"
	"github.com/techsquidtv/uhs-cli/models/services/sonarr"
)

type UHSConfig struct {
	Services services.ServicesConfig `yaml:"services"`
}

func DefaultServiceConfig() services.ServicesConfig {
	config := make(services.ServicesConfig)

	qbt := &qbittorrent.Qbittorrent{}
	plex := &plex.Plex{}
	sonarr := &sonarr.Sonarr{}
	radarr := &radarr.Radarr{}

	config["qbittorrent"] = qbt.Default()
	config["plex"] = plex.Default()
	config["sonarr"] = sonarr.Default()
	config["radarr"] = radarr.Default()

	return config
}
