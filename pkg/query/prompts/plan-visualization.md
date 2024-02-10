# Data Visualization Task

## Task
We need to write JavaScript that uses D3 to create a visualization that satisfies this user prompt:

---
{{.prompt}}
---

{{template "plan"}}

Elaborate as much as you can on this plan. But don't write the final implementation--just
write step-by-step instructions that an implementer can follow.

## Dataset Description
{{template "data_description" .data}}

## Technical Details
{{template "technical_details"}}

You can assume that we've already installed D3, and that an SVG element has been created.
The implementation only needs to create a function
`drawVisualization(svg, data)`, which accepts an svg element and a string containing the data.
