package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "meter",
		Type:        "default",
		Name:        "E3DC Batterie",
		Sample:      `power:
  type: modbus
  uri: e3dc.fritz.box:502
  id: 1
  register: # manual register configuration
    address: 40070
    type: holding
    decode: int32
  scale: -1 # reverse direction`,
	}

	registry.Registry.Add(template)
}
