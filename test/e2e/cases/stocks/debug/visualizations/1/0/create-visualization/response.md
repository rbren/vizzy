# Stock Prices Over Time

## Plan

1. **Parse the data**: Given the CSV format, use `d3.csvParse` to convert the string data into an array of objects.
   
2. **Transform the data**: Map the parsed data to ensure `date` fields are converted to JavaScript Date objects and `price` values are treated as floats. This will facilitate time scale for the x-axis and linear scale for the y-axis.
   
3. **Structure of the transformed data**: Each object will have a `symbol`, a `date` (as a Date object), and a `price` (as a float).
   
4. **Sanitize the data**: Filter the data to remove any objects with null or missing values in any of the three fields (`symbol`, `date`, `price`). Throw an error if no valid data points remain after sanitation.
   
5. **Filter the data**: Since the task involves plotting prices over time for two random symbols, first group the data by `symbol` using `d3.group`. Then, select two random symbols to visualize. If the dataset for a chosen symbol contains more than 50 data points, downsample it to ensure a reasonable number of points.

## Code

```javascript
async function drawVisualization(svg, data) {
    const parseDate = d3.timeParse("%b %d %Y");
    const csvData = d3.csvParse(data, ({symbol, date, price}) => ({
        symbol,
        date: parseDate(date),
        price: +price
    }));

    // Sanitize data
    const sanitizedData = csvData.filter(d => d.symbol && d.date && !isNaN(d.price));
    if (sanitizedData.length === 0) throw new Error('No valid data points');

    // Randomly select two symbols
    const symbols = Array.from(new Set(sanitizedData.map(d => d.symbol)));
    if (symbols.length < 2) throw new Error('Insufficient symbols for comparison');
    const selectedSymbols = _.sampleSize(symbols, 2);

    // Filter the data for the selected symbols
    const symbolData = sanitizedData.filter(d => selectedSymbols.includes(d.symbol));
    
    // Group data by symbol
    const groupedData = d3.group(symbolData, d => d.symbol);

    const width = parseInt(svg.attr('width'));
    const height = parseInt(svg.attr('height'));
    const margin = { top: 20, right: 30, bottom: 20, left: 50 };

    const x = d3.scaleTime()
        .domain(d3.extent(symbolData, d => d.date))
        .range([margin.left, width - margin.right]);

    const y = d3.scaleLinear()
        .domain([0, d3.max(symbolData, d => d.price)])
        .range([height - margin.bottom, margin.top]);

    const color = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, selectedSymbols.length));

    svg.append("g")
       .attr("transform", `translate(0,${height - margin.bottom})`)
       .call(d3.axisBottom(x).ticks(width / 80).tickSizeOuter(0))
       .attr("color", "#fff");

    svg.append("g")
       .attr("transform", `translate(${margin.left},0)`)
       .call(d3.axisLeft(y))
       .attr("color", "#fff");

    const line = d3.line()
        .defined(d => !isNaN(d.price))
        .x(d => x(d.date))
        .y(d => y(d.price));

    groupedData.forEach((value, key) => {
        svg.append("path")
           .datum(value)
           .attr("fill", "none")
           .attr("stroke", color(key))
           .attr("stroke-width", 1.5)
           .attr("d", line);
    });
}
```