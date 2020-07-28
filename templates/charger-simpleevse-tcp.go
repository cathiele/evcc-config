package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "simpleevse",
		Name:   "Simple EVSE (Ethernet/Modbus TCP)",
		Sample: `uri: 192.168.0.8:502 # TCP ModBus address`,
	}

	registry.Add(template)
}
