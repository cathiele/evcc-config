package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "default",
		Name:   "Generisch",
		Sample: `status: # charger status A..F
  type: ...
  # ...
enabled: # charger enabled state (true/false or 0/1)
  type: ...
  # ...
enable: # set charger enabled state
  type: ...
  # ...
maxcurrent: # set charger max current
  type: ...
  # ...`,
	}

	registry.Add(template)
}
