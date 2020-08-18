package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Sonnenbatterie Eco (PV meter/ HTTP)",
		Sample: `power: # power reading
  type: http # use http plugin
  uri: http://192.168.1.75:8080/api/v1/status
  jq: .Production_W`,
	}

	registry.Add(template)
}
