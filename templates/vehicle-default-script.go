package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "vehicle",
		Type:        "default",
		Name:        "Generisch (Script)",
		Sample:      `title: Mein Auto # name
capacity: 50 # kWh
charge:
  type: script # use script
  cmd: /bin/sh -c "echo 50" # actual command
  timeout: 3s # kill script after 3 seconds
cache: 5m # cache duration`,
	}

	registry.Registry.Add(template)
}
