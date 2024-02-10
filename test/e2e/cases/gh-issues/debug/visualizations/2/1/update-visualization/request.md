# Data Visualization Task

## Current Code
We currently have this D3 code to generate a visualization of the data:

```javascript

async function drawVisualization(svg, data) {
    const width = parseInt(svg.attr('width'));
    const height = parseInt(svg.attr('height'));
    const margin = { top: 30, right: 30, bottom: 70, left: 100 };

    const transformedData = parseAndTransformData(data);

    if (transformedData.length === 0) {
        throw new Error("No valid data points found.");
    }

    const xScale = d3.scaleBand()
        .range([margin.left, width - margin.right])
        .domain(transformedData.map(d => d.title))
        .padding(0.1);

    const yScale = d3.scaleLinear()
        .domain([0, d3.max(transformedData, d => d.reactionsCount)])
        .range([height - margin.bottom, margin.top]);

    svg.append("g")
        .attr("transform", `translate(0,${height - margin.bottom})`)
        .call(d3.axisBottom(xScale))
        .selectAll("text")
        .attr("transform", "translate(-10,0)rotate(-45)")
        .style("text-anchor", "end");

    svg.append("g")
        .attr("transform", `translate(${margin.left}, 0)`)
        .call(d3.axisLeft(yScale));

    svg.append("g")
        .selectAll("rect")
        .data(transformedData)
        .join("rect")
        .attr("x", d => xScale(d.title))
        .attr("y", d => yScale(d.reactionsCount))
        .attr("width", xScale.bandwidth())
        .attr("height", d => height - margin.bottom - yScale(d.reactionsCount))
        .attr("fill", d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, transformedData.length)));
}

```

The visualization currently has the title:
```
GitHub Issues with Most Reactions
```
Keep this title the same, unless the code is changing in a way that
makes this title inaccurate.

## Dataset Description

### Filetype and Summary
The data is in this format: JSON

The title of the dataset is: GitHub Issues for Kubernetes

Contains detailed information about issues and pull requests for the Kubernetes repository on GitHub, including metadata and labels.

### Structure
This data is an array of objects, each representing a GitHub issue or pull request for the Kubernetes repository. Each object contains fields such as `url`, `html_url`, `title`, and `labels`, among others. The `user` and `assignee` fields are nested objects containing information about the issue creator and assignee respectively, including their `login` and `avatar_url`. The `labels` field is an array of objects, where each object describes a label applied to the issue, including the label's `name` and `description`. Special considerations for preprocessing might include handling nested objects and arrays, and parsing the `labels` array to extract specific labels of interest.

### Fields

Each data point has these fields:
* `url`
* `repository_url`
* `labels_url`
* `comments_url`
* `events_url`
* `html_url`
* `id`
* `node_id`
* `number`
* `title`
* `user.login`
* `user.id`
* `user.node_id`
* `user.avatar_url`
* `user.url`
* `user.html_url`
* `user.followers_url`
* `user.following_url`
* `user.gists_url`
* `user.starred_url`
* `user.subscriptions_url`
* `user.organizations_url`
* `user.repos_url`
* `user.events_url`
* `user.received_events_url`
* `user.type`
* `user.site_admin`
* `labels[].id`
* `labels[].node_id`
* `labels[].url`
* `labels[].name`
* `labels[].color`
* `labels[].default`
* `labels[].description`
* `state`
* `locked`
* `assignee`
* `assignees[]`
* `assignee.login`
* `assignee.id`
* `assignee.node_id`
* `assignee.avatar_url`
* `assignee.url`
* `assignee.html_url`
* `assignee.followers_url`
* `assignee.following_url`
* `assignee.gists_url`
* `assignee.starred_url`
* `assignee.subscriptions_url`
* `assignee.organizations_url`
* `assignee.repos_url`
* `assignee.events_url`
* `assignee.received_events_url`
* `assignee.type`
* `assignee.site_admin`

Be sure to respect the capitalization and spaces in the above fields.

### Sample Data
Here is a small sample of the data:

```
[
  {
    "url": "https://api.github.com/repos/kubernetes/kubernetes/issues/122973",
    "repository_url": "https://api.github.com/repos/kubernetes/kubernetes",
    "labels_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/122973/labels{/name}",
    "comments_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/122973/comments",
    "events_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/122973/events",
    "html_url": "https://github.com/kubernetes/kubernetes/pull/122973",
    "id": 2101118275,
    "node_id": "PR_kwDOAToIks5lGpWn",
    "number": 122973,
    "title": "Missed a spot! testing verify-licenses.sh",
    "user": {
      "login": "dims",
      "id": 23304,
      "node_id": "MDQ6VXNlcjIzMzA0",
      "avatar_url": "https://avatars.githubusercontent.com/u/23304?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/dims",
      "html_url": "https://github.com/dims",
      "followers_url": "https://api.github.com/user
```

## Task
We need to modify the code above to satisfy this user prompt:
---
Display all the reaction emojis and counts for each issue
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
Be sure to put this title OUTSIDE the javascript block, above the first backticks.
The current title is "GitHub Issues with Most Reactions".
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
