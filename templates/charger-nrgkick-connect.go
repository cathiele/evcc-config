package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "nrgkick-connect",
		Name:        "NRGKick Connect",
		Sample:      `ip: 192.168.1.4 # IP
macaddress: 00:99:22 # MAC address
password: # password`,
	}

	registry.Registry.Add(template)
}
