package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "evsewifi",
		Name:        "EVSE Wifi",
		Sample:      `uri: http://192.168.1.4 # SimpleEVSE-Wifi address`,
	}

	registry.Registry.Add(template)
}
