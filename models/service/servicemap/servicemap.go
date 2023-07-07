package servicemap

import (
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service/services"
)

var Registered = map[string]func() manager.Configurer{
	"qbittorrent":       services.NewQbittorrent,
	"plex":              services.NewPlex,
	"sonarr":            services.NewSonarr,
	"radarr":            services.NewRadarr,
	"sabnzbd":           services.NewSabnzbd,
	"cloudflared":       services.NewCloudflared,
	"overseerr":         services.NewOverseerr,
	"autobrr":           services.NewAutobrr,
	"prowlarr":          services.NewProwlarr,
	"kavita":            services.NewKavita,
	"gotify":            services.NewGotify,
	"tautulli":          services.NewTautulli,
	"playwright":        services.NewPlaywright,
	"thelounge":         services.NewThelounge,
	"apprise":           services.NewApprise,
	"changedetectionio": services.NewChangedetectionio,
	"huginn":            services.NewHuginn,
	"nginx":             services.NewNginx,
	"homepage":          services.NewHomepage,
	"jellyfin":          services.NewJellyfin,
}

// Keys is a utility function which returns the (string) keys of a map.
func Keys(m map[string]func() manager.Configurer) []string {
	var r []string
	for k := range m {
		r = append(r, k)
	}
	return r
}
