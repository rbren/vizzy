# Stock Prices Over Time for Microsoft, Apple, and Google

## Plan
1. Parse the CSV data using `d3.csvParse` to convert the string data into an array of objects. Each object will represent a row in the CSV.
2. Filter out any companies not requested (in this case, we'll keep only Microsoft, Apple, and Google).
3. Convert the `date` field from string to JavaScript `Date` objects and ensure the `price` field is a float. This will help in plotting the data over time and ensure accuracy in the visualization.
4. Group the data by the `symbol` to make it easier to plot each company's data as a separate line.
5. Use D3 to scale the data to fit within the SVG dimensions, considering the domain of dates and prices.
6. Draw axes, and for each group (company), plot a line representing the stock prices over time.

## Code
```javascript
async function drawVisualization(svg, data) {
    // Parse the CSV data
    const parsedData = d3.csvParse(data);

    // Filter for Microsoft, Apple, and Google
    const filteredData = parsedData.filter(d => d.symbol === 'MSFT' || d.symbol === 'AAPL' || d.symbol === 'GOOGL');

    // Convert date strings to Date objects and price strings to numbers
    const transformedData = filteredData.map(d => {
        return { symbol: d.symbol, date: new Date(d.date), price: +d.price };
    });

    // Group data by symbol
    const groupedData = d3.group(transformedData, d => d.symbol);

    // Set up the SVG dimensions
    const width = parseInt(svg.attr('width'));
    const height = parseInt(svg.attr('height'));
    const margin = { top: 20, right: 20, bottom: 30, left: 50 };
    const chartWidth = width - margin.left - margin.right;
    const chartHeight = height - margin.top - margin.bottom;

    // Create scales
    const xScale = d3.scaleTime()
        .domain(d3.extent(transformedData, d => d.date))
        .range([0, chartWidth]);

    const yScale = d3.scaleLinear()
        .domain([0, d3.max(transformedData, d => d.price)])
        .range([chartHeight, 0]);

    // Append the G element
    const g = svg.append('g')
        .attr('transform', `translate(${margin.left},${margin.top})`);

    // Create and append the X axis
    g.append('g')
        .attr('transform', `translate(0,${chartHeight})`)
        .call(d3.axisBottom(xScale))
        .attr('color', '#fff');

    // Create and append the Y axis
    g.append('g')
        .call(d3.axisLeft(yScale))
        .attr('color', '#fff');

    // Line generator
    const line = d3.line()
        .x(d => xScale(d.date))
        .y(d => yScale(d.price))
        .curve(d3.curveMonotoneX);

    // Set the color scale for different symbols
    const color = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, groupedData.size));

    // Plot line for each group
    groupedData.forEach((value, key) => {
        g.append('path')
            .datum(value)
            .attr('fill', 'none')
            .attr('stroke', color(key))
            .attr('stroke-width', 2)
            .attr('d', line);
    });
}
```