Below are the first {{ .characters }} of a file. What can you tell me about this file? Please respond in JSON, with the following fields:
* `type`: one of `HTML`, `CSV`, `JSON`, or `other`
* `title`: a concise title for the file, like "Stock price history"
* `description`: a one-sentence description of the data, like "Contains hour-by-hour prices for different stock tickers, including bid, ask, and spread"
* `data_format`: a paragraph describing of the structure of the data, which could be used by a programmer who needs to analyze the data. Be sure to note if there are any special considerations or preprocessing that should be done to the data.
* `fields`: an array of strings indicating what fields are available for each data point. These must be the same names used in the data file provided, with the same capitalization and punctuation. If the data file does not provide names (e.g. a CSV with no headers), you can provide names like `column_1`. If there are nested JSON fields, use dot notation; for arrays, use `[]`.
* `suggested_visualizations`: A list of 3-5 suggested ways to visualize this data. Suggestions should be in imperative form, like "Draw a graph" or "Show the data".

If the document is unstructured text or HTML, look for any tables or structure that could potentially be extracted from the page.

An example response might be:
```
{
  "type": "JSON",
  "title": "Stock price history",
  "description": "Contains hour-by-hour prices for different stock tickers, including bid, ask, and spread",
  "data_format": "The data is a JSON array, where each entry in the array contains the fields `ticker` and `time`, as well as JSON objects for `price` and `company`. The `price` object has two subfields: `bid` and `ask`. The `company` object has two subfields: `sector` and `name`.",
  "fields": ["ticker", "time", "price.bid", "price.ask", "company.sector", "company.name", "tags[]"],
  "suggested_visualizations": [
    "Graph the asking price over time for each ticker",
    "Show the bid/ask spread over time, as a percentage of ask",
    "Display the average price by sector over time",
    "Draw a bar graph that shows average price by sector for the latest data"
  ]
}
```

Here's the data:

```
{{ .data }}
```

Remember, you're only seeing *part* of the data, maybe only part of one data point.
The structure of all data points are probably similar, but the data might be quite different.
