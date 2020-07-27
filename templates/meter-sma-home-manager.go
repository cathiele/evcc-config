package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "meter",
		Type:        "sma",
		Name:        "SMA Home Manager 2.0 oder SMA Energy Meter 30",
		Sample:      `serial: 1234567890 # Serial number of the device, if this is defined uri is not needed!
uri: 192.168.1.4 # IP address of the device, if this is defined serial is not needed!`,
	}

	registry.Registry.Add(template)
}
