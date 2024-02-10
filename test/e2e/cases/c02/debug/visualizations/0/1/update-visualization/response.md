# Atmospheric CO2 Concentrations Over Time

```javascript
async function drawVisualization(svg, data) {
  const parsedData = d3.csvParse(data);
  
  const transformedData = parsedData.map(d => {
    return {
      date: new Date(d.Date),
      CO2: +d.CO2,
      adjustedCO2: +d['adjusted CO2']
    };
  }).filter(d => {
    return !isNaN(d.CO2) && d.CO2 !== null && !isNaN(d.adjustedCO2) && d.adjustedCO2 !== null;
  });
  
  if (transformedData.length === 0) {
    throw new Error("No valid data points.");
  }
  
  const margin = {top: 20, right: 30, bottom: 30, left: 50},
      width = svg.attr('width') - margin.left - margin.right,
      height = svg.attr('height') - margin.top - margin.bottom;
  
  const g = svg.append("g").attr("transform", `translate(${margin.left},${margin.top})`);
  
  const xScale = d3.scaleTime()
                   .domain(d3.extent(transformedData, d => d.date))
                   .range([0, width]);
                   
  const yScale = d3.scaleLinear()
                   .domain([250, d3.max(transformedData, d => d.CO2)])
                   .range([height, 0]);
                   
  g.append("g")
   .call(d3.axisLeft(yScale))
   .attr("color", "#fff"); // Fixing incorrect attribute for font color
  
  g.append("g")
   .attr("transform", `translate(0,${height})`)
   .call(d3.axisBottom(xScale))
   .attr("color", "#fff"); // Fixing incorrect attribute for font color
  
  const line = d3.line()
                 .x(d => xScale(d.date))
                 .y(d => yScale(d.CO2));
  
  g.append("path")
   .datum(transformedData)
   .attr("fill", "none")
   .attr("stroke", "white")
   .attr("stroke-width", 1.5)
   .attr("d", line);
}
```