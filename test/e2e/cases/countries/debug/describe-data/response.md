```json
{
  "type": "JSON",
  "title": "Global Fertility and Life Expectancy",
  "description": "Tracks changes in fertility rates and life expectancy across different countries over years",
  "data_format": "The data is structured as a JSON array with each object representing a data point for a country in a given year. Each object contains the fields `year`, `fertility`, `life_expect`, and `country`. Optionally, objects may contain previous (`p_`) and next (`n_`) values for both `fertility` and `life_expect`, indicating trends over time. When reading this data, special attention should be given to these predictive and historical values to understand temporal trends. Preprocessing may involve filtering by country or time range, and handling missing `p_` and `n_` fields when they are absent.",
  "fields": ["year", "fertility", "life_expect", "n_fertility", "n_life_expect", "p_fertility", "p_life_expect", "country"],
  "suggested_visualizations": [
    "Plot fertility rates over time for selected countries",
    "Graph life expectancy changes over time for selected countries",
    "Compare the fertility rates to life expectancy in a scatter plot",
    "Create a heatmap showing changes in life expectancy across all countries",
    "Generate line graphs comparing the previous and next fertility and life expectancy trends for a country"
  ]
}
```