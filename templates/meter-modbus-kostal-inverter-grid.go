package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Inverter (Grid Meter)",
		Sample: `model: kostal
uri: 192.168.178.52:1502 
id: 71 # Configured Modbus Device ID 
register: # manual register configuration
  address: 252 # (see https://www.kostal-solar-electric.com/de-de/download/-/media/document-library-folder---kse/2018/08/30/08/53/ba_kostal_interface_modbus-tcp_sunspec.pdf)
  type: holding
  decode: float32s #swapped float encoding`,
	}

	registry.Add(template)
}
