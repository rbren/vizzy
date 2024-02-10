Below are the first 8000 of a file. What can you tell me about this file? Please respond in JSON, with the following fields:
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
[
    {
        "_comment": "Data courtesy of Gapminder.org",
        "year": 1955,
        "fertility": 7.7,
        "life_expect": 30.332,
        "n_fertility": 7.7,
        "n_life_expect": 31.997,
        "country": "Afghanistan"
    },
    {
        "year": 1960,
        "fertility": 7.7,
        "life_expect": 31.997,
        "p_fertility": 7.7,
        "n_fertility": 7.7,
        "p_life_expect": 30.332,
        "n_life_expect": 34.02,
        "country": "Afghanistan"
    },
    {
        "year": 1965,
        "fertility": 7.7,
        "life_expect": 34.02,
        "p_fertility": 7.7,
        "n_fertility": 7.7,
        "p_life_expect": 31.997,
        "n_life_expect": 36.088,
        "country": "Afghanistan"
    },
    {
        "year": 1970,
        "fertility": 7.7,
        "life_expect": 36.088,
        "p_fertility": 7.7,
        "n_fertility": 7.7,
        "p_life_expect": 34.02,
        "n_life_expect": 38.438,
        "country": "Afghanistan"
    },
    {
        "year": 1975,
        "fertility": 7.7,
        "life_expect": 38.438,
        "p_fertility": 7.7,
        "n_fertility": 7.8,
        "p_life_expect": 36.088,
        "n_life_expect": 39.854,
        "country": "Afghanistan"
    },
    {
        "year": 1980,
        "fertility": 7.8,
        "life_expect": 39.854,
        "p_fertility": 7.7,
        "n_fertility": 7.9,
        "p_life_expect": 38.438,
        "n_life_expect": 40.822,
        "country": "Afghanistan"
    },
    {
        "year": 1985,
        "fertility": 7.9,
        "life_expect": 40.822,
        "p_fertility": 7.8,
        "n_fertility": 8,
        "p_life_expect": 39.854,
        "n_life_expect": 41.674,
        "country": "Afghanistan"
    },
    {
        "year": 1990,
        "fertility": 8,
        "life_expect": 41.674,
        "p_fertility": 7.9,
        "n_fertility": 8,
        "p_life_expect": 40.822,
        "n_life_expect": 41.763,
        "country": "Afghanistan"
    },
    {
        "year": 1995,
        "fertility": 8,
        "life_expect": 41.763,
        "p_fertility": 8,
        "n_fertility": 7.4792,
        "p_life_expect": 41.674,
        "n_life_expect": 42.129,
        "country": "Afghanistan"
    },
    {
        "year": 2000,
        "fertility": 7.4792,
        "life_expect": 42.129,
        "p_fertility": 8,
        "p_life_expect": 41.763,
        "country": "Afghanistan"
    },
    {
        "year": 1955,
        "fertility": 3.1265,
        "life_expect": 64.399,
        "n_fertility": 3.0895,
        "n_life_expect": 65.142,
        "country": "Argentina"
    },
    {
        "year": 1960,
        "fertility": 3.0895,
        "life_expect": 65.142,
        "p_fertility": 3.1265,
        "n_fertility": 3.049,
        "p_life_expect": 64.399,
        "n_life_expect": 65.634,
        "country": "Argentina"
    },
    {
        "year": 1965,
        "fertility": 3.049,
        "life_expect": 65.634,
        "p_fertility": 3.0895,
        "n_fertility": 3.1455,
        "p_life_expect": 65.142,
        "n_life_expect": 67.065,
        "country": "Argentina"
    },
    {
        "year": 1970,
        "fertility": 3.1455,
        "life_expect": 67.065,
        "p_fertility": 3.049,
        "n_fertility": 3.44,
        "p_life_expect": 65.634,
        "n_life_expect": 68.481,
        "country": "Argentina"
    },
    {
        "year": 1975,
        "fertility": 3.44,
        "life_expect": 68.481,
        "p_fertility": 3.1455,
        "n_fertility": 3.15,
        "p_life_expect": 67.065,
        "n_life_expect": 69.942,
        "country": "Argentina"
    },
    {
        "year": 1980,
        "fertility": 3.15,
        "life_expect": 69.942,
        "p_fertility": 3.44,
        "n_fertility": 3.053,
        "p_life_expect": 68.481,
        "n_life_expect": 70.774,
        "country": "Argentina"
    },
    {
        "year": 1985,
        "fertility": 3.053,
        "life_expect": 70.774,
        "p_fertility": 3.15,
        "n_fertility": 2.9,
        "p_life_expect": 69.942,
        "n_life_expect": 71.868,
        "country": "Argentina"
    },
    {
        "year": 1990,
        "fertility": 2.9,
        "life_expect": 71.868,
        "p_fertility": 3.053,
        "n_fertility": 2.63,
        "p_life_expect": 70.774,
        "n_life_expect": 73.275,
        "country": "Argentina"
    },
    {
        "year": 1995,
        "fertility": 2.63,
        "life_expect": 73.275,
        "p_fertility": 2.9,
        "n_fertility": 2.35,
        "p_life_expect": 71.868,
        "n_life_expect": 74.34,
        "country": "Argentina"
    },
    {
        "year": 2000,
        "fertility": 2.35,
        "life_expect": 74.34,
        "p_fertility": 2.63,
        "p_life_expect": 73.275,
        "country": "Argentina"
    },
    {
        "year": 1955,
        "fertility": 5.15,
        "life_expect": 64.381,
        "n_fertility": 4.399,
        "n_life_expect": 66.606,
        "country": "Aruba"
    },
    {
        "year": 1960,
        "fertility": 4.399,
        "life_expect": 66.606,
        "p_fertility": 5.15,
        "n_fertility": 3.301,
        "p_life_expect": 64.381,
        "n_life_expect": 68.336,
        "country": "Aruba"
    },
    {
        "year": 1965,
        "fertility": 3.301,
        "life_expect": 68.336,
        "p_fertility": 4.399,
        "n_fertility": 2.651,
        "p_life_expect": 66.606,
        "n_life_expect": 70.941,
        "country": "Aruba"
    },
    {
        "year": 1970,
        "fertility": 2.651,
        "life_expect": 70.941,
        "p_fertility": 3.301,
        "n_fertility": 2.45,
        "p_life_expect": 68.336,
        "n_life_expect": 71.83,
        "country": "Aruba"
    },
    {
        "year": 1975,
        "fertility": 2.45,
        "life_expect": 71.83,
        "p_fertility": 2.651,
        "n_fertility": 2.358,
        "p_life_expect": 70.941,
        "n_life_expect": 74.116,
        "country": "Aruba"
    },
    {
        "year": 1980,
        "fertility": 2.358,
        "life_expect": 74.116,
        "p_fertility": 2.45,
        "n_fertility": 2.3,
        "p_life_expect": 71.83,
        "n_life_expect": 74.494,
        "country": "Aruba"
    },
    {
        "year": 1985,
        "fertility": 2.3,
        "life_expect": 74.494,
        "p_fertility": 2.358,
        "n_fertility": 2.28,
        "p_life_expect": 74.116,
        "n_life_expect": 74.108,
        "country": "Aruba"
    },
    {
        "year": 1990,
        "fertility": 2.28,
        "life_expect": 74.108,
        "p_fertility": 2.3,
        "n_fertility": 2.208,
        "p_life_expect": 74.494,
        "n_life_expect": 73.011,
        "country": "Aruba"
    },
    {
        "year": 1995,
        "fertility": 2.208,
        "life_expect": 73.011,
        "p_fertility": 2.28,
        "n_fertility": 2.124,
        "p_life_expect": 74.108,
        "n_life_expect": 73.451,
        "country": "Aruba"
    },
    {
        "year": 2000,
        "fertility": 2.124,
        "life_expect": 73.451,
        "p_fertility": 2.208,
        "p_life_expect": 73.011,
        "country": "Aruba"
    },
    {
        "year": 1955,
        "fertility": 3.406,
        "life_expect": 70.33,
        "n_fertility": 3.274,
        "n_life_expect": 70.93,
        "country": "Australia"
    },
    {
        "year": 1960,
        "fertility": 3.274,
        "life_expect": 70.93,
        "p_fertility": 3.406,
        "n_fertility": 2.871,
        "p_life_expect": 70.33,
        "n_life_expect": 71.1,
        "country": "Australia"
    },
    {
        "year": 1965,
        "fertility": 2.871,
        "life_expect": 71.1,
        "p_fertility": 3.274,
        "n_fertility": 2.535,
        "p_life_expect": 70.93,
        "n_life_expect": 71.93,
        "country": "Australia"
    },
    {
        "year": 1970,
        "fertility": 2.535,
        "life_expect": 71.93,
        "p_fertility": 2.871,
        "n_fertility": 1.989,
        "p_life_expect": 71.1,
        "n_life_expect": 73.49,
        "country
```

Remember, you're only seeing *part* of the data, maybe only part of one data point.
The structure of all data points are probably similar, but the data might be quite different.
