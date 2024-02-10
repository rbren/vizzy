# CO2 Levels Over Time

## Plan

To create a visualization plotting the difference between CO2 and adjusted CO2 levels over time, the following steps will be detailed:

1. **Parse the Data**: Since the provided data is a CSV format string, we will use `d3.csvParse` to parse it into a JavaScript object.

2. **Transform the Data**: We need to transform the data to include a new field that represents the difference between CO2 and adjusted CO2 levels for each data point.

3. **Sanitize the Data**: Ensure all data points are complete and valid, removing any that contain null or missing values.

4. **Filter the Data**: If necessary, filter the data to ensure the visualization does not get overcrowded. We will achieve this by selecting a subset of the data to display.

5. **Plot the Data**: Create a line chart where the x-axis represents time (Date) and the y-axis represents the CO2 difference.

### Transform the Data

To calculate the difference and transform the data, we can use the following code:

```javascript
const dataFormatted = data.map(d => ({
    ...d,
    Date: d3.timeParse("%Y-%m-%d")(d.Date),
    CO2: +d.CO2, 
    adjustedCO2: +d.adjustedCO2,
    difference: +d.CO2 - +d.adjustedCO2
}));
```

### Sanitize the Data

To ensure data integrity, we will filter out any records with missing or null values:

```javascript
const sanitizedData = dataFormatted.filter(d => d.CO2 && d.adjustedCO2);
```

### Filter the Data

We will not specifically filter the data in this case because we assume that a detailed analysis is desired over time. If data volume necessitates filtering, one could downsample or select a more specific time range.

## Code

```javascript
async function drawVisualization(svg, data) {
    const margin = {top: 20, right: 30, bottom: 30, left: 50},
          width = +svg.attr("width") - margin.left - margin.right,
          height = +svg.attr("height") - margin.top - margin.bottom;
    
    // Parse and transform the data
    const parsedData = d3.csvParse(data, d => ({
        Date: d3.timeParse("%Y-%m-%d")(d.Date),
        CO2: +d.CO2,
        adjustedCO2: +d["adjusted CO2"],
        difference: +d.CO2 - +d["adjusted CO2"]
    })).filter(d => d.Date && !isNaN(d.difference)); // Sanitize data

    // Create scales
    const x = d3.scaleTime()
        .domain(d3.extent(parsedData, d => d.Date))
        .range([0, width]);
    
    const y = d3.scaleLinear()
        .domain([d3.min(parsedData, d => d.difference), d3.max(parsedData, d => d.difference)])
        .nice()
        .range([height, 0]);
    
    // Append the svg object to the body of the page
     svg = svg.append("g")
           .attr("transform", `translate(${margin.left},${margin.top})`);

    // Add X axis
    svg.append("g")
        .attr("transform", `translate(0,${height})`)
        .call(d3.axisBottom(x))
        .attr("stroke", "#fff");
    
    // Add Y axis
    svg.append("g")
        .call(d3.axisLeft(y))
        .attr("stroke", "#fff");
    
    // Add the line
    svg.append("path")
        .datum(parsedData)
        .attr("fill", "none")
        .attr("stroke", "steelblue")
        .attr("stroke-width", 1.5)
        .attr("d", d3.line()
            .x(d => x(d.Date))
            .y(d => y(d.difference)));

    // Add labels
    svg.append("text")
        .attr("text-anchor", "end")
        .attr("stroke", "#fff")
        .attr("x", width)
        .attr("y", height - 6)
        .text("Date");
    
    svg.append("text")
        .attr("text-anchor", "end")
        .attr("transform", "rotate(-90)")
        .attr("stroke", "#fff")
        .attr("y", 6)
        .attr("dy", ".75em")
        .text("Difference in CO2 Levels (ppm)");
}
```