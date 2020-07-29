package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "wallbe",
		Name:   "Wallbe (pre 2019 EV-CC-AC1 controller)",
		Sample: `uri: 192.168.0.8:502 # TCP ModBus address
legacy: true # enable for older Wallbes with Phoenix EV-CC-AC1-M3-CBC-RCM controller`,
	}

	registry.Add(template)
}
