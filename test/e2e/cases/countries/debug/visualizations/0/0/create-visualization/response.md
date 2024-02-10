# Global Fertility Rates

## Plan

1. **Parse the JSON Data**: Use `JSON.parse()` to convert the given JSON string into JavaScript objects.
```javascript
const jsonData = JSON.parse(data);
```

2. **Transform the Data for Easier Use**: Transform the data into a format that groups data by country. This will make it easier to calculate the latest fertility rate per country.
```javascript
const groupedData = _.groupBy(jsonData, 'country');
```

3. **Sanitize and Aggregate the Data**: Iterate over each country, selecting the most recent year's data that has a fertility rate, and removing any entries without a fertility rate.
```javascript
const latestFertilityByCountry = _.mapValues(groupedData, (entries) => {
    const sortedEntries = _.sortBy(entries, 'year').reverse();
    return sortedEntries.find(entry => entry.fertility != null);
});
```

4. **Remove Null or Missing Values**: Ensure all data points have a valid fertility rate.
```javascript
const validData = _.omitBy(latestFertilityByCountry, value => value == null);
```

5. **Set Up D3 Projection and Path**: Set up a GeoMercator projection and path for the countries.
```javascript
const projection = d3.geoMercator().fitSize([width, height], topojson.feature(world, world.objects.countries));
const path = d3.geoPath().projection(projection);
```

6. **Draw the Map and Color**: Download a TopoJSON for global countries. Draw each country and color it based on its fertility rate.
```javascript
const colorScale = d3.scaleLinear(d3.quantize(d3.interpolateSpectral, validData.length));
```

7. **Filtering Data**: If necessary, filter the data to ensure the visualization is not overloaded. This step may involve selecting only the top 50 countries based on fertility rate, but given the visualization, it might be more useful to map all available data.

## Code

```javascript
async function drawVisualization(svg, data) {
    const width = +svg.attr('width');
    const height = +svg.attr('height');
    
    const jsonData = JSON.parse(data);
    const groupedData = _.groupBy(jsonData, 'country');
    const latestFertilityByCountry = _.mapValues(groupedData, entries => {
        const sortedEntries = _.sortBy(entries, 'year').reverse();
        return sortedEntries.find(entry => entry.fertility != null);
    });
    const validData = _.omitBy(latestFertilityByCountry, value => value == null);

    const world = await d3.json("https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json");
    const countries = topojson.feature(world, world.objects.countries);

    const projection = d3.geoMercator().fitSize([width, height], countries);
    const path = d3.geoPath().projection(projection);

    const fertilityRates = Object.values(validData).map(d => d.fertility);
    const colorScale = d3.scaleLinear()
                         .domain([_.min(fertilityRates), _.max(fertilityRates)])
                         .range(["#FFFFCC", "#800026"]);
    
    svg.selectAll("path")
        .data(countries.features)
        .join("path")
        .attr("fill", d => {
            const countryData = validData[d.properties.name];
            return countryData ? colorScale(countryData.fertility) : "#ccc";
        })
        .attr("d", path)
        .attr("stroke", "#fff")
        .attr("stroke-width", 0.5);
    
    svg.selectAll("text")
        .data(countries.features)
        .enter().append("text")
        .attr("class", "country-name")
        .attr("transform", d => `translate(${path.centroid(d)})`)
        .attr("text-anchor", "middle")
        .attr("alignment-baseline", "central")
        .attr("fill", "#fff")
        .attr("font-size", "8px")
        .attr("style", "pointer-events:none;")
        .text(d => d.properties.name);
}
```