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
﻿FID,NAME,Obesity,SHAPE_Length,SHAPE_Area
1,Texas,32.4,15408321.8698148,7672329221262.61
2,California,24.2,14518698.4566722,5327809415403.01
3,Kentucky,34.6,6346698.58704225,1128830286256.75
4,Georgia,30.7,5795595.91737164,1652980281864.84
5,Wisconsin,30.7,6806782.08358127,1567816094351.9
6,Oregon,30.1,7976010.9415358,3178445626071.46
7,Virginia,29.2,7710803.89869229,1158804356933.99
8,Tennessee,33.8,6350376.86082221,1177054472558.46
9,Louisiana,36.2,7383856.91994343,1355093764697.92
10,New York,25,7981383.07788145,1411440550734.63
11,Michigan,31.2,12407988.0244766,1657221839418.83
12,Idaho,28.6,9081125.22209832,2593598268244.12
13,Florida,26.8,10533929.1444184,1674411223041.2
14,Alaska,29.8,6663726.59054098,530113882116.939
15,Montana,23.6,9745471.75354953,4371466366614.52
16,Minnesota,26.1,8530250.42958512,2367874733333.85
17,Nebraska,31.4,6921599.98880176,2207444436435.34
18,Washington,26.4,7508966.8395569,2214214821535.73
19,Ohio,29.8,5015484.92538504,1158015967734.46
20,Illinois,30.8,6172530.80822631,1571568299027.73
21,Missouri,32.4,7165632.96877297,1949020823655.51
22,Iowa,32.1,5558144.62470353,1575588858404.13
23,South Dakota,30.4,6680911.24820537,2199053487708.16
24,Arkansas,34.5,5707634.00276971,1488699273848.93
25,Mississippi,35.6,5834201.76787132,1327853198175.37
26,Colorado,20.2,7092296.49769191,3066878609047.81
27,North Carolina,30.1,6714056.23757219,1495755986090.79
28,Utah,24.5,6798972.50395198,2614280967942.81
29,Oklahoma,33.9,7857119.57078742,1982448890144.3
30,Wyoming,29,6860971.19013795,2897400399041.46
31,West Virginia,35.6,5374279.69717876,685167361615.589
32,Indiana,31.3,4858882.21795382,1010364222333.8
33,Massachusetts,24.3,4197218.1811822,247224998635.541
34,Nevada,26.7,8240794.81048687,3564354191161.93
35,Connecticut,25.3,1780630.95573379,145421898735.938
36,District of Columbia,22.1,200324.960720334,1972102312.29117
37,Rhode Island,26,984524.581180468,36178076477.7396
38,Alabama,35.6,5750657.81244629,1442806615250.92
39,Puerto Rico,29.5,2136273.27059009,114783882870.193
40,South Carolina,31.7,4370940.54750344,878270407919.019
41,Maine,30,5560035.36660915,990384869184.921
42,Hawaii,22.7,0,0
43,Arizona,28.4,8044184.23304366,3562685913793.35
44,New Mexico,28.8,8075166.55513897,3622933199179.54
45,Maryland,28.9,5850363.29291603,303943203091.598
46,Delaware,29.7,1383604.11753278,59081101616.5539
47,Pennsylvania,30,5024348.11219534,1288451936889.14
48,Kansas,34.2,6540498.4197016,2340365916439.03
49,Vermont,25.1,2653732.23226363,278931330216.061
50,New Jersey,25.6,2599119.26086415,224606547518.286
51,North Dakota,31,5872756.35857996,2013151767399.38
52,New Hampshire,26.3,2674767.14806674,270529403259.303

```

Remember, you're only seeing *part* of the data, maybe only part of one data point.
The structure of all data points are probably similar, but the data might be quite different.
