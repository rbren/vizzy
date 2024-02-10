# Top 10 Music Genres by Latest Popularity

```javascript
async function drawVisualization(svg, data) {
    // Parse the CSV data
    const parsedData = d3.csvParse(data);

    // Get the latest date column name
    const latestDate = Object.keys(parsedData[0]).pop();

    // Transform data into an array of {genre, value} objects
    let transformedData = parsedData.map(d => ({
        genre: d.genre,
        value: +d[latestDate]
    })).filter(d => !isNaN(d.value)); // Remove any rows with NaN values

    // Sort by value in descending order to get the top 10
    transformedData.sort((a, b) => b.value - a.value);
    const top10Genres = transformedData.slice(0, 10);

    // Set dimensions and margins
    const margin = {top: 40, right: 20, bottom: 30, left: 100},
        width = parseFloat(svg.attr("width")) - margin.left - margin.right,
        height = parseFloat(svg.attr("height")) - margin.top - margin.bottom;

    // Append a 'g' element to the SVG, to hold the bars, 
    // translated to account for the margins
    const g = svg.append("g").attr("transform", `translate(${margin.left},${margin.top})`);

    // Set up the scales
    const xScale = d3.scaleLinear()
        .domain([0, d3.max(top10Genres, d => d.value)])
        .range([0, width]);
    
    const yScale = d3.scaleBand()
        .domain(top10Genres.map(d => d.genre))
        .range([0, height])
        .padding(0.1);

    // Create the bars
    g.selectAll(".bar")
        .data(top10Genres)
        .enter().append("rect")
        .attr("class", "bar")
        .attr("x", 0)
        .attr("y", d => yScale(d.genre))
        .attr("width", d => xScale(d.value))
        .attr("height", yScale.bandwidth())
        .attr("fill", d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, top10Genres.length)));

    // Add the x-axis
    g.append("g")
        .attr("transform", `translate(0,${height})`)
        .call(d3.axisBottom(xScale))
        .attr("color", "#fff"); // White color to stand out against a dark background

    // Add the y-axis
    g.append("g")
        .call(d3.axisLeft(yScale))
        .attr("color", "#fff"); // White color

    // Setting text styles
    svg.selectAll("text").attr("fill", "#fff").attr("font-family", "sans-serif");
}
```