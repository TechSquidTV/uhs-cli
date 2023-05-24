package homepage

import (
	"github.com/techsquidtv/uhs-cli/models/services"
)

func GetString(s string) *string {
	return &s
}

type Homepage struct {
	services.Service `yaml:",inline"`
	HomepageOptions  `yaml:",inline"`
}

type HomepageOptions struct {
	Bookmarks []BookmarkGroup `yaml:"bookmarks"`
  Services  []ServiceGroup  `yaml:"services"`
  Widgets   []Widget        `yaml:"widgets"`
  Settings  Settings        `yaml:"settings"`
}

// Return default values for service
func (s *Homepage) Default() services.ServiceInterface {
	http := 3000

	bookmarksConfig := []BookmarkGroup{
		{
			"Development": []Bookmark{
				CreateBookmark("github", "mdi-github", "gh", "https://github.com"),
			},
		},
    {
      "Media": []Bookmark{
				CreateBookmark("youtube", "mdi-youtube", "yt", "https://youtube.com"),
				CreateBookmark("plex", "mdi-plex", "plex", "https://app.plex.tv/desktop"),
			},
    },
	}
  servicesConfig := []ServiceGroup{
    {
      "Media": []Service{
        CreateService("plex", "mdi-plex", "https://app.plex.tv/desktop", "Plex"),
      },
    },
  }
  widgetConfig := []Widget{
    CreateWidget("search", map[string]interface{}{
        "provider": "[duckduckgo, google]",
        "focus": true,
        "target": "_blank",
      }),
  }
  settingsConfig := Settings{
    "title": "Homepage",
    "theme": "dark",
  }

	p := &Homepage{
		Service: services.Service{
			Enabled:      false,
			ReplicaCount: 1,
			Image: services.Image{
				Repository: "ghcr.io/benphelps/homepage",
				Tag:        "latest",
				PullPolicy: "Always",
			},
			Ports: services.Ports{
				Http: &http,
			},
		},
		HomepageOptions: HomepageOptions{
			Bookmarks: bookmarksConfig,
      Services:  servicesConfig,
      Widgets:   widgetConfig,
      Settings:  settingsConfig,
		},
	}
	return p
}

func (s *Homepage) Configure() services.ServiceInterface {
	s.Enabled = true
	return s
}
