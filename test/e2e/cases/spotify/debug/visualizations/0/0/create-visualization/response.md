# Music Genre Popularity Over Time

```javascript
async function drawVisualization(svg, data) {
    // Parse the CSV data
    const parsedData = d3.csvParse(data);
    
    // Transform the data: Convert date strings to date objects, ensure values are numbers
    const transformedData = parsedData.map(d => {
        const result = { genre: d.genre };
        for (const key of Object.keys(d).filter(k => k !== 'genre')) {
            result[key] = { date: d3.timeParse("%m/%d/%Y")(key), value: +d[key] };
        }
        return result;
    });
    
    // Extract all dates for the x-axis and flatten genre values for the y-axis
    const dates = _.uniq(_.flatten(transformedData.map(d => Object.values(d).map(v => v.date))));
    const values = _.flatten(transformedData.map(d => Object.values(d).map(v => v.value)));
    
    // Filter out undefined or null values
    const filteredData = transformedData.map(d => {
        return {
            genre: d.genre,
            values: Object.values(d).filter(v => v.date && !isNaN(v.value) && v.value !== null)
        };
    }).filter(d => d.values.length > 0);
    
    // Create scales
    const xScale = d3.scaleTime()
                     .domain(d3.extent(dates))
                     .range([40, +svg.attr('width') - 40]);
    const yScale = d3.scaleLinear()
                     .domain([0, d3.max(values)])
                     .range([+svg.attr('height') - 40, 20]);
    
    // Draw axes
    svg.append("g")
       .attr("transform", `translate(0,${+svg.attr('height') - 40})`)
       .call(d3.axisBottom(xScale))
       .attr("color", "#fff");
    svg.append("g")
       .attr("transform", "translate(40,0)")
       .call(d3.axisLeft(yScale))
       .attr("color", "#fff");
    
    // Draw lines
    const line = d3.line()
                   .x(d => xScale(d.date))
                   .y(d => yScale(d.value))
                   .curve(d3.curveMonotoneX);
    
    const colorScale = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, filteredData.length));
    
    svg.selectAll(".line")
       .data(filteredData)
       .enter()
       .append("path")
       .attr("fill", "none")
       .attr("stroke", (d, i) => colorScale(i))
       .attr("stroke-width", 2)
       .attr("d", d => line(d.values))
       .attr("data-genre", d => d.genre);
    
    // Draw genre labels at the end of each line
    filteredData.forEach(genreData => {
        svg.append("text")
           .attr("x", xScale(d3.max(genreData.values, d => d.date)) + 5)
           .attr("y", yScale(genreData.values[genreData.values.length - 1].value))
           .text(genreData.genre)
           .attr("fill", "#fff")
           .style("font-size", "10px")
           .style("font-family", "sans-serif");
    });
}
```