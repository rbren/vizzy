# Historical CO2 Concentrations Over Time

```javascript
async function drawVisualization(svg, data) {
    let parsedData = d3.csvParse(data);

    parsedData = parsedData.filter(row => row.Date && row.CO2 && row['adjusted CO2'])
                           .map(row => ({
                                date: d3.timeParse('%Y-%m-%d')(row.Date),
                                co2: +row.CO2,
                                adjustedCo2: +row['adjusted CO2']
                            }));

    if (parsedData.length === 0) throw new Error("No valid data points found.");
    
    const margin = { top: 20, right: 30, bottom: 50, left: 60 },
          width = svg.attr('width') - margin.left - margin.right,
          height = svg.attr('height') - margin.top - margin.bottom;

    const xScale = d3.scaleTime()
                     .domain(d3.extent(parsedData, d => d.date))
                     .range([0, width]);

    const yScale = d3.scaleLinear()
                     .domain([250, d3.max(parsedData, d => Math.max(d.co2, d.adjustedCo2))])
                     .nice()
                     .range([height, 0]);

    const xAxis = d3.axisBottom(xScale).ticks(width > 500 ? 16 : 8);
    const yAxis = d3.axisLeft(yScale);

    svg.append("g")
       .attr("transform", `translate(${margin.left},${height + margin.top})`)
       .call(xAxis)
       .style("color", "#fff");

    svg.append("g")
       .attr("transform", `translate(${margin.left},${margin.top})`)
       .call(yAxis)
       .style("color", "#fff");

    const lineGeneratorCO2 = d3.line()
                                .x(d => xScale(d.date))
                                .y(d => yScale(d.co2));

    const lineGeneratorAdjustedCO2 = d3.line()
                                       .x(d => xScale(d.date))
                                       .y(d => yScale(d.adjustedCo2));

    svg.append('path')
       .datum(parsedData)
       .attr('d', lineGeneratorCO2)
       .attr('fill', 'none')
       .attr('stroke', 'limegreen')
       .attr('stroke-width', 2)
       .attr('transform', `translate(${margin.left},${margin.top})`);

    svg.append('path')
       .datum(parsedData)
       .attr('d', lineGeneratorAdjustedCO2)
       .attr('fill', 'none')
       .attr('stroke', 'orange')
       .attr('stroke-width', 2)
       .attr('transform', `translate(${margin.left},${margin.top})`);

    svg.append("text")
       .attr("transform", `translate(${width / 2 + margin.left},${height + margin.top + 40})`)
       .style("text-anchor", "middle")
       .style("fill", "#fff")
       .text("Time");

    svg.append("text")
       .attr("transform", "rotate(-90)")
       .attr("y", 0)
       .attr("x", -(height / 2) - margin.top)
       .attr("dy", "1em")
       .style("text-anchor", "middle")
       .style("fill", "#fff")
       .text("CO2 Concentrations (ppm)");
}
```