# Map of National Parks

## Plan
1. Parse the input data string using `JSON.parse`.
2. Extract the `data` array from the parsed object. This array contains objects for each national park, including their latitude and longitude coordinates.
3. Filter and sanitize the data to ensure that all necessary fields (`latitude`, `longitude`, and `fullName`) exist and are valid. Null or missing values will be removed.
4. Since the visualization is a map of all the parks, pre-select a reasonable number of parks for visualization if the dataset is large, focusing on unique latitudes and longitudes to ensure a diverse representation across the map.
5. Use a world map from the provided TopoJSON datasets to serve as the base of our map.
6. Plot the parks on the map using their latitude and longitude coordinates. Each park will be represented by a circle.
7. Color the circles using a scale based on the total number of activities in each park, applying the `interpolateSpectral` color scheme.
8. Ensure that all text is white (#fff) to be visible against the dark background.

## Code
```javascript
async function drawVisualization(svg, data) {
  // Parse the data from a JSON string
  const rawData = JSON.parse(data);
  const parksData = rawData.data.filter(d => d.latitude && d.longitude && d.fullName);

  // Load world map geography
  const worldData = await d3.json('https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json');
  
  // Get SVG dimensions
  const width = parseInt(svg.attr('width')), height = parseInt(svg.attr('height'));

  // Create a map projection and path generator
  const projection = d3.geoMercator().fitSize([width, height], topojson.feature(worldData, worldData.objects.countries));
  const path = d3.geoPath().projection(projection);

  // Draw the world map
  svg.append("g")
    .selectAll("path")
    .data(topojson.feature(worldData, worldData.objects.countries).features)
    .enter().append("path")
      .attr("fill", "#bbb")
      .attr("d", path)
      .attr("stroke", "#fff");
  
  // Scale for circle sizes, based on the number of activities
  const sizeScale = d3.scaleSqrt()
                       .domain([0, d3.max(parksData, d => d.activities.length)])
                       .range([1, 10]);

  // Convert latitude and longitude to x and y coordinates using the projection
  parksData.forEach(park => {
    const coords = projection([parseFloat(park.longitude), parseFloat(park.latitude)]);
    park.x = coords[0];
    park.y = coords[1];
  });

  // Draw circles for each national park
  svg.append("g")
    .selectAll("circle")
    .data(parksData)
    .enter().append("circle")
      .attr("cx", d => d.x)
      .attr("cy", d => d.y)
      .attr("r", d => sizeScale(d.activities.length))
      .attr("fill", d => d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, parksData.length))(d.activities.length))
      .attr("stroke", "#fff");

  // Ensure text is white and readable against the dark background
  svg.style("color", "#fff")
     .style("font-family", "sans-serif");
}
```