package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "nissan",
		Name:   "Nissan (Leaf)",
		Sample: `title: Leaf # display name for UI
capacity: 60 # kWh
user: # user
password: # password
region: NE # carwings region, leave empty for Europe
cache: 5m # cache API response`,
	}

	registry.Add(template)
}
