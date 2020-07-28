package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "phoenix-emcp",
		Name:   "Phoenix EMCP Controller (Ethernet/Modbus TCP)",
		Sample: `uri: 192.168.0.8:502 # ModBus address
id: 1`,
	}

	registry.Add(template)
}
