package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:       "vehicle",
		Type:        "default",
		Name:        "Generisch",
		Sample:      `title: Mein Auto # name
capacity: 50 # kWh
charge:
  type: ...
  # ...
cache: 5m # cache duration`,
	}

	registry.Registry.Add(template)
}
