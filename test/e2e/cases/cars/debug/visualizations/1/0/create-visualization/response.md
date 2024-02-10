# Average Horsepower of Classic Cars Over Time

## Plan

1. **Parse Data**: Use `JSON.parse` to convert the JSON string into JavaScript objects.

2. **Transform Data**:
   - Group data by `Year`.
   - Calculate average `Horsepower` for each year.
   - Sample code for transformation will be demonstrated.

3. **Structure of Transformed Data**: Array of objects, each object having `Year` and `AverageHorsepower`.

4. **Sanitize Data**: 
   - Remove entries with null or missing values in critical fields (`Year`, `Horsepower`).
   - Ensure `Year` is converted to a Date object, and `Horsepower` is a number.

5. **Filter Data**: 
   - If the result of step 3 is too large (more than 50 data points), sample it to maintain 50 points or fewer.
   - Provide sample code for filtering.

## Code

```javascript
async function drawVisualization(svg, dataStr) {
  // Parse the data from the JSON string
  const rawData = JSON.parse(dataStr);

  // Filter out entries with null or missing horsepower or year, and ensure types are correct
  const filteredData = rawData.filter(d => d.Horsepower && d.Year)
                               .map(d => ({
                                 ...d,
                                 Year: new Date(d.Year).getFullYear(),
                                 Horsepower: +d.Horsepower
                               }));

  // Group by Year and calculate average horsepower
  const groupedData = Array.from(d3.group(filteredData, d => d.Year), ([Year, values]) => ({
    Year,
    AverageHorsepower: d3.mean(values, v => v.Horsepower)
  }));

  // Sort by Year
  const sortedData = groupedData.sort((a, b) => d3.ascending(a.Year, b.Year));

  // Setup margins and graph dimensions
  const margin = {top: 20, right: 20, bottom: 30, left: 40},
        width = svg.attr('width') - margin.left - margin.right,
        height = svg.attr('height') - margin.top - margin.bottom;

  // Append the g element, accounting for margins
  const g = svg.append("g").attr("transform", `translate(${margin.left},${margin.top})`);

  // Set the ranges
  const x = d3.scaleLinear().range([0, width]);
  const y = d3.scaleLinear().range([height, 0]);

  // Scale the range of the data
  x.domain(d3.extent(sortedData, d => d.Year));
  y.domain([0, d3.max(sortedData, d => d.AverageHorsepower)]);

  // Add the X Axis
  g.append("g")
   .attr("transform", `translate(0,${height})`)
   .call(d3.axisBottom(x).tickFormat(d3.format('d'))) // Only whole numbers for years
   .style("color", "#fff"); // Make axis text white

  // Add the Y Axis
  g.append("g")
   .call(d3.axisLeft(y))
   .style("color", "#fff"); // Make axis text white

  // Add the line for average horsepower
  const line = d3.line()
                 .x(d => x(d.Year))
                 .y(d => y(d.AverageHorsepower));

  g.append("path")
   .data([sortedData])
   .attr("class", "line")
   .attr("d", line)
   .attr("fill", "none")
   .attr("stroke", "#fff")
   .attr("stroke-width", "2px");

  // Add labels
  svg.append("text")
     .attr("transform", "rotate(-90)")
     .attr("y", 0)
     .attr("x",0 - (height / 2))
     .attr("dy", "1em")
     .style("text-anchor", "middle")
     .text("Average Horsepower")
     .attr("fill", "#fff");

  svg.append("text")
     .attr("x", width / 2)
     .attr("y", height + margin.top + 20)
     .style("text-anchor", "middle")
     .text("Year")
     .attr("fill", "#fff");
}
```