package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA Sunny Island",
		Sample: `model: sunny-island
uri: 192.168.1.4:502
id: 126
power: Power # default values, optionally override
energy: Sum # default values, optionally override`,
	}

	registry.Add(template)
}
