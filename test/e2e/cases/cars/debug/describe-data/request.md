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
      "Name":"chevrolet chevelle malibu",
      "Miles_per_Gallon":18,
      "Cylinders":8,
      "Displacement":307,
      "Horsepower":130,
      "Weight_in_lbs":3504,
      "Acceleration":12,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"buick skylark 320",
      "Miles_per_Gallon":15,
      "Cylinders":8,
      "Displacement":350,
      "Horsepower":165,
      "Weight_in_lbs":3693,
      "Acceleration":11.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"plymouth satellite",
      "Miles_per_Gallon":18,
      "Cylinders":8,
      "Displacement":318,
      "Horsepower":150,
      "Weight_in_lbs":3436,
      "Acceleration":11,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"amc rebel sst",
      "Miles_per_Gallon":16,
      "Cylinders":8,
      "Displacement":304,
      "Horsepower":150,
      "Weight_in_lbs":3433,
      "Acceleration":12,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"ford torino",
      "Miles_per_Gallon":17,
      "Cylinders":8,
      "Displacement":302,
      "Horsepower":140,
      "Weight_in_lbs":3449,
      "Acceleration":10.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"ford galaxie 500",
      "Miles_per_Gallon":15,
      "Cylinders":8,
      "Displacement":429,
      "Horsepower":198,
      "Weight_in_lbs":4341,
      "Acceleration":10,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"chevrolet impala",
      "Miles_per_Gallon":14,
      "Cylinders":8,
      "Displacement":454,
      "Horsepower":220,
      "Weight_in_lbs":4354,
      "Acceleration":9,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"plymouth fury iii",
      "Miles_per_Gallon":14,
      "Cylinders":8,
      "Displacement":440,
      "Horsepower":215,
      "Weight_in_lbs":4312,
      "Acceleration":8.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"pontiac catalina",
      "Miles_per_Gallon":14,
      "Cylinders":8,
      "Displacement":455,
      "Horsepower":225,
      "Weight_in_lbs":4425,
      "Acceleration":10,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"amc ambassador dpl",
      "Miles_per_Gallon":15,
      "Cylinders":8,
      "Displacement":390,
      "Horsepower":190,
      "Weight_in_lbs":3850,
      "Acceleration":8.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"citroen ds-21 pallas",
      "Miles_per_Gallon":null,
      "Cylinders":4,
      "Displacement":133,
      "Horsepower":115,
      "Weight_in_lbs":3090,
      "Acceleration":17.5,
      "Year":"1970-01-01",
      "Origin":"Europe"
   },
   {
      "Name":"chevrolet chevelle concours (sw)",
      "Miles_per_Gallon":null,
      "Cylinders":8,
      "Displacement":350,
      "Horsepower":165,
      "Weight_in_lbs":4142,
      "Acceleration":11.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"ford torino (sw)",
      "Miles_per_Gallon":null,
      "Cylinders":8,
      "Displacement":351,
      "Horsepower":153,
      "Weight_in_lbs":4034,
      "Acceleration":11,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"plymouth satellite (sw)",
      "Miles_per_Gallon":null,
      "Cylinders":8,
      "Displacement":383,
      "Horsepower":175,
      "Weight_in_lbs":4166,
      "Acceleration":10.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"amc rebel sst (sw)",
      "Miles_per_Gallon":null,
      "Cylinders":8,
      "Displacement":360,
      "Horsepower":175,
      "Weight_in_lbs":3850,
      "Acceleration":11,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"dodge challenger se",
      "Miles_per_Gallon":15,
      "Cylinders":8,
      "Displacement":383,
      "Horsepower":170,
      "Weight_in_lbs":3563,
      "Acceleration":10,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"plymouth 'cuda 340",
      "Miles_per_Gallon":14,
      "Cylinders":8,
      "Displacement":340,
      "Horsepower":160,
      "Weight_in_lbs":3609,
      "Acceleration":8,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"ford mustang boss 302",
      "Miles_per_Gallon":null,
      "Cylinders":8,
      "Displacement":302,
      "Horsepower":140,
      "Weight_in_lbs":3353,
      "Acceleration":8,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"chevrolet monte carlo",
      "Miles_per_Gallon":15,
      "Cylinders":8,
      "Displacement":400,
      "Horsepower":150,
      "Weight_in_lbs":3761,
      "Acceleration":9.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"buick estate wagon (sw)",
      "Miles_per_Gallon":14,
      "Cylinders":8,
      "Displacement":455,
      "Horsepower":225,
      "Weight_in_lbs":3086,
      "Acceleration":10,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"toyota corona mark ii",
      "Miles_per_Gallon":24,
      "Cylinders":4,
      "Displacement":113,
      "Horsepower":95,
      "Weight_in_lbs":2372,
      "Acceleration":15,
      "Year":"1970-01-01",
      "Origin":"Japan"
   },
   {
      "Name":"plymouth duster",
      "Miles_per_Gallon":22,
      "Cylinders":6,
      "Displacement":198,
      "Horsepower":95,
      "Weight_in_lbs":2833,
      "Acceleration":15.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"amc hornet",
      "Miles_per_Gallon":18,
      "Cylinders":6,
      "Displacement":199,
      "Horsepower":97,
      "Weight_in_lbs":2774,
      "Acceleration":15.5,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"ford maverick",
      "Miles_per_Gallon":21,
      "Cylinders":6,
      "Displacement":200,
      "Horsepower":85,
      "Weight_in_lbs":2587,
      "Acceleration":16,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"datsun pl510",
      "Miles_per_Gallon":27,
      "Cylinders":4,
      "Displacement":97,
      "Horsepower":88,
      "Weight_in_lbs":2130,
      "Acceleration":14.5,
      "Year":"1970-01-01",
      "Origin":"Japan"
   },
   {
      "Name":"volkswagen 1131 deluxe sedan",
      "Miles_per_Gallon":26,
      "Cylinders":4,
      "Displacement":97,
      "Horsepower":46,
      "Weight_in_lbs":1835,
      "Acceleration":20.5,
      "Year":"1970-01-01",
      "Origin":"Europe"
   },
   {
      "Name":"peugeot 504",
      "Miles_per_Gallon":25,
      "Cylinders":4,
      "Displacement":110,
      "Horsepower":87,
      "Weight_in_lbs":2672,
      "Acceleration":17.5,
      "Year":"1970-01-01",
      "Origin":"Europe"
   },
   {
      "Name":"audi 100 ls",
      "Miles_per_Gallon":24,
      "Cylinders":4,
      "Displacement":107,
      "Horsepower":90,
      "Weight_in_lbs":2430,
      "Acceleration":14.5,
      "Year":"1970-01-01",
      "Origin":"Europe"
   },
   {
      "Name":"saab 99e",
      "Miles_per_Gallon":25,
      "Cylinders":4,
      "Displacement":104,
      "Horsepower":95,
      "Weight_in_lbs":2375,
      "Acceleration":17.5,
      "Year":"1970-01-01",
      "Origin":"Europe"
   },
   {
      "Name":"bmw 2002",
      "Miles_per_Gallon":26,
      "Cylinders":4,
      "Displacement":121,
      "Horsepower":113,
      "Weight_in_lbs":2234,
      "Acceleration":12.5,
      "Year":"1970-01-01",
      "Origin":"Europe"
   },
   {
      "Name":"amc gremlin",
      "Miles_per_Gallon":21,
      "Cylinders":6,
      "Displacement":199,
      "Horsepower":90,
      "Weight_in_lbs":2648,
      "Acceleration":15,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"ford f250",
      "Miles_per_Gallon":10,
      "Cylinders":8,
      "Displacement":360,
      "Horsepower":215,
      "Weight_in_lbs":4615,
      "Acceleration":14,
      "Year":"1970-01-01",
      "Origin":"USA"
   },
   {
      "Name":"chevy c20",
      "Miles_per_Gallon":10,
      "Cylinders":8,
      
```

Remember, you're only seeing *part* of the data, maybe only part of one data point.
The structure of all data points are probably similar, but the data might be quite different.
