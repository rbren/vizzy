# Data Visualization Task

## Task
Please write JavaScript that uses D3 to create a visualization that satisfies this user prompt:

---
Show a list of the ten most popular genres as of the latest date
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

The title of the dataset is: Music Genre Popularity Over Time

Tracks the popularity of various music genres over multiple dates spanning from April 23, 2016 to August 9, 2023.

### Structure
The data is structured as a CSV where each row represents a music genre and each column represents a date. The first row contains the column headers, which are dates in MM/DD/YYYY format starting from '4/23/2016' and ending with '8/9/2023'. Each subsequent row starts with the genre name followed by the popularity scores (or rankings) of that genre on each date. Scores are numeric. Some cells in the table are empty, indicating missing data points for those genre-date combinations. Preprocessing steps might include filling missing values and converting date strings into a date-time format for time series analysis.

### Fields

Each data point has these fields:
* `genre`
* `4/23/2016`
* `6/25/2016`
* `2/17/2017`
* `4/19/2017`
* `5/11/2017`
* `7/21/2017`
* `9/12/2017`
* `11/12/2017`
* `12/20/2017`
* `1/12/2018`
* `3/6/2018`
* `5/1/2018`
* `6/30/2018`
* `9/29/2018`
* `12/12/2018`
* `1/19/2019`
* `3/30/2019`
* `6/19/2019`
* `9/8/2019`
* `12/12/2019`
* `1/6/2020`
* `3/2/2020`
* `6/29/2020`
* `9/21/2020`
* `12/4/2020`
* `1/12/2021`
* `3/5/2021`
* `6/15/2021`
* `9/13/2021`
* `12/3/2021`
* `6/15/2022`
* `9/22/2022`
* `1/27/2023`
* `4/11/2023`
* `6/9/2023`
* `8/9/2023`

Be sure to respect the capitalization and spaces in the above fields.

### Sample Data
Here is a small sample of the data:

```
genre,4/23/2016,6/25/2016,2/17/2017,4/19/2017,5/11/2017,7/21/2017,9/12/2017,11/12/2017,12/20/2017,1/12/2018,3/6/2018,5/1/2018,6/30/2018,9/29/2018,12/12/2018,1/19/2019,3/30/2019,6/19/2019,9/8/2019,12/12/2019,1/6/2020,3/2/2020,6/29/2020,9/21/2020,12/4/2020,1/12/2021,3/5/2021,6/15/2021,9/13/2021,12/3/2021,6/15/2022,9/22/2022,1/27/2023,4/11/2023,6/9/2023,8/9/2023
pop,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1
rap,3,4,4,4,4,4,4,4,4,3,2,2,2,3,3,3,4,3,3,4,3,3,3,3,4,3,4,3,3,3,3,3,3,3,2,2
rock,5,5,7,7,8,7,7,7,6,6,7,7,7,7,7,5,6,6,5,6,6,7,5,5,6,6,5,6,6,4,5,4,4,4,3,3
urbano latino,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,6,6,4,4
hip hop,10,8,11,9,7,10,10,13,13,10,6,6,5,6,6,6,7,7,6,8,8,10,8,8,9,9,9,8,8,6,7,6,5,5,6,5
trap latino,,,,85,94,85,88,64,72,65,63,49,48,66,61,43,54,46,53,55,53,17,39,44,10,10,10,9,9,8,6,5,15,16,5,6
reggaeton,68,60,40,39,38,29,29,29,33,23,19,15,15,22,21,20,22,23,21,11,10,8,10,11,11,13,12,12,10,10,8,7,7,7,7,7
filmi,551,562,570,570,563,574,576,568,580
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
