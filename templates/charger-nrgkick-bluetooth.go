package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "nrgkick-bluetooth",
		Name:        "NRGKick Bluetooth",
		Sample:      `macaddress: 00:99:22 # MAC address
pin: # pin`,
	}

	registry.Registry.Add(template)
}
