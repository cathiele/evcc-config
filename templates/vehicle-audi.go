package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "audi",
		Name:   "Audi (eTron etc)",
		Sample: `title: eTron # display name for UI
capacity: 14 # kWh
user: # user
password: # password
vin: WAUZZZ... # optional
cache: 5m # cache API response`,
	}

	registry.Add(template)
}
