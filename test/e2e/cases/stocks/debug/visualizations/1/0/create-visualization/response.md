# Monthly Closing Prices of Stocks Over Time

## Plan

1. Parse the CSV string data using `d3.csvParse`.
2. Transform the data to ensure the `date` is a JavaScript Date object and the `price` is a number. Filter out any rows with null or invalid values in the process.
3. Decide which two random symbols to plot. Since the example is with Microsoft (MSFT), Amazon (AMZN), and IBM, pick two of these randomly using JavaScript's Math functions.
4. Use `d3.line` to create two line graphs for the selected symbols.
5. Utilize `d3.scaleTime` for the x-axis (time) and `d3.scaleLinear` for the y-axis (price).
6. Make sure to adjust all elements (axes, lines) to fit within the given SVG area, taking into account the margins.
7. Append all necessary elements to the SVG element passed to `drawVisualization`.

Below is an implementation outline with sample code for each step.

## Code

```javascript
async function drawVisualization(svg, data) {
    // Parse the CSV Data
    const parsedData = d3.csvParse(data, d => ({
        symbol: d.symbol,
        date: new Date(d.date),
        price: +d.price
    })).filter(d => !isNaN(d.date) && !isNaN(d.price));

    // Make sure there's data to plot
    if (parsedData.length === 0) {
        throw new Error("No valid data points found");
    }

    // Pick two random symbols
    const uniqueSymbols = [...new Set(parsedData.map(d => d.symbol))];
    const randomSymbols = _.sampleSize(uniqueSymbols, 2);

    // Filter data for these two symbols
    const filteredData = parsedData.filter(d => randomSymbols.includes(d.symbol));

    // Set up dimensions
    const margin = { top: 20, right: 30, bottom: 30, left: 40 };
    const width = +svg.attr('width') - margin.left - margin.right;
    const height = +svg.attr('height') - margin.top - margin.bottom;

    // Create scales
    const x = d3.scaleTime()
        .domain(d3.extent(filteredData, d => d.date))
        .range([0, width]);

    const y = d3.scaleLinear()
        .domain([0, d3.max(filteredData, d => d.price)]).nice()
        .range([height, 0]);

    const color = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, randomSymbols.length));

    // Add scales to axis
    const xAxis = d3.axisBottom(x);
    const yAxis = d3.axisLeft(y);

    // Prepare the data for line generation
    const line = d3.line()
        .x(d => x(d.date))
        .y(d => y(d.price));

    const lines = svg.append('g')
        .selectAll('.line')
        .data(randomSymbols)
        .enter().append('g')
        .attr("transform", `translate(${margin.left},${margin.top})`);

    lines.append("path")
        .attr("fill", "none")
        .attr("stroke", d => color(d))
        .attr("stroke-linejoin", "round")
        .attr("stroke-linecap", "round")
        .attr("stroke-width", 1.5)
        .attr("d", d => line(filteredData.filter(p => p.symbol === d)));

    // Append axes
    svg.append("g")
        .attr("transform", `translate(${margin.left},${height + margin.top})`)
        .call(xAxis);

    svg.append("g")
        .attr("transform", `translate(${margin.left},${margin.top})`)
        .call(yAxis);

    // Label axes
    svg.append("text")
        .attr("text-anchor", "end")
        .attr("x", width / 2 + margin.left)
        .attr("y", height + margin.top + 20)
        .attr("fill", "#fff")
        .text("Time");

    svg.append("text")
        .attr("text-anchor", "end")
        .attr("x", -height / 2 + margin.top)
        .attr("y", 15)
        .attr("transform", "rotate(-90)")
        .attr("fill", "#fff")
        .text("Price ($)");
}
```