# Average Horsepower of Vehicles Over Time

```javascript
async function drawVisualization(svg, data) {
  const jsonData = JSON.parse(data);

  const filteredData = jsonData.filter(d => d.Horsepower != null && d.Year != null);
  
  filteredData.forEach(d => {
    d.Year = parseInt(d.Year.substring(0, 4));
  });

  const averageHorsepowerByYear = Array.from(d3.rollup(filteredData, v => d3.mean(v, d => d.Horsepower), d => d.Year),
    ([Year, Horsepower]) => ({Year, Horsepower}));

  const margin = {top: 20, right: 20, bottom: 30, left: 50},
      width = +svg.attr('width') - margin.left - margin.right,
      height = +svg.attr('height') - margin.top - margin.bottom,
      g = svg.append("g").attr("transform", `translate(${margin.left},${margin.top})`);

  const x = d3.scaleLinear()
              .rangeRound([0, width])
              .domain(d3.extent(averageHorsepowerByYear, d => d.Year));
  
  const y = d3.scaleLinear()
              .rangeRound([height, 0])
              .domain([0, d3.max(averageHorsepowerByYear, d => d.Horsepower)]);

  const line = d3.line()
                 .x(d => x(d.Year))
                 .y(d => y(d.Horsepower));

  g.append("g")
    .attr("transform", `translate(0,${height})`)
    .call(d3.axisBottom(x).tickFormat(d3.format("d")))
    .select(".domain")
    .remove();

  g.append("g")
    .call(d3.axisLeft(y))
    .append("text")
    .attr("fill", "#fff")
    .attr("transform", "rotate(-90)")
    .attr("y", 6)
    .attr("dy", "0.71em")
    .attr("text-anchor", "end")
    .text("Horsepower");

  g.append("path")
    .datum(averageHorsepowerByYear)
    .attr("fill", "none")
    .attr("stroke", "purple") // Changed the line color to purple to satisfy the user prompt
    .attr("stroke-linejoin", "round")
    .attr("stroke-linecap", "round")
    .attr("stroke-width", 1.5)
    .attr("d", line);
}
```