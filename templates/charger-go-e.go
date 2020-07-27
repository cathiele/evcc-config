package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "go-e",
		Name:        "go-eCharger (Lokal)",
		Sample:      `uri: http://192.168.1.4 # either go-e local address`,
	}

	registry.Registry.Add(template)
}
