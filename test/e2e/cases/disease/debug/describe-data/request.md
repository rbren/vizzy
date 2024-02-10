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
﻿cod,cdc,google,guardian,nyt
alzheimers,0.036286747,0.025974356,0.004943649,0.012451976
cancer,0.278526958,0.37617788,0.14101401,0.200501124
automobile_crashes,0.056812203,0.100184937,0.026506877,0.023220094
diabetes,0.034887479,0.101264952,0.021648934,0.033706173
heart_disease,0.305395675,0.02700541,0.024504116,0.034294267
homicide,0.008158781,0.030665256,0.281759805,0.220748665
kidney_disease,0.022814805,0.009054294,0.001316163,0.002911404
lower_respiratory_disease,0.064496341,0.020960729,0.01690009,0.019741052
overdose,0.014048823,0.018564208,0.001552862,0.002380666
pnuemonia_and_influenza,0.027566931,0.044959903,0.039092684,0.042853976
stroke,0.052789362,0.059869777,0.0623558,0.055810968
suicide,0.017168692,0.123948746,0.132229656,0.103950693
terrorism,7.82E-05,0.06136955,0.304998882,0.306252471
```

Remember, you're only seeing *part* of the data, maybe only part of one data point.
The structure of all data points are probably similar, but the data might be quite different.
