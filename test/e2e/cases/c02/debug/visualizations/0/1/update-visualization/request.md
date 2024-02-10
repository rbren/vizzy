# Data Visualization Task

## Current Code
We currently have this D3 code to generate a visualization of the data:

```javascript

async function drawVisualization(svg, data) {
    // Parse the CSV data string
    let parsedData = d3.csvParse(data);

    // Transform data: Parse date strings to Date objects and filter out any rows with missing data.
    parsedData = parsedData.filter(row => row.Date && row.CO2 && row['adjusted CO2'])
                           .map(row => ({
                                date: d3.timeParse('%Y-%m-%d')(row.Date),
                                co2: +row.CO2,
                                adjustedCo2: +row['adjusted CO2']
                            }));

    if(parsedData.length === 0) throw new Error("No valid data points found.");
    
    // Prepare SVG dimensions and margins
    const margin = { top: 20, right: 30, bottom: 50, left: 60 },
          width = svg.attr('width') - margin.left - margin.right,
          height = svg.attr('height') - margin.top - margin.bottom;

    // Create scales
    const xScale = d3.scaleTime()
                     .domain(d3.extent(parsedData, d => d.date))
                     .range([0, width]);

    const yScale = d3.scaleLinear()
                     .domain([d3.min(parsedData, d => Math.min(d.co2, d.adjustedCo2)), 
                              d3.max(parsedData, d => Math.max(d.co2, d.adjustedCo2))])
                     .nice()
                     .range([height, 0]);

    // Create and append axes
    const xAxis = d3.axisBottom(xScale).ticks(width > 500 ? 16 : 8);
    const yAxis = d3.axisLeft(yScale);

    svg.append("g")
       .attr("transform", `translate(${margin.left},${height + margin.top})`)
       .call(xAxis)
       .style("color", "#fff");

    svg.append("g")
       .attr("transform", `translate(${margin.left},${margin.top})`)
       .call(yAxis)
       .style("color", "#fff");

    // Prepare the lines to be drawn
    const lineGeneratorCO2 = d3.line()
                                .x(d => xScale(d.date))
                                .y(d => yScale(d.co2));

    const lineGeneratorAdjustedCO2 = d3.line()
                                       .x(d => xScale(d.date))
                                       .y(d => yScale(d.adjustedCo2));

    // Append paths for CO2 and adjusted CO2 lines
    svg.append('path')
       .datum(parsedData)
       .attr('d', lineGeneratorCO2)
       .attr('fill', 'none')
       .attr('stroke', 'limegreen')
       .attr('stroke-width', 2)
       .attr('transform', `translate(${margin.left},${margin.top})`);

    svg.append('path')
       .datum(parsedData)
       .attr('d', lineGeneratorAdjustedCO2)
       .attr('fill', 'none')
       .attr('stroke', 'orange')
       .attr('stroke-width', 2)
       .attr('transform', `translate(${margin.left},${margin.top})`);

    // Add labels
    svg.append("text")
       .attr("transform", `translate(${width / 2 + margin.left},${height + margin.top + 40})`)
       .style("text-anchor", "middle")
       .style("fill", "#fff")
       .text("Time");

    svg.append("text")
       .attr("transform", "rotate(-90)")
       .attr("y", 0)
       .attr("x", -(height / 2) - margin.top)
       .attr("dy", "1em")
       .style("text-anchor", "middle")
       .style("fill", "#fff")
       .text("CO2 Concentrations (ppm)");
}

```

The visualization currently has the title:
```
Historical CO2 Concentrations Over Time
```
Keep this title the same, unless the code is changing in a way that
makes this title inaccurate.

## Dataset Description

### Filetype and Summary
The data is in this format: CSV

The title of the dataset is: CO2 Historical Data

Contains monthly measurements of atmospheric CO2 concentrations, both unadjusted and adjusted values.

### Structure
The data is in a CSV format, with each row representing a monthly measurement of CO2. There are three columns in the data: `Date`, `CO2`, and `adjusted CO2`. The 'Date' field is formatted as YYYY-MM-DD, though only year and month are significant since the day is always set to the first of the month. The `CO2` and `adjusted CO2` fields contain floating point numbers representing the measured and adjusted values of atmospheric CO2 concentrations in parts per million (ppm). Before analyzing this data, it might be necessary to parse the date fields into a date/time representation suitable for time series analysis. Furthermore, handling missing data points (for example, there appears to be no entries for certain months) may be required for comprehensive analysis.

### Fields

Each data point has these fields:
* `Date`
* `CO2`
* `adjusted CO2`

Be sure to respect the capitalization and spaces in the above fields.

### Sample Data
Here is a small sample of the data:

```
Date,CO2,adjusted CO2
1958-03-01,315.70,314.44
1958-04-01,317.46,315.16
1958-05-01,317.51,314.71
1958-07-01,315.86,315.19
1958-08-01,314.93,316.19
1958-09-01,313.21,316.08
1958-11-01,313.33,315.20
1958-12-01,314.67,315.43
1959-01-01,315.58,315.54
1959-02-01,316.49,315.86
1959-03-01,316.65,315.38
1959-04-01,317.72,315.42
1959-05-01,318.29,315.49
1959-06-01,318.15,316.03
1959-07-01,316.54,315.86
1959-08-01,314.80,316.06
1959-09-01,313.84,316.73
1959-10-01,313.33,316.33
1959-11-01,314.81,316.68
1959-12-01,315.58,316.35
1960-01-01,316.43,316.39
1960-02-01,316.98,316.35
1960-03-01,317.58,316.28
1960-04-01,319.03,316.70
1960-05-01,320.04,317.22
1960-06-01,319.58,317.48
1960-07-01,318.18,317.52
1960-08-01,315.90,317.20
1960-09-01,314.17,317.08
1960-10-01,313.83,316.83
1960-11-01,315.00,316.88
1960-12-01,316.19,316.96
1961-01-01,316.90,316.85
1961-02-01,317.70,317.07
1961-03-01,318.54,317.26
1961-04-01,319.48,317.16
1961-05-01,320.58,317.76
1961-06-01,319.77,317.63
1961-07-01,318.57,317.88
196
```

