package homepage

type Settings map[interface{}]interface{}

func CreateSettings(settings map[interface{}]interface{}) Settings {
	return Settings{
		settings: settings,
	}
}