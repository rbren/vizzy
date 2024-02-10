# GitHub Issues Reaction Summary

```javascript
async function drawVisualization(svg, data) {
    const width = parseInt(svg.attr('width'));
    const height = parseInt(svg.attr('height'));
    const margin = { top: 20, right: 20, bottom: 100, left: 60 };

    // Parsing and transforming data
    let jsonData = JSON.parse(data);
    jsonData = jsonData.filter(d => d.reactions && Object.keys(d.reactions).length > 0);

    // Summarizing reactions for each issue
    const reactionsSummary = jsonData.map(d => ({
        title: d.title,
        reactions: d.reactions,
        totalReactionsCount: Object.values(d.reactions).reduce((acc, curr) => acc + curr, 0)
    }));

    // Scale for the x-axis
    const xScale = d3.scaleBand()
        .range([margin.left, width - margin.right])
        .padding(0.1)
        .domain(reactionsSummary.map(d => d.title));

    // Scale for the y-axis
    const yScale = d3.scaleLinear()
        .domain([0, d3.max(reactionsSummary, d => d.totalReactionsCount)])
        .nice()
        .range([height - margin.bottom, margin.top]);

    // Adding the X-axis
    svg.append("g")
        .attr("transform", `translate(0,${height - margin.bottom})`)
        .call(d3.axisBottom(xScale))
        .selectAll("text")
        .attr("transform", "translate(-10,0)rotate(-45)")
        .style("text-anchor", "end")
        .style("fill", "#fff");

    // Adding the Y-axis
    svg.append("g")
        .attr("transform", `translate(${margin.left}, 0)`)
        .call(d3.axisLeft(yScale))
        .style("color", "#fff");

    // Bars
    svg.append("g")
        .selectAll("rect")
        .data(reactionsSummary)
        .join("rect")
        .attr("x", d => xScale(d.title))
        .attr("y", d => yScale(d.totalReactionsCount))
        .attr("width", xScale.bandwidth())
        .attr("height", d => height - margin.bottom - yScale(d.totalReactionsCount))
        .attr("fill", d => d3.scaleOrdinal(d3.quantize(d3.interpolateSpectral, reactionsSummary.length))(d.title));

    // Tooltip
    const tooltip = d3.select("body").append("div")
        .attr("class", "tooltip")
        .style("position", "absolute")
        .style("background-color", "#fff")
        .style("border", "solid")
        .style("border-width", "2px")
        .style("border-radius", "5px")
        .style("padding", "10px")
        .style("visibility", "hidden");

    svg.selectAll("rect")
        .on("mouseover", (event, d) => {
            let content = `<strong>${d.title}</strong><br>`;
            for (const [reaction, count] of Object.entries(d.reactions)) {
                content += `${reaction}: ${count}<br>`;
            }
            tooltip.html(content)
                .style("visibility", "visible");
        })
        .on("mousemove", (event) => {
            tooltip.style("top", (event.pageY - 10) + "px")
                .style("left", (event.pageX + 10) + "px");
        })
        .on("mouseout", () => {
            tooltip.style("visibility", "hidden");
        });
}
```