# Global Fertility Rate Map

```javascript
async function drawVisualization(svg, dataString) {
  const data = JSON.parse(dataString);

  const latestFertilityByCountry = data.reduce((acc, item) => {
    const { country, year, fertility } = item;
    if (!acc[country] || acc[country].year < year) {
      acc[country] = { year, fertility };
    }
    return acc;
  }, {});

  const sanitizedData = Object.entries(latestFertilityByCountry)
    .filter(([_, { fertility }]) => fertility != null)
    .map(([country, { fertility }]) => ({ country, fertility }));

  const mapData = await d3.json("https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json");

  const projection = d3.geoNaturalEarth1().fitSize([+svg.attr("width"), +svg.attr("height")], topojson.feature(mapData, mapData.objects.countries));
  const path = d3.geoPath().projection(projection);

  const colorScale = d3.scaleLinear()
    .domain([d3.min(sanitizedData.map(d => d.fertility)), d3.max(sanitizedData.map(d => d.fertility))])
    .range(["#800080", "#008000"]) 
    .interpolate(d3.interpolateSpectral);

  svg.selectAll("path")
    .data(topojson.feature(mapData, mapData.objects.countries).features)
    .join("path")
    .attr("fill", d => {
      const countryData = sanitizedData.find(c => c.country === d.properties.name);
      return countryData ? colorScale(countryData.fertility) : "#ccc";
    })
    .attr("d", path)
    .attr("stroke", "white")
    .style("stroke-width", 0.5)
    .attr("title", d => d.properties.name);

  svg.selectAll("text")
    .style("fill", "#fff")
    .style("font-family", "sans-serif");
}
```