# Data Visualization Task

## Current Code
We currently have this D3 code to generate a visualization of the data:

```javascript
{{ .current.code }}
```

The visualization currently has the title:
```
{{ .current.title }}
```
Keep this title the same, unless the code is changing in a way that
makes this title inaccurate.

## Dataset Description
{{template "data_description" .data}}

## Task
We need to modify the code above to satisfy this user prompt:
---
{{.prompt}}
---
The prompt might mention an issue with the current visualization, or ask for an enhancement.
You may need to rewrite the code significantly or refactor it. It's possible the current
code is incorrect, or needs a major change to address the user's prompt.

Don't make any changes to the visualization that were not explicitly requested by the user.

## Technical Implementation
{{template "technical_details"}}

{{ template "javascript_response" .current }}

## Style Guide
{{ template "style_guide" }}
