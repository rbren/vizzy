# Failure Patterns
This is a document of common failure patterns to work on. If your visualization is having issues, open a PR
to add them to this document. Look for an existing section that might apply and add a link or details if
you think it'd be helpful. If there are no existing sections that match your issue, start a new one.

Patterns are roughly in order of how unsolved they are.

## Bad JavaScript usage

Example 1: calling `new Date()` with a year instead of milliseconds
```js
years[index] = new Date(splitYr[0] - 1);
```

Example 2: using an undefined but common variable, like `height`

#### Mitigation
These are often one-off issues. If we see the same one again and again, we can
add special instructions.

#### Workarounds
* Usually asking Vizzy to fix the issue helps, if you can spot it in the code.

## Semantic Misunderstanding
Sometimes the LLM doesn't understand the semantic meaning of a field. E.g. `rank` is interpreted
as "higher is better", when in reality, "lower is better". If the field is actually labeled `rank`
this might be inferred, but in the Spotify data it's implied.

#### Mitigation
* [ ] Add field descriptions/interpretations to the initial data description, including the
significance of high/low/null values

#### Workarounds
* Manually edit the data description to specify the semantic meaning of ambiguous fields

## String Matching
If the user asks for a graph of Google's stock, and the data has `ticker=GOOG`, the LLM
might successfully make that jump from world knowledge--it knows its looking for a ticker,
and that Google's ticker is `GOOG`.

However, sometimes the strings in the data are not as easy to match. E.g. if I ask for a graph of
"rock and roll", and the data has `genre=rock`, there's no way for the LLM to canonicalize that
string. It would have to implement some kind of fuzzy matching.

#### Workarounds
* Hard code string maps into the JavaScript
* Ask Vizzy to hard code the maps using its world knowledge
  * This has proven hard to do...

## Huge amounts of data
The dataset might provide enough samples for 10k bars in a bar chart. Often the LLM tries to
draw all of them.

#### Mitigation
* [x] Encourage Vizzy to limit the number of bars/lines/etc shown

#### Workarounds
* Specifically tell Vizzy to only show N data points

## Updates try to build on top of bad code
Sometimes instead of editing the code, the LLM tries to _add_ new code. E.g. check out this
code where the user asked for zero-data to be colored white:

```
async function drawVisualization(svg, data) {
    /* ... */
    // Draw map
    svg.selectAll('path')
      .data(filteredCountries)
      .join('path')
      .attr('d', path)
      .attr('fill', d => colorScale(d.properties.winners))

    // Color countries with no winners to be white
    svg.selectAll('path')
        .data(countries)
        .join('path')
        .attr('d', path)
        .attr('fill', d => {
            if (d.properties.winners > 0) {
                return colorScale(d.properties.winners);
            } else {
                return "#ffffff";
            }
        });
}
```

It ends up drawing _two_ maps, which causes a weird bug when tooltips are added.

## Positioning
The LLM often gets the inital position wrong, and struggles to take feedback.

Sometimes elements are drawn in the wrong order, causing e.g. text to be hidden behind a graph

#### Mitigation
* [ ] Experiment with sending an image of the visualization back to ChatGPT
* [x] Add prompt instructions about drawing text last

#### Workarounds
* Redo the visualization from scratch, rather than trying to get Vizzy to fix it
* Fix it manually

## Invalid SVG attributes
Often we end up with invalid attributes for SVG elements that are generated using the data, e.g.
```
Error: Invalid value for attribute d="M1.2121212121212122,NaNL1.4141414141414141,NaNC1.6161616161616161,NaN,2.0202020202020203,NaN,2.4242424...
```
These errors are very hard to catch and bubble up.

#### Mitigation
* [x] We've added code that checks for specific issues using CSS selectors, like `path[d*="NaN"]`
* [x] We've added prompt instructions to always sanitize data before passing it to d3's `data()` function

## Using third-party libraries
Often the LLM tries to use third-party libraries that haven't been included

#### Mitigation
* [x] We've selectively added common and useful libraries. We've also explicitly added the list
of libraries to the prompt, and said others can't be used

## Using deprecated d3 functionality
Sometimes the LLM tries to use functions that were removed in v6.

#### Mitigation
* [x] We've added v6 migration instructions to the prompt, though this is somewhat expensive.
* [x] We string-replace function calls like `d3.keys` -> `Object.keys`


