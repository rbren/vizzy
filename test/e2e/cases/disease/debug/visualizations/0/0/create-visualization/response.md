# Perceived Media Coverage vs. Actual Causes of Death

## Plan
We will create a stacked bar chart comparing the perceived media coverage (Google search trends, The Guardian newspaper, and The New York Times) versus the actual data from the CDC for various causes of death. The process involves parsing the CSV data, sanitizing, and transforming it into a suitable format for D3 to visualize.

1. **Parsing the data**: Use `d3.csvParse` to parse the CSV string into an array of objects.
2. **Data Transformation**: Transform the parsed data into a format suitable for creating a stacked bar chart.
3. **Data Sanitization**: Ensure all data points are valid, i.e., no null or undefined values. Normalize if needed.
4. **Filtering**: If necessary, filter the dataset to a manageable size, though given the sample it seems all data can be displayed.
5. **Visualization**: Use D3 to create a stacked bar chart, taking the SVG's width and height into account.

## Code
```javascript
async function drawVisualization(svg, data) {
    // Parse the CSV data
    const parsedData = d3.csvParse(data);

    // Data transformation
    const causes = parsedData.map(d => d.cod);
    const sources = ["cdc", "google", "guardian", "nyt"];
    const stackedData = d3.stack().keys(sources)(parsedData);

    // Set up scales
    const xScale = d3.scaleBand()
        .domain(causes)
        .rangeRound([0, svg.attr('width')])
        .padding(0.1);
    const yScale = d3.scaleLinear()
        .domain([0, 1])
        .range([svg.attr('height'), 0]);
    const color = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, sources.length));

    // Axis
    const xAxis = d3.axisBottom(xScale).tickSizeOuter(0);
    const yAxis = d3.axisLeft(yScale).ticks(10, "%");

    // Append axes to the svg
    svg.append("g")
        .attr("transform", `translate(0,${svg.attr('height')})`)
        .call(xAxis)
        .selectAll("text")
        .style("text-anchor", "end")
        .attr("dx", "-.8em")
        .attr("dy", ".15em")
        .attr("transform", "rotate(-45)");
    svg.append("g").call(yAxis);

    // Create and fill the bars
    svg.append("g")
        .selectAll("g")
        .data(stackedData)
        .join("g")
          .attr("fill", d => color(d.key))
        .selectAll("rect")
        .data(d => d)
        .join("rect")
          .attr("x", d => xScale(d.data.cod))
          .attr("y", d => yScale(d[1]))
          .attr("height", d => yScale(d[0]) - yScale(d[1]))
          .attr("width", xScale.bandwidth());

    // Text color and font
    svg.selectAll("text").attr("fill", "#fff").attr("font-family", "sans-serif");

    // Title (SVG does not support rendering titles, guided by the style guide)
    // However, ensure labels for axes are in place for clarity
    svg.append("text")
      .attr("x", svg.attr('width') / 2)
      .attr("y", 0 - (10 / 2))
      .attr("text-anchor", "middle")
      .style("font-size", "16px")
      .style("text-decoration", "underline")
      .attr("fill", "#fff")
      .text("Causes of Death: Perceived vs. Actual");
}
```