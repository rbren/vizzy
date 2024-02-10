# Data Visualization Task

## Current Code
We currently have this D3 code to generate a visualization of the data:

```javascript

async function drawVisualization(svg, data) {
    // Parse the CSV data
    const parsedData = d3.csvParse(data);

    // Transform the data
    const transformedData = parsedData.map(row => {
        return {
            symbol: row.symbol,
            date: new Date(row.date),
            price: +row.price
        };
    }).filter(row => row.symbol === "MSFT"); // Filtering to include only Microsoft for demonstration as per plan

    // Check for any missing or null values
    const cleanedData = transformedData.filter(row => row.date && !isNaN(row.price) && row.price !== null);

    if (cleanedData.length === 0) {
        throw new Error("No valid data points.");
    }
    
    // Set basic dimensions
    const margin = {top: 20, right: 25, bottom: 30, left: 60},
          width = +svg.attr("width") - margin.left - margin.right,
          height = +svg.attr("height") - margin.top - margin.bottom;

    // Create scales
    const xScale = d3.scaleTime()
                     .domain(d3.extent(cleanedData, d => d.date))
                     .range([0, width]);
    
    const yScale = d3.scaleLinear()
                     .domain([0, d3.max(cleanedData, d => d.price)])
                     .range([height, 0]);
    
    // Append the SVG group to transform the chart position
    const chartGroup = svg.append("g")
                          .attr("transform", `translate(${margin.left},${margin.top})`);

    // Draw the X axis
    chartGroup.append("g")
              .attr("transform", `translate(0,${height})`)
              .call(d3.axisBottom(xScale))
              .attr("color", "#fff"); // White color for dark backgrounds
    
    // Draw the Y axis
    chartGroup.append("g")
              .call(d3.axisLeft(yScale))
              .attr("color", "#fff"); // White color for dark backgrounds

    // Data line
    const line = d3.line()
                   .x(d => xScale(d.date))
                   .y(d => yScale(d.price));
    
    // Draw the line
    chartGroup.append("path")
              .datum(cleanedData)
              .attr("fill", "none")
              .attr("stroke", "yellow") // Use a bright color for demonstration, although not according to interpolateSpectral
              .attr("stroke-width", 1.5)
              .attr("d", line);
}

```

The visualization currently has the title:
```
Stock Prices Over Time
```
Keep this title the same, unless the code is changing in a way that
makes this title inaccurate.

## Dataset Description

### Filetype and Summary
The data is in this format: CSV

The title of the dataset is: Historical Stock Prices

Monthly closing prices of Microsoft, Amazon, and IBM stocks from January 2000 to March 2010.

### Structure
The data is in CSV format, where each row represents a stock's monthly closing price at the end of the given month. Each row contains three fields: the stock's symbol (`symbol`), the date (`date`), and the price (`price`) at which the stock closed at the end of that month. The `date` field is formatted as `Mon d yyyy`, where `Mon` is the three-letter abbreviation of the month. Before analysis, you may need to convert the `date` field into a date object and the `price` field into a numeric type.

### Fields

Each data point has these fields:
* `symbol`
* `date`
* `price`

Be sure to respect the capitalization and spaces in the above fields.

### Sample Data
Here is a small sample of the data:

```
symbol,date,price
MSFT,Jan 1 2000,39.81
MSFT,Feb 1 2000,36.35
MSFT,Mar 1 2000,43.22
MSFT,Apr 1 2000,28.37
MSFT,May 1 2000,25.45
MSFT,Jun 1 2000,32.54
MSFT,Jul 1 2000,28.4
MSFT,Aug 1 2000,28.4
MSFT,Sep 1 2000,24.53
MSFT,Oct 1 2000,28.02
MSFT,Nov 1 2000,23.34
MSFT,Dec 1 2000,17.65
MSFT,Jan 1 2001,24.84
MSFT,Feb 1 2001,24
MSFT,Mar 1 2001,22.25
MSFT,Apr 1 2001,27.56
MSFT,May 1 2001,28.14
MSFT,Jun 1 2001,29.7
MSFT,Jul 1 2001,26.93
MSFT,Aug 1 2001,23.21
MSFT,Sep 1 2001,20.82
MSFT,Oct 1 2001,23.65
MSFT,Nov 1 2001,26.12
MSFT,Dec 1 2001,26.95
MSFT,Jan 1 2002,25.92
MSFT,Feb 1 2002,23.73
MSFT,Mar 1 2002,24.53
MSFT,Apr 1 2002,21.26
MSFT,May 1 2002,20.71
MSFT,Jun 1 2002,22.25
MSFT,Jul 1 2002,19.52
MSFT,Aug 1 2002,19.97
MSFT,Sep 1 2002,17.79
MSFT,Oct 1 2002,21.75
MSFT,Nov 1 2002,23.46
MSFT,Dec 1 2002,21.03
MSFT,Jan 1 2003,19.31
MSFT,Feb 1 2003,19.34
MSFT,Mar 1 2003,19.76
MSFT,Apr 1 2003,20.87
MSFT,May 1 2003,20.09
MSFT,Jun 1 2003,20.93
MSFT,Jul 1 2003,21.56
MSFT,Aug 1 2003,21.65
MSFT,Sep 1 2003,22.6
```

## Task
We need to modify the code above to satisfy this user prompt:
---
Start the plot at 2005, index each of the stock prices so they're at 0 in 2005, and show the relative increase and decrease over time
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
The current title is "Stock Prices Over Time".
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
