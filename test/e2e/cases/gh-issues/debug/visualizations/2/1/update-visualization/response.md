# GitHub Issues Reaction Summary

```javascript
async function drawVisualization(svg, data) {
    const issues = JSON.parse(data);

    let reactionData = issues.map(issue => ({
        title: issue.title,
        reactions: issue.reactions // This presumes reactions structure is parsed directly from issue object
    }))
    .filter(issue => Object.values(issue.reactions).reduce((a, b) => a + b, 0) > 0)
    .slice(0, 20); // Taking top 20 for visualization

    if(reactionData.length === 0) throw new Error("No valid data points found.");

    const margin = {top: 50, right: 30, bottom: 100, left: 60},
        width = parseInt(svg.attr('width')) - margin.left - margin.right,
        height = parseInt(svg.attr('height')) - margin.top - margin.bottom;

    const x = d3.scaleBand()
        .range([0, width])
        .domain(reactionData.map(d => d.title))
        .padding(0.1);

    const y = d3.scaleLinear()
        .domain([0, d3.max(reactionData, d => Object.values(d.reactions).reduce((a, b) => a + b, 0))])
        .range([height, 0]);

    const color = d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, reactionData.length));

    const svgGroup = svg.append("g")
        .attr("transform", `translate(${margin.left},${margin.top})`);

    svgGroup.selectAll(".bar")
        .data(reactionData)
        .enter().append("rect")
            .attr("class", "bar")
            .attr("x", d => x(d.title))
            .attr("width", x.bandwidth())
            .attr("y", d => y(Object.values(d.reactions).reduce((a, b) => a + b, 0)))
            .attr("height", d => height - y(Object.values(d.reactions).reduce((a, b) => a + b, 0)))
            .attr("fill", (d, i) => color(i));

    svgGroup.append("g")
        .attr("transform", `translate(0,${height})`)
        .call(d3.axisBottom(x))
        .selectAll("text")  
            .style("text-anchor", "end")
            .attr("dx", "-.8em")
            .attr("dy", ".15em")
            .attr("transform", "rotate(-65)");

    svgGroup.append("g")
        .call(d3.axisLeft(y));

    svg.selectAll("text").attr("fill", "#fff").attr("font-family", "sans-serif");
    svg.selectAll("path.domain, g.tick line").attr("stroke", "#fff");

    const tooltip = d3.select("body").append("div")
        .attr("class", "tooltip")
        .style("background", "#fff")
        .style("color", "#000")
        .style("padding", "10px")
        .style("border-radius", "5px")
        .style("font-family", "sans-serif")
        .style("opacity", 0);

    svgGroup.selectAll(".bar")
        .on("mouseover", function(event, d) {
            tooltip.transition()
                .duration(200)
                .style("opacity", .9);
            tooltip.html(d.title + "<br/>" + JSON.stringify(d.reactions))
                .style("left", (event.pageX) + "px")
                .style("top", (event.pageY - 28) + "px");
        })
        .on("mouseout", function() {
            tooltip.transition()
                .duration(500)
                .style("opacity", 0);
        });
}
```