# Data Visualization Task

## Current Code
We currently have this D3 code to generate a visualization of the data:

```javascript

async function drawVisualization(svg, dataString) {
  // Parsing data
  const data = JSON.parse(dataString);

  // Transforming data: extracting latest fertility rate for each country
  const latestFertilityByCountry = data.reduce((acc, item) => {
    const { country, year, fertility } = item;
    if (!acc[country] || acc[country].year < year) {
      acc[country] = { year, fertility };
    }
    return acc;
  }, {});

  // Sanitizing data: removing entries without a valid fertility rate
  const sanitizedData = Object.entries(latestFertilityByCountry)
    .filter(([_, { fertility }]) => fertility != null)
    .map(([country, { fertility }]) => ({ country, fertility }));

  // Loading map data
  const mapData = await d3.json("https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json");

  // Setting up projection and path generator
  const projection = d3.geoNaturalEarth1().fitSize([+svg.attr("width"), +svg.attr("height")], topojson.feature(mapData, mapData.objects.countries));
  const path = d3.geoPath().projection(projection);

  // Setting up color scale
  const fertilityValues = sanitizedData.map(d => d.fertility);
  const colorScale = d3.scaleLinear()
    .domain([d3.min(fertilityValues), d3.max(fertilityValues)])
    .range(["#fff", "red"]) // use interpolateSpectral for a more detailed color scheme as required
    .interpolate(d3.interpolateSpectral);

  // Drawing countries
  svg.selectAll("path")
    .data(topojson.feature(mapData, mapData.objects.countries).features)
    .join("path")
    .attr("fill", d => {
      const countryData = sanitizedData.find(c => c.country === d.properties.name);
      return countryData ? colorScale(countryData.fertility) : "#ccc";
    })
    .attr("d", path)
    .attr("stroke", "white")
    .style("stroke-width", 0.5)
    .attr("title", d => d.properties.name); // Adding country name as title for accessibility

  // Ensure text and other elements are white for contrast
  svg.selectAll("text")
    .style("fill", "#fff")
    .style("font-family", "sans-serif");
}

```

The visualization currently has the title:
```
Global Fertility Rate Map
```
Keep this title the same, unless the code is changing in a way that
makes this title inaccurate.

## Dataset Description

### Filetype and Summary
The data is in this format: JSON

The title of the dataset is: Global Fertility and Life Expectancy

Tracks changes in fertility rates and life expectancy across different countries over years

### Structure
The data is structured as a JSON array with each object representing a data point for a country in a given year. Each object contains the fields `year`, `fertility`, `life_expect`, and `country`. Optionally, objects may contain previous (`p_`) and next (`n_`) values for both `fertility` and `life_expect`, indicating trends over time. When reading this data, special attention should be given to these predictive and historical values to understand temporal trends. Preprocessing may involve filtering by country or time range, and handling missing `p_` and `n_` fields when they are absent.

### Fields

Each data point has these fields:
* `year`
* `fertility`
* `life_expect`
* `n_fertility`
* `n_life_expect`
* `p_fertility`
* `p_life_expect`
* `country`

Be sure to respect the capitalization and spaces in the above fields.

### Sample Data
Here is a small sample of the data:

```
[
    {
        "_comment": "Data courtesy of Gapminder.org",
        "year": 1955,
        "fertility": 7.7,
        "life_expect": 30.332,
        "n_fertility": 7.7,
        "n_life_expect": 31.997,
        "country": "Afghanistan"
    },
    {
        "year": 1960,
        "fertility": 7.7,
        "life_expect": 31.997,
        "p_fertility": 7.7,
        "n_fertility": 7.7,
        "p_life_expect": 30.332,
        "n_life_expect": 34.02,
        "country": "Afghanistan"
    },
    {
        "year": 1965,
        "fertility": 7.7,
        "life_expect": 34.02,
        "p_fertility": 7.7,
        "n_fertility": 7.7,
        "p_life_expect": 31.997,
        "n_life_expect": 36.088,
        "country": "Afghanistan"
    },
    {
        "year": 1970,
        "fertility": 7.7,
        "life_expect": 36.088,
        "p_fertility": 7.7,
        "n_fertility": 7.7,
        "p_life_expect": 34.02,
        "n_life_expect": 38.438,
        "country": "Afghanistan"
    },
    {
        "year"
```

## Task
We need to modify the code above to satisfy this user prompt:
---
Make very fertile countries green, and very infertile countries deep purple
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
The current title is "Global Fertility Rate Map".
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
