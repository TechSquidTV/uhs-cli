package homepage

type Widget map[string]map[string]interface{}

func CreateWidget(name string, resources map[string]interface{}) Widget {
	return Widget{
		name: resources,
	}
}
