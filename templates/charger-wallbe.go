package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "wallbe",
		Name:        "Wallbe (Eco, Pro)",
		Sample:      `uri: 192.168.0.8:502 # ModBus address`,
	}

	registry.Registry.Add(template)
}