## Task
We need to modify the code above to satisfy this user prompt:
---
Set the bottom of the y axis to be 250
---
The prompt might mention an issue with the current visualization, or ask for an enhancement.
You may need to rewrite the code significantly or refactor it. It's possible the current
code is incorrect, or needs a major change to address the user's prompt.

Don't make any changes to the visualization that were not explicitly requested by the user.

## Technical Implementation

The following libraries are available:
* D3 v7, as `d3`
* topojson v3, as `topojson`
* lodash v4, as `_`
* d3-hexbin v0.2, as `d3.hexbin`

These libraries are already available in the global scope. Do not try to import them inside the code.
No other third-party libraries are available for use. Do not import any third-party libraries.

You may download third-party data, e.g. for map data. Here are some datasets that may be useful:
* TopoJSON with country and land geometry: https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json
* TopoJSON with US states: https://cdn.jsdelivr.net/npm/us-atlas@3/states-10m.json
* TopoJSON with US counties: https://cdn.jsdelivr.net/npm/us-atlas@3/counties-10m.json

Here are the changes in D3 that have taken place as of v7, which you should be aware of:

* You can no longer use `d3.event`. Instead, `event` is passed as the first argument to all listeners
* `d3.mouse`, `d3.touch`, `d3.touches`, `d3.clientPoint` are no longer available. Instead, use `d3.pointer`.
* `d3.voronoi` is deprecated, and replaced by `d3.Delaunay`
* `d3.nest` is no longer available. Instead, use `d3.group` and `d3.rollup` (from `d3-array`)
* `d3.map` is no longer available. Instead, use `Map`
* `d3.set` is no longer available. Instead, use `Set`
* `d3.keys` is no longer available. Instead, use `Object.keys`
* `d3.values` is no longer available. Instead, use `Object.values`
* `d3.entries` is no loonger available. Instead, use `Object.entries`
* `d3.histogram` is now called `d3.bin`
* `d3.scan` is now called `d3.leastIndex`


Your code must provide a function `drawVisualization(svg, data)`, which accepts an svg element
that has already been created using `d3.create('svg')`, and a string containing all of the data,
already downloaded. The data string must be parsed into structured data based on the description above.

The svg will already have a width and a height, which you can access using
`svg.attr('width')` and `svg.attr('height')`. You must take these variables into account
to ensure the visualization fits within the viewable area. All sizes and positions should
take these variables into account.

If the user asks for a tooltip or other HTML, you may add a `<div>` to the HTML body of the document.
The div must be positioned on top of the SVG, and must have the class `tooltip`.

The `drawVisualization(svg, data)` function must be an `async` function. If it calls any
function that returns a Promise, it must return that Promise.
Be sure to bubble up any errors to the caller, using `throw` or `reject`.

The data may not be perfectly sanitized. Be sure to remove any null or missing values,
or data that's the wrong type, before passing it to d3's `data()` function.
You should throw an error if there are no valid data points. Do not take any data point as
canonical--use `_.flatten()` when necessary to construct things like axes and scales.

Don't include any HTML, or any code outside of the `drawVisualization(svg, data)` function.

Don't include any comments in your code.

Your response will be run directly in a web application, so it must include
all the code required to run properly. Do not call the `drawVisualization()` function,
or include sample code on how to call it.

In your response, please use a large markdown header to give a title to
the visualization this code will generate, per the instructions in the Style Guide section.
The current title is "Historical CO2 Concentrations Over Time".
You should use this as the title unless it would be completely inaccurate, or
the user has explicitly requested a different title.
If the user prompt asks you to change the title, this is the title you should change. Don't
add a title in the JavaScript code.

Place three backticks at the start and end of your code. Here's an example of the format
for your response:

---
# Hello World
## Plan
I will use `console.log` to print the message `Hello world`, like this:
```javascript
console.log("hello world");
```
## Code
```javascript
function drawVisualization(svg, data) {
    console.log("hello world");
}
```
---

## Style Guide

The markdown title should be a concise description of the data being shown,
not a description of how it's being displayed.
The title shouldn't include the word "visualization", or words like "tooltip",
"pie chart", "error bar", etc.
Some examples of good titles might be "Stock Prices Over Time",
"Relative Horsepower of Cars by Country", or "Average Rank of Genre by Year".
Do not render a title inside the SVG.

Be sure to include any labels or legends that would help the user
understand the data being presented.

Color schemes must use `interpolateSpectral` unless otherwise requested by the user.
If you have N discreate colors to render, use
`d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, N))`.
If you're rendering a range of colors, use `d3.scaleLinear` with semantically appropriate colors.

The visualization will be presented against a dark background. Text, axes, and other information
should be rendered as #fff white. All fonts should be sans-serif.

Tooltips must use a white background, black text, and sans-serif font.
They should have 10px of padding and a 5px border radius.

Be sure to draw elements in the correct order. Text must be drawn last, so it appears on top.

Be sure to remove the default black fill from any paths.
