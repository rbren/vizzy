{{- define "plan" }}
{{- if .plan }}
Here's the plan for the implementation:
---
{{.plan}}
---
{{- else }}
Please plan the implementation step-by-step. Include sample code wherever possible. Specifically:
* Describe how to parse the data string described below, e.g. using `d3.csvParse`, `JSON.parse`, an XML parser, or using a custom parsing function if it's in a different format
* Describe how to transform the data into a shape and format that will be easiest to work with when creating the visualization, and provide sample code
* Describe the structure of the transformed data
* Describe how to sanitize the data, to ensure that any null or missing values are removed before being passed to D3
* Describe how to filter the data, to ensure that if there's a lot of data, only a reasonable number of bars/lines/etc (10-50) are presented to the user, and provide sample code

In all your sample code, be sure to respect the libraries that are available,
as specified in the Technical Details section.
{{- end }}
{{- end }}
