package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "vw",
		Name:   "VW (ID.3, ID.4, etc)",
		Sample: `title: ID.3 # display name for UI
capacity: 10 # kWh
user: # user
password: # password
vin: WVWZZZ...
cache: 5m # cache API response`,
	}

	registry.Add(template)
}
