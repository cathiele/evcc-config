package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "tesla",
		Name:   "Tesla Powerwall (Battery meter)",
		Sample: `uri: http://192.168.1.4/api/meters/aggregates
usage: battery # grid meter: `+"`"+`site`+"`"+`, pv: `+"`"+`solar`+"`"+`, battery: `+"`"+`battery`+"`"+``,
	}

	registry.Add(template)
}
