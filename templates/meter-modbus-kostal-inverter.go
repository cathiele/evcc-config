package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Inverter (PV Meter)",
		Sample: `model: kostal
uri: 192.168.0.1:1502
id: 71
power: Power`,
	}

	registry.Add(template)
}
