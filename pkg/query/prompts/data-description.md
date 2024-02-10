{{- define "data_description" }}
### Filetype and Summary
The data is in this format: {{.type}}

The title of the dataset is: {{.title}}

{{.description}}

### Structure
{{.data_format}}

### Fields
{{- if .fieldsMetadata }}
There are {{ index .fieldsMetadata "$dataPoints" }} data points in the set.
{{- end }}

Each data point has these fields:

{{- if .fieldsMetadata }}
{{-   range $key, $val := .fieldsMetadata }}
{{-     if ne $key "$dataPoints" }}
* `{{ $key }}` | {{ range $subkey, $subval := $val }} {{ $subkey }}:{{ formatValue $subval }}; {{ end }}
{{-     end }}
{{-   end }}
{{- else }}
{{-   range $key := .fields }}
* `{{ $key }}`
{{-   end }}
{{- end }}

Be sure to respect the capitalization and spaces in the above fields.

### Sample Data
Here is a small sample of the data:

```
{{ .sample }}
```
{{- end }}
