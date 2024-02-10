# Stock Prices Over Time

```javascript
async function drawVisualization(svg, data) {
    const parsedData = d3.csvParse(data);

    const transformedData = parsedData.map(row => {
        return {
            symbol: row.symbol,
            date: d3.timeParse("%b %d %Y")(row.date),
            price: +row.price
        };
    });

    // Filter data starting from 2005
    let filteredData = transformedData.filter(row => row.date >= new Date("Jan 1 2005"));

    // Group data by symbol to calculate base prices in 2005
    let basePrices = Array.from(d3.rollup(filteredData, v => v[0].price, d => d.symbol));

    // Index prices to start at 0 in 2005 by subtracting base price of 2005 for each symbol
    filteredData = filteredData.map(row => {
        return {
            symbol: row.symbol,
            date: row.date,
            price: row.price - basePrices.find(d => d[0] === row.symbol)[1]
        };
    });

    // Nest the data by symbol to create one line per group
    const nestedData = d3.group(filteredData, d => d.symbol);

    const margin = {top: 20, right: 25, bottom: 30, left: 60},
          width = +svg.attr("width") - margin.left - margin.right,
          height = +svg.attr("height") - margin.top - margin.bottom;

    const xScale = d3.scaleTime()
                     .domain(d3.extent(filteredData, d => d.date))
                     .range([0, width]);

    const yScale = d3.scaleLinear()
                     .domain([
                        d3.min(filteredData, d => d.price),
                        d3.max(filteredData, d => d.price)
                     ])
                     .range([height, 0]);

    const chartGroup = svg.append("g")
                          .attr("transform", `translate(${margin.left},${margin.top})`);

    chartGroup.append("g")
              .attr("transform", `translate(0,${height})`)
              .call(d3.axisBottom(xScale))
              .attr("color", "#fff");

    chartGroup.append("g")
              .call(d3.axisLeft(yScale))
              .attr("color", "#fff");

    const color = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, nestedData.size));

    nestedData.forEach((values, key) => {
        const line = d3.line()
                       .x(d => xScale(d.date))
                       .y(d => yScale(d.price));

        chartGroup.append("path")
                  .datum([...values])
                  .attr("fill", "none")
                  .attr("stroke", color(key))
                  .attr("stroke-width", 1.5)
                  .attr("d", line);
    });
}
```