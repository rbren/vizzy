# Horsepower vs Weight for Vehicles

## Plan
To construct the scatter plot of `Horsepower` vs `Weight_in_lbs`, we will follow these steps:

1. **Parse the JSON data**: Use the native JavaScript function `JSON.parse` to convert the string data into an array of objects.
2. **Data Transformation**: Create a new array containing only the necessary fields (`Horsepower` and `Weight_in_lbs`) for each vehicle, removing any objects that have null or missing values in those fields.
3. **Data Sanitization**: Ensure all data points in the `Horsepower` and `Weight_in_lbs` fields are of the correct type (numbers) and remove any invalid entries.
4. **Draw the Visualization**: Using D3, draw the scatter plot on the provided SVG element, with `Horsepower` on the x-axis and `Weight_in_lbs` on the y-axis. Adjust the scales appropriately to fit the SVG dimensions.

Sample code snippets illustrating each step will be included.

## Code
```javascript
async function drawVisualization(svg, data) {
  // Step 1: Parse the JSON data
  const vehicles = JSON.parse(data);

  // Step 2 & 3: Transform and sanitize the data
  const filteredData = vehicles.filter(v => v.Horsepower !== null && v.Weight_in_lbs !== null && !isNaN(v.Horsepower) && !isNaN(v.Weight_in_lbs)).map(v => ({
    Horsepower: +v.Horsepower,
    Weight_in_lbs: +v.Weight_in_lbs
  }));

  // Throw an error if there are no valid data points
  if (filteredData.length === 0) {
    throw new Error("No valid data points");
  }

  // Drawing starts here
  const margin = {top: 20, right: 20, bottom: 30, left: 40},
        width = +svg.attr('width') - margin.left - margin.right,
        height = +svg.attr('height') - margin.top - margin.bottom;

  // Append a 'g' element to svg
  const g = svg.append('g').attr('transform', `translate(${margin.left},${margin.top})`);

  // X and Y scales
  const x = d3.scaleLinear()
              .domain(d3.extent(filteredData, d => d.Horsepower))
              .rangeRound([0, width]);
  const y = d3.scaleLinear()
              .domain(d3.extent(filteredData, d => d.Weight_in_lbs))
              .rangeRound([height, 0]);

  // Append X axis
  g.append('g')
    .attr('transform', `translate(0,${height})`)
    .call(d3.axisBottom(x))
    .append('text')
    .attr('fill', '#fff')
    .attr('x', width)
    .attr('dy', '-0.5em')
    .attr('text-anchor', 'end')
    .text('Horsepower');

  // Append Y axis
  g.append('g')
    .call(d3.axisLeft(y))
    .append('text')
    .attr('fill', '#fff')
    .attr('transform', 'rotate(-90)')
    .attr('y', 6)
    .attr('dy', '0.71em')
    .attr('text-anchor', 'end')
    .text('Weight (lbs)');

  // Add the scatterplot points
  g.selectAll('.dot')
    .data(filteredData)
    .enter().append('circle')
    .attr('class', 'dot')
    .attr('r', 3.5)
    .attr('cx', d => x(d.Horsepower))
    .attr('cy', d => y(d.Weight_in_lbs))
    .style('fill', '#fff');
}
```