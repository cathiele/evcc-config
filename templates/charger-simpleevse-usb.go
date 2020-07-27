package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "simpleevse",
		Name:        "Simple EVSE (USB)",
		Sample:      `device: /dev/usb1 # RS485 ModBus device`,
	}

	registry.Registry.Add(template)
}
