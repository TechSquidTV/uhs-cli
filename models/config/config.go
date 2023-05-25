package config

import (
	"github.com/techsquidtv/uhs-cli/models/services"
	"github.com/techsquidtv/uhs-cli/models/services/apprise"
	"github.com/techsquidtv/uhs-cli/models/services/autobrr"
	"github.com/techsquidtv/uhs-cli/models/services/changedetectionio"
	"github.com/techsquidtv/uhs-cli/models/services/cloudflared"
	"github.com/techsquidtv/uhs-cli/models/services/gotify"
	"github.com/techsquidtv/uhs-cli/models/services/homepage"
	"github.com/techsquidtv/uhs-cli/models/services/huginn"
	"github.com/techsquidtv/uhs-cli/models/services/kavita"
	"github.com/techsquidtv/uhs-cli/models/services/nginx"
	"github.com/techsquidtv/uhs-cli/models/services/overseerr"
	"github.com/techsquidtv/uhs-cli/models/services/playwright"
	"github.com/techsquidtv/uhs-cli/models/services/plex"
	"github.com/techsquidtv/uhs-cli/models/services/prowlarr"
	"github.com/techsquidtv/uhs-cli/models/services/qbittorrent"
	"github.com/techsquidtv/uhs-cli/models/services/radarr"
	"github.com/techsquidtv/uhs-cli/models/services/sabnzbd"
	"github.com/techsquidtv/uhs-cli/models/services/sonarr"
	"github.com/techsquidtv/uhs-cli/models/services/tautulli"
	"github.com/techsquidtv/uhs-cli/models/services/thelounge"
)

type Config struct {
	Common	 services.ServiceInterface          `yaml:"common,omitempty"`
	Services services.ServicesConfig `yaml:"services"`
}

func DefaultServiceConfig() services.ServicesConfig {
	uhsConfig := make(services.ServicesConfig)

	qbt := &qbittorrent.Qbittorrent{}
	plex := &plex.Plex{}
	sonarr := &sonarr.Sonarr{}
	radarr := &radarr.Radarr{}
	sabnzbd := &sabnzbd.Sabnzbd{}
	cloudflared := &cloudflared.Cloudflared{}
	overseerr := &overseerr.Overseerr{}
	autobrr := &autobrr.Autobrr{}
	prowlarr := &prowlarr.Prowlarr{}
	kavita := &kavita.Kavita{}
	gotify := &gotify.Gotify{}
	tautulli := &tautulli.Tautulli{}
	playwright := &playwright.Playwright{}
	thelounge := &thelounge.Thelounge{}
	apprise := &apprise.Apprise{}
	changedetectionio := &changedetectionio.Changedetectionio{}
	huginn := &huginn.Huginn{}
	nginx := &nginx.Nginx{}
	homepage := &homepage.Homepage{}

	uhsConfig["qbittorrent"] = qbt.Default()
	uhsConfig["plex"] = plex.Default()
	uhsConfig["sonarr"] = sonarr.Default()
	uhsConfig["radarr"] = radarr.Default()
	uhsConfig["sabnzbd"] = sabnzbd.Default()
	uhsConfig["cloudflared"] = cloudflared.Default()
	uhsConfig["overseerr"] = overseerr.Default()
	uhsConfig["autobrr"] = autobrr.Default()
	uhsConfig["prowlarr"] = prowlarr.Default()
	uhsConfig["kavita"] = kavita.Default()
	uhsConfig["gotify"] = gotify.Default()
	uhsConfig["tautulli"] = tautulli.Default()
	uhsConfig["playwright"] = playwright.Default()
	uhsConfig["thelounge"] = thelounge.Default()
	uhsConfig["apprise"] = apprise.Default()
	uhsConfig["changedetectionio"] = changedetectionio.Default()
	uhsConfig["huginn"] = huginn.Default()
	uhsConfig["nginx"] = nginx.Default()
	uhsConfig["homepage"] = homepage.Default()

	return uhsConfig
}