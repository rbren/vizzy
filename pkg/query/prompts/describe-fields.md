# Data Analysis Task

## Task
Please write some JavaScript code that computes important metadata about each primitive field,
(i.e. strings, numerics, and booleans). Metadata might include things
like maximum, minimum, average, mode, enum, etc. For string fields, you can compute
things like the most common entry, or the number of unique entries.
You can get creative with what metadata is computed.

The code should provide a function called `computeMetadata(data)`,
which takes in a string containing the file's data. The data must be parsed, e.g.
using `JSON.parse`, `d3.csvParse`, or a custom parsing function if it's in a different format.
The function must output a JavaScript object.
Each key in the object must be one of the fields listed above,
and each value must be an object contaning metadata about that field.

Each metadata object should contain a `sampleValue` pulled from the data. It should have
the same type (i.e. string, numeric, or boolean) as the data.
If a field contains less than 10 unique string values,
include an `enum` in the metadata, which contains an array those values.

Every other metadata field must be a number, string, or boolean. Metadata fields must not be objects
or arrays.

There should be an additional field, `$dataPoints`, which contains the total number of data points
in the set.

For example, if we had two fields "ticker" and "price", the function might return an object like:
```
{
  "$dataPoints": 230,
  "ticker": {
    "sampleValue": "MSFT",
    "mostCommonEntry": "AAPL",
  },
  "price": {
    "sampleValue": 42.0,
    "average": 53.25,
    "maximum": 100,
    "minimum": 2.34
  }
}
```

And the function might look like this:
```
function computeMetadata(data) {
  const parsed = JSON.parse(data);
  let mostCommonTicker = _.head(_(parsed.map(d => d.ticker))
    .countBy()
    .entries()
    .maxBy(_.last));
  const uniqueTickers = _.uniq(parsed.map(d => d.ticker));

  const prices = parsed.map(d => d.price);

  let metadata = {
    "$dataPoints": parsed.length,
    "ticker": {
      "sampleValue": parsed[0].ticker,
      "mostCommonEntry": mostCommonTicker,
    },
    "price": {
      "sampleValue": parsed[0].price,
      "maximum": _.max(prices),
      "minimum": _.min(prices),
      "average": _.mean(prices),
    },
  }
  if (uniqueTickers.length < 10) metadata.ticker.enum = uniqueTickers;
}
```

The code must take into account the possibility of null or missing values. Make sure
to add guards in case a field is undefined, an array is empty, etc.

## Dataset Description
{{template "data_description" .data}}

## Technical Details
The following libraries are available:
* D3 v7, as `d3`
* lodash v4, as `_`

No other libraries are available. You should only use JavaScript functionality that is
available in most modern browsers.

Lodash is already available as `window._`. You should not use `import`, `require` or make any other
attempts to import third-party libraries. Your code should _only_ define the `computeMetadata`
function.
