package homepage

type BookmarkGroup map[string][]Bookmark

type Bookmark map[string][]BookmarkOptions

type BookmarkOptions struct {
	Icon *string `yaml:"icon,omitempty"`
	Abbr *string `yaml:"abbr,omitempty"`
	Href string  `yaml:"href"`
}

func CreateBookmark(name, icon, abbr, href string) Bookmark {
	return Bookmark{
		name: []BookmarkOptions{
			{
				Icon: &icon,
				Abbr: &abbr,
				Href: href,
			},
		},
	}
}
