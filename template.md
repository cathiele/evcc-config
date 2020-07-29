# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters
{{range filter "meter" .}}
- [{{.Name}}](#{{href "meter" .Name}}){{end}}

## Chargers
{{range filter "charger" .}}
- [{{.Name}}](#{{href "charger" .Name}}){{end}}

## Vehicles
{{range filter "vehicle" .}}
- [{{.Name}}](#{{href "vehicle" .Name}}){{end}}

## Details

### Meters

{{range filter "meter" .}}
<a id="{{href "meter" .Name}}"></a>
#### {{.Name}}

```yaml
- type: {{.Type}}
{{.Sample | indent 2}}
```
{{end}}

### Chargers

{{range filter "charger" .}}
<a id="{{href "charger" .Name}}"></a>
#### {{.Name}}

```yaml
- type: {{.Type}}
{{.Sample | indent 2}}
```
{{end}}

### Vehicles

{{range filter "vehicle" .}}
<a id="{{href "vehicle" .Name}}"></a>
#### {{.Name}}

```yaml
- type: {{.Type}}
{{.Sample | indent 2}}
```
{{end}}
