# Global Fertility Rates

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
                         .range(["#800026", "#006837"]); // Deep purple to green scale
    
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