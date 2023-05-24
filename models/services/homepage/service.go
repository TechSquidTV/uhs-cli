package homepage

type ServiceGroup map[string][]Service

type Service map[string][]ServiceOptions

type ServiceOptions struct {
	Icon *string `yaml:"icon,omitempty"`
	Href string  `yaml:"href"`
	Description *string `yaml:"description,omitempty"`
}

func CreateService(name, icon, href, description string) Service {
	return Service{
		name: []ServiceOptions{
			{
				Icon: &icon,
				Href: href,
				Description: &description,
			},
		},
	}
}