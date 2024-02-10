# CO2 Difference Over Time

## Plan
The plan involves the following steps:
1. **Parsing the Data:** Use `d3.csvParse` to parse the CSV string.
2. **Transforming Data:** Transform the parsed data into an easier format for plotting, calculating the difference between `CO2` and `adjusted CO2` for each entry.
3. **Sanitizing Data:** Ensure data entries are complete and cast numerical fields from strings to numbers where necessary.
4. **Filtering Data:** If necessary, limit the data to a reasonable number of points to plot.
5. **Creating the Visualization:** Use D3 to plot the difference over time on an SVG element.
6. **Styling:** Ensure text and lines are white to contrast against a dark background.

## Code
```javascript
async function drawVisualization(svg, dataString) {
  // Parsing the CSV data
  const data = d3.csvParse(dataString, d => {
    return {
      date: d3.timeParse("%Y-%m-%d")(d.Date),
      co2: +d.CO2,
      adjustedCo2: +d["adjusted CO2"],
      difference: +d.CO2 - +d["adjusted CO2"]
    };
  });

  // Sanitizing the data: Remove entries with missing or null values
  const filteredData = data.filter(d => !isNaN(d.difference) && d.date);

  if (filteredData.length === 0) {
    throw new Error("No valid data points");
  }

  // Setting up dimensions
  const margin = {top: 20, right: 30, bottom: 30, left: 40},
      width = +svg.attr('width') - margin.left - margin.right,
      height = +svg.attr('height') - margin.top - margin.bottom;

  // Scales
  const x = d3.scaleTime()
    .domain(d3.extent(filteredData, d => d.date))
    .range([margin.left, width - margin.right]);

  const y = d3.scaleLinear()
    .domain([d3.min(filteredData, d => d.difference), d3.max(filteredData, d => d.difference)])
    .nice()
    .range([height - margin.bottom, margin.top]);

  // Axes
  const xAxis = g => g
    .attr("transform", `translate(0,${height - margin.bottom})`)
    .call(d3.axisBottom(x).ticks(width / 80).tickSizeOuter(0))
    .attr("color", "#fff");

  const yAxis = g => g
    .attr("transform", `translate(${margin.left},0)`)
    .call(d3.axisLeft(y))
    .call(g => g.select(".domain").remove())
    .call(g => g.select(".tick:last-of-type text").clone()
        .attr("x", 3)
        .attr("text-anchor", "start")
        .attr("font-weight", "bold")
        .text(data.y))
    .attr("color", "#fff");

  // Drawing the line
  const line = d3.line()
    .defined(d => !isNaN(d.difference))
    .x(d => x(d.date))
    .y(d => y(d.difference));

  svg.append("g")
    .call(xAxis);

  svg.append("g")
    .call(yAxis);

  svg.append("path")
    .datum(filteredData)
    .attr("fill", "none")
    .attr("stroke", "white")
    .attr("stroke-width", 1.5)
    .attr("stroke-linejoin", "round")
    .attr("stroke-linecap", "round")
    .attr("d", line);

  // Optional: Add axis labels, title, etc. here according to the requirements.
}
```