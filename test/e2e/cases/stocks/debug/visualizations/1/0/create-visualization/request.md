# Data Visualization Task

## Task
Please write JavaScript that uses D3 to create a visualization that satisfies this user prompt:

---
Plot the price over time for two random symbols
---

Keep your visualization as simple as possible, while still satisfying the user prompt.
Avoid adding things like tooltips etc if the user hasn't asked for them.

## Plan

Please plan the implementation step-by-step. Include sample code wherever possible. Specifically:
* Describe how to parse the data string described below, e.g. using `d3.csvParse`, `JSON.parse`, or using an XML parser
* Describe how to transform the data into a shape and format that will be easiest to work with when creating the visualization, and provide sample code
* Describe the structure of the transformed data
* Describe how to sanitize the data, to ensure that any null or missing values are removed before being passed to D3
* Describe how to filter the data, to ensure that if there's a lot of data, only a reasonable number of bars/lines/etc (10-50) are presented to the user, and provide sample code

In all your sample code, be sure to respect the libraries that are available,
as specified in the Technical Details section.

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
