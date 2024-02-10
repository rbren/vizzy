# Stock Prices Over Time

## Plan

1. **Parse the CSV Data**: I'll use `d3.csvParse()` function to parse the input `data` string into JavaScript objects.
2. **Transform the Data**: Convert the `date` field from string to JavaScript Date objects and ensure `price` is a floating point number.
3. **Sanitize the Data**: Filter out any rows with null, missing values, or bad types.
4. **Data Structure and Filtering**: Since the user wants stock prices over time, we'll make sure to have data points for each month available. If too many data points, we'll select a subset.
5. **Draw the Visualization**: Use D3 to plot this data on an SVG.

### Sample code to transform data

```javascript
const parsedData = d3.csvParse(data);

// Transform dates and prices; filter out invalid data
const transformedData = parsedData
  .filter(d => d.symbol && d.date && d.price)  // Basic existence check
  .map(d => ({
    symbol: d.symbol,
    date: new Date(d.date),
    price: +d.price  // Convert string to number
  }));
```

## Code

```javascript
async function drawVisualization(svg, data) {
  const margin = {top: 20, right: 20, bottom: 30, left: 50},
      width = +svg.attr("width") - margin.left - margin.right,
      height = +svg.attr("height") - margin.top - margin.bottom;

  const parsedData = d3.csvParse(data);

  // Basic data cleaning
  const cleanedData = parsedData.filter(d => d.symbol && d.date && !isNaN(Date.parse(d.date)) && d.price && !isNaN(+d.price))
    .map(d => ({
      symbol: d.symbol,
      date: new Date(d.date),
      price: +d.price
    }));

  // Ensure there is data to plot
  if (cleanedData.length === 0) {
    throw new Error("No valid data points to plot.");
  }

  // Group data by symbol
  const groupedData = d3.group(cleanedData, d => d.symbol);

  const x = d3.scaleTime()
    .domain(d3.extent(cleanedData, d => d.date))
    .range([margin.left, width - margin.right]);

  const y = d3.scaleLinear()
    .domain([0, d3.max(cleanedData, d => d.price)]).nice()
    .range([height - margin.bottom, margin.top]);

  const line = d3.line()
    .defined(d => !isNaN(d.price))
    .x(d => x(d.date))
    .y(d => y(d.price));

  const xAxis = g => g
    .attr("transform", `translate(0,${height - margin.bottom})`)
    .call(d3.axisBottom(x).ticks(width / 80).tickSizeOuter(0));

  const yAxis = g => g
    .attr("transform", `translate(${margin.left},0)`)
    .call(d3.axisLeft(y))
    .call(g => g.select(".domain").remove())
    .call(g => g.select(".tick:last-of-type text").clone()
        .attr("x", 3)
        .attr("text-anchor", "start")
        .attr("font-weight", "bold")
        .text(cleanedData.price));

  svg.append("g")
    .call(xAxis);

  svg.append("g")
    .call(yAxis);

  const colors = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, groupedData.size));

  for (const [symbol, data] of groupedData) {
    svg.append("path")
      .datum(data)
      .attr("fill", "none")
      .attr("stroke", colors(symbol))
      .attr("stroke-width", 1.5)
      .attr("d", line);
  }

  const text = svg.selectAll("text")
    .data(groupedData.keys())
    .enter().append("text")
    .attr("transform", (d, i) => `translate(${width - margin.right},${margin.top + i * 20})`)
    .attr("font-family", "sans-serif")
    .attr("font-size", 12)
    .attr("fill", d => colors(d))
    .text(d => d);
}
```