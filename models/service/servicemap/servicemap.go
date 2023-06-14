package servicemap

import (
	"github.com/techsquidtv/uhs-cli/models/service"
	"github.com/techsquidtv/uhs-cli/models/service/services/apprise"
	"github.com/techsquidtv/uhs-cli/models/service/services/autobrr"
	"github.com/techsquidtv/uhs-cli/models/service/services/changedetectionio"
	"github.com/techsquidtv/uhs-cli/models/service/services/cloudflared"
	"github.com/techsquidtv/uhs-cli/models/service/services/gotify"
	"github.com/techsquidtv/uhs-cli/models/service/services/homepage"
	"github.com/techsquidtv/uhs-cli/models/service/services/huginn"
	"github.com/techsquidtv/uhs-cli/models/service/services/jellyfin"
	"github.com/techsquidtv/uhs-cli/models/service/services/kavita"
	"github.com/techsquidtv/uhs-cli/models/service/services/nginx"
	"github.com/techsquidtv/uhs-cli/models/service/services/overseerr"
	"github.com/techsquidtv/uhs-cli/models/service/services/playwright"
	"github.com/techsquidtv/uhs-cli/models/service/services/plex"
	"github.com/techsquidtv/uhs-cli/models/service/services/prowlarr"
	"github.com/techsquidtv/uhs-cli/models/service/services/qbittorrent"
	"github.com/techsquidtv/uhs-cli/models/service/services/radarr"
	"github.com/techsquidtv/uhs-cli/models/service/services/sabnzbd"
	"github.com/techsquidtv/uhs-cli/models/service/services/sonarr"
	"github.com/techsquidtv/uhs-cli/models/service/services/tautulli"
	"github.com/techsquidtv/uhs-cli/models/service/services/thelounge"
)

var ServiceMap = map[string]service.ServiceInterface{
    "qbittorrent":      &qbittorrent.Qbittorrent{},
    "plex":             &plex.Plex{},
    "sonarr":           &sonarr.Sonarr{},
    "radarr":           &radarr.Radarr{},
    "sabnzbd":          &sabnzbd.Sabnzbd{},
    "cloudflared":      &cloudflared.Cloudflared{},
    "overseerr":        &overseerr.Overseerr{},
    "autobrr":          &autobrr.Autobrr{},
    "prowlarr":         &prowlarr.Prowlarr{},
    "kavita":           &kavita.Kavita{},
    "gotify":           &gotify.Gotify{},
    "tautulli":         &tautulli.Tautulli{},
    "playwright":       &playwright.Playwright{},
    "thelounge":        &thelounge.Thelounge{},
    "apprise":          &apprise.Apprise{},
    "changedetectionio": &changedetectionio.Changedetectionio{},
    "huginn":           &huginn.Huginn{},
    "nginx":            &nginx.Nginx{},
    "homepage":         &homepage.Homepage{},
    "jellyfin":         &jellyfin.Jellyfin{},
}