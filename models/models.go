package models

import (
	"github.com/techsquidtv/uhs-cli/models/services"
	"github.com/techsquidtv/uhs-cli/models/services/apprise"
	"github.com/techsquidtv/uhs-cli/models/services/autobrr"
	"github.com/techsquidtv/uhs-cli/models/services/cloudflared"
	"github.com/techsquidtv/uhs-cli/models/services/gotify"
	"github.com/techsquidtv/uhs-cli/models/services/kavita"
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
	overseerr := &overseerr.Overseerr{}
	autobrr := &autobrr.Autobrr{}
	prowlarr := &prowlarr.Prowlarr{}
	kavita := &kavita.Kavita{}
	gotify := &gotify.Gotify{}
	tautulli := &tautulli.Tautulli{}
	playwright := &playwright.Playwright{}
	thelounge := &thelounge.Thelounge{}
	apprise := &apprise.Apprise{}

	config["qbittorrent"] = qbt.Default()
	config["plex"] = plex.Default()
	config["sonarr"] = sonarr.Default()
	config["radarr"] = radarr.Default()
	config["sabnzbd"] = sabnzbd.Default()
	config["cloudflared"] = cloudflared.Default()
	config["overseerr"] = overseerr.Default()
	config["autobrr"] = autobrr.Default()
	config["prowlarr"] = prowlarr.Default()
	config["kavita"] = kavita.Default()
	config["gotify"] = gotify.Default()
	config["tautulli"] = tautulli.Default()
	config["playwright"] = playwright.Default()
	config["thelounge"] = thelounge.Default()
	config["apprise"] = apprise.Default()
	
	return config
}
