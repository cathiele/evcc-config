package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "default",
		Name:   "Generisch (MQTT)",
		Sample: `status: # charger status A..F
  type: mqtt
  topic: some/topic1
enabled: # charger enabled state (true/false or 0/1)
  type: mqtt
  topic: some/topic2
enable: # set charger enabled state
  type: script
  cmd: /bin/sh -c "echo ${enable}"
maxcurrent: # set charger max current
  type: script
  cmd: /bin/sh -c "echo ${maxcurrent}"`,
	}

	registry.Add(template)
}
