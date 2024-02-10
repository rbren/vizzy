```json
{
  "type": "JSON",
  "title": "Global Fertility and Life Expectancy",
  "description": "Tracks changes in fertility rates and life expectancy across different countries over years",
  "data_format": "The data is structured as a JSON array, where each object represents a data point for a specific year in a specific country. Each object contains several key-value pairs: 'year' indicates the year of the data, 'fertility' and 'p_fertility' (previous fertility), 'n_fertility' (next fertility), 'life_expect', 'p_life_expect' (previous life expectancy), and 'n_life_expect' (next life expectancy), with 'country' specifying the country. Some objects may not have all these fields, indicating the start or end of the available data for a country. Pay attention to possible missing data when aggregating or comparing across years or countries. The 'p_' prefix indicates the value from the previous data point, and 'n_' prefix indicates the value from the next data point, which might require interpolation or estimation for analysis.",
  "fields": [
    "year",
    "fertility",
    "life_expect",
    "n_fertility",
    "n_life_expect",
    "country",
    "p_fertility",
    "p_life_expect"
  ],
  "suggested_visualizations": [
    "Plot fertility rate changes over time for selected countries.",
    "Graph life expectancy over time for selected countries.",
    "Compare the fertility rate versus life expectancy in a scatter plot for various years.",
    "Create a heatmap for fertility rates or life expectancy across different countries over time.",
    "Animate changes in fertility rates and life expectancy over time on a global map."
  ]
}
```