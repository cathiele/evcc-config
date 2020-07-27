# Configuration examples for EVCC

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
