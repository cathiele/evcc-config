package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "charger",
		Type:        "keba",
		Name:        "KEBA Connect",
		Sample:      `uri: 192.168.1.4:7090 # KEBA address
rfid:
  tag: 765765348 # RFID tag, see `+"`"+`evcc charger`+"`"+` to show tag`,
	}

	registry.Registry.Add(template)
}
