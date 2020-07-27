package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "meter",
		Type:        "default",
		Name:        "Generisch (Script)",
		Sample:      `power:
  type: script # use script
  cmd: /bin/sh -c "echo 0" # actual command
  timeout: 3s # kill script after 3 seconds`,
	}

	registry.Registry.Add(template)
}
