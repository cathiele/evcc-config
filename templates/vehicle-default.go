package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "default",
		Name:   "Generisch",
		Sample: `title: Mein Auto # display name for UI
capacity: 50 # kWh
charge:
  type: ...
  # ...
cache: 5m # cache duration`,
	}

	registry.Add(template)
}
