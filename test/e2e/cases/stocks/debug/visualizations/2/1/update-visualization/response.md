# Stock Prices Over Time for Microsoft, Apple, and Google

```javascript
async function drawVisualization(svg, data) {
    // Parse the CSV data
    const parsedData = d3.csvParse(data);

    // Filter for Microsoft, Apple, and Google, and start from 2005 onwards
     const filteredData = parsedData.filter(d => (d.symbol === 'MSFT' || d.symbol === 'AAPL' || d.symbol === 'GOOGL') && new Date(d.date) >= new Date('Jan 1 2005'));

    // Convert date strings to Date objects and price strings to numbers
    let transformedData = filteredData.map(d => ({
        symbol: d.symbol,
        date: new Date(d.date),
        price: +d.price
    }));

    // Group data by symbol to find base prices in 2005
    const basePrices = new Map();
    const groupedDataBySymbol = d3.group(transformedData, d => d.symbol);
    groupedDataBySymbol.forEach((value, key) => {
        const basePrice = value.find(d => d.date.getFullYear() === 2005).price;
        basePrices.set(key, basePrice);
    });

    // Index prices to start at 0 in 2005
    transformedData = transformedData.map(d => {
        return { symbol: d.symbol, date: d.date, price: d.price - basePrices.get(d.symbol) };
    });

    // Group data by symbol for plotting
    const groupedData = d3.group(transformedData, d => d.symbol);

    // Set up the SVG dimensions
    const width = +svg.attr('width');
    const height = +svg.attr('height');
    const margin = { top: 20, right: 20, bottom: 30, left: 50 };
    const chartWidth = width - margin.left - margin.right;
    const chartHeight = height - margin.top - margin.bottom;

    // Create scales
    const xScale = d3.scaleTime()
        .domain(d3.extent(transformedData, d => d.date))
        .range([0, chartWidth]);

    const yScale = d3.scaleLinear()
        .domain([d3.min(transformedData, d => d.price), d3.max(transformedData, d => d.price)])
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