# Map of National Parks

## Code

```javascript
async function drawVisualization(svg, data) {
    // Parse the JSON data
    const jsonData = JSON.parse(data);

    // Validate and extract park locations from parsed data
    const parkLocations = jsonData.data.map(d => ({
        name: d.fullName,
        latitude: parseFloat(d.latitude),
        longitude: parseFloat(d.longitude)
    })).filter(d => !isNaN(d.latitude) && !isNaN(d.longitude));

    if (parkLocations.length === 0) {
        throw new Error("No valid park locations found in data.");
    }

    // Set up projection and path generator for map
    const projection = d3.geoMercator().fitSize([+svg.attr('width'), +svg.attr('height')], {type: "MultiPoint", coordinates: parkLocations.map(d => [d.longitude, d.latitude])});
    const pathGenerator = d3.geoPath().projection(projection);

    // Load and draw base map (USA)
    const usMap = await d3.json("https://cdn.jsdelivr.net/npm/us-atlas@3/states-10m.json");
    svg.append("path")
        .datum(topojson.feature(usMap, usMap.objects.states))
        .attr("fill", "#444")
        .attr("d", pathGenerator);

    // Add parks as points on the map
    svg.selectAll("circle")
        .data(parkLocations)
        .join("circle")
        .attr("cx", d => projection([d.longitude, d.latitude])[0])
        .attr("cy", d => projection([d.longitude, d.latitude])[1])
        .attr("r", 3)
        .attr("fill", d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, parkLocations.length)));   

    // Optionally, style the SVG to ensure it contrast against a dark background
    svg.style("background-color", "#222")
        .attr("stroke", "#fff")
        .attr("stroke-width", 1);
}
```