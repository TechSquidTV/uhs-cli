package services

import (
	"github.com/techsquidtv/uhs-cli/models/config/manager"
	"github.com/techsquidtv/uhs-cli/models/service"
	"gopkg.in/yaml.v3"
)

type homepage struct {
	service.Service `yaml:",inline"`
	homepageOptions `yaml:",inline"`
}

type homepageOptions struct {
	Bookmarks []homepageBookmarkGroup `yaml:"bookmarks"`
	Services  []homepageServiceGroup  `yaml:"services"`
	Widgets   []homepageWidget        `yaml:"widgets"`
	Settings  map[any]any             `yaml:"settings"`
}

// Return default values for service
func NewHomepage() manager.Configurer {
	http := 3000

	bookmarksConfig := []homepageBookmarkGroup{
		{
			"Development": []homepageBookmark{
				createHomepageBookmark("github", "mdi-github", "gh", "https://github.com"),
			},
		},
		{
			"Media": []homepageBookmark{
				createHomepageBookmark("youtube", "mdi-youtube", "yt", "https://youtube.com"),
				createHomepageBookmark("plex", "mdi-plex", "plex", "https://app.plex.tv/desktop"),
			},
		},
	}
	servicesConfig := []homepageServiceGroup{
		{
			"Media": []homepageService{
				createHomepageService("plex", "mdi-plex", "https://app.plex.tv/desktop", "Plex"),
			},
		},
	}
	widgetConfig := []homepageWidget{
		createHomepageWidget("search", map[string]any{
			"provider": "[duckduckgo, google]",
			"focus":    true,
			"target":   "_blank",
		}),
	}
	settingsConfig := map[any]any{
		"title": "Homepage",
		"theme": "dark",
	}

	return &homepage{
		Service: service.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: service.Image{
				Repository: "ghcr.io/benphelps/homepage",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: service.Ports{
				Http: &http,
			},
		},
		homepageOptions: homepageOptions{
			Bookmarks: bookmarksConfig,
			Services:  servicesConfig,
			Widgets:   widgetConfig,
			Settings:  settingsConfig,
		},
	}
}

func (s *homepage) Configure() {
	s.Enabled = true
}

func (s *homepage) Fill(data []byte) error { return yaml.Unmarshal(data, s) }

//////////////
// Bookmark //
//////////////

type homepageBookmarkGroup map[string][]homepageBookmark
type homepageBookmark map[string][]homepageBookmarkOptions

type homepageBookmarkOptions struct {
	Icon *string `yaml:"icon,omitempty"`
	Abbr *string `yaml:"abbr,omitempty"`
	Href string  `yaml:"href"`
}

func createHomepageBookmark(name, icon, abbr, href string) homepageBookmark {
	return homepageBookmark{
		name: []homepageBookmarkOptions{
			{
				Icon: &icon,
				Abbr: &abbr,
				Href: href,
			},
		},
	}
}

/////////////
// Service //
/////////////

type homepageServiceGroup map[string][]homepageService
type homepageService map[string][]homepageServiceOptions

type homepageServiceOptions struct {
	Icon        *string `yaml:"icon,omitempty"`
	Href        string  `yaml:"href"`
	Description *string `yaml:"description,omitempty"`
}

func createHomepageService(name, icon, href, description string) homepageService {
	return homepageService{
		name: []homepageServiceOptions{
			{
				Icon:        &icon,
				Href:        href,
				Description: &description,
			},
		},
	}
}

////////////
// Widget //
////////////

type homepageWidget map[string]map[string]any

func createHomepageWidget(name string, resources map[string]any) homepageWidget {
	return homepageWidget{
		name: resources,
	}
}
