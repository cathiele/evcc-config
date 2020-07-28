package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "phoenix-evcc",
		Name:   "Phoenix EVCC Controller (Modbus)",
		Sample: `device: /dev/ttyUSB0
baudrate: 9600
comset: 8N1
id: 1 # slave id`,
	}

	registry.Add(template)
}
