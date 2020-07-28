# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters
{{range filter "meter" .}}
- {{.Name}}{{end}}

## Chargers
{{range filter "charger" .}}
- {{.Name}}{{end}}

## Details

### Meters

{{range filter "meter" .}}
#### {{.Name}}

```yaml
- type: {{.Class}}
{{.Sample | indent 2}}
```
{{end}}

### Chargers

{{range filter "charger" .}}
#### {{.Name}}

```yaml
- type: {{.Class}}
{{.Sample | indent 2}}
```
{{end}}
