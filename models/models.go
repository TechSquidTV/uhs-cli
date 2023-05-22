package models

import (
	"github.com/techsquidtv/uhs-cli/models/services"
	"github.com/techsquidtv/uhs-cli/models/services/cloudflared"
	"github.com/techsquidtv/uhs-cli/models/services/plex"
	"github.com/techsquidtv/uhs-cli/models/services/qbittorrent"
	"github.com/techsquidtv/uhs-cli/models/services/radarr"
	"github.com/techsquidtv/uhs-cli/models/services/sabnzbd"
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
	sabnzbd := &sabnzbd.Sabnzbd{}
	cloudflared := &cloudflared.Cloudflared{}

	config["qbittorrent"] = qbt.Default()
	config["plex"] = plex.Default()
	config["sonarr"] = sonarr.Default()
	config["radarr"] = radarr.Default()
	config["sabnzbd"] = sabnzbd.Default()
	config["cloudflared"] = cloudflared.Default()

	return config
}
