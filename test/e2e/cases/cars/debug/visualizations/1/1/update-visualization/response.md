# Average Horsepower of Classic Cars Over Time

```javascript
async function drawVisualization(svg, dataStr) {
  const rawData = JSON.parse(dataStr);

  const filteredData = rawData.filter(d => d.Horsepower && d.Year)
                               .map(d => ({
                                 ...d,
                                 Year: new Date(d.Year).getFullYear(),
                                 Horsepower: +d.Horsepower
                               }));

  const groupedData = Array.from(d3.group(filteredData, d => d.Year), ([Year, values]) => ({
    Year,
    AverageHorsepower: d3.mean(values, v => v.Horsepower)
  }));

  const sortedData = groupedData.sort((a, b) => d3.ascending(a.Year, b.Year));

  const margin = {top: 20, right: 20, bottom: 30, left: 40},
        width = svg.attr('width') - margin.left - margin.right,
        height = svg.attr('height') - margin.top - margin.bottom;

  const g = svg.append("g").attr("transform", `translate(${margin.left},${margin.top})`);

  const x = d3.scaleLinear().range([0, width]);
  const y = d3.scaleLinear().range([height, 0]);

  x.domain(d3.extent(sortedData, d => d.Year));
  y.domain([0, d3.max(sortedData, d => d.AverageHorsepower)]);

  g.append("g")
   .attr("transform", `translate(0,${height})`)
   .call(d3.axisBottom(x).tickFormat(d3.format('d')))
   .style("color", "#fff");

  g.append("g")
   .call(d3.axisLeft(y))
   .style("color", "#fff");

  const line = d3.line()
                 .x(d => x(d.Year))
                 .y(d => y(d.AverageHorsepower));

  g.append("path")
   .data([sortedData])
   .attr("class", "line")
   .attr("d", line)
   .attr("fill", "none")
   .attr("stroke", "purple")
   .attr("stroke-width", "2px");

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