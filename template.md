# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters
{{range $id, $el := filter "meter" .}}
- [{{.Name}}](#meter-{{$id}}){{end}}

## Chargers
{{range $id, $el := filter "charger" .}}
- [{{.Name}}](#charger-{{$id}}){{end}}

## Vehicles
{{range $id, $el := filter "vehicle" .}}
- [{{.Name}}](#vehicle-{{$id}}){{end}}

## Details

### Meters

{{range $id, $el := filter "meter" .}}
<a id="meter-{{$id}}"></a>
#### {{.Name}}

```yaml
- type: {{.Type}}
{{.Sample | indent 2}}
```
{{end}}

### Chargers

{{range $id, $el := filter "charger" .}}
<a id="charger-{{$id}}"></a>
#### {{.Name}}

```yaml
- type: {{.Type}}
{{.Sample | indent 2}}
```
{{end}}

### Vehicles

{{range $id, $el := filter "vehicle" .}}
<a id="vehicle-{{$id}}"></a>
#### {{.Name}}

```yaml
- type: {{.Type}}
{{.Sample | indent 2}}
```
{{end}}
