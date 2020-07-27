package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "meter",
		Type:        "default",
		Name:        "Generisch (MQTT)",
		Sample:      `power: # power reading
  type: mqtt # use mqtt
  topic: mbmd/sdm1-1/Power # mqtt topic
  timeout: 10s # don't use older values`,
	}

	registry.Registry.Add(template)
}
