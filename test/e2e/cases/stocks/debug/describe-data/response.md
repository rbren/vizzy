```json
{
  "type": "CSV",
  "title": "Historical Stock Prices",
  "description": "Monthly closing prices of Microsoft, Amazon, and IBM stocks from January 2000 to March 2010.",
  "data_format": "The data is in CSV format with three columns: `symbol`, `date`, and `price`. Each row represents the closing price of a stock (identified by `symbol`) for a given month (specified in `date` as 'Mon dd yyyy'). The `price` is a floating point number representing the stock's closing price at the end of that month. When analyzing this data, it may be necessary to convert the `date` strings into a date-time format that can be easily manipulated and to ensure the `price` values are treated as numerical data for calculations.",
  "fields": ["symbol", "date", "price"],
  "suggested_visualizations": [
    "Plot the monthly closing prices for each stock over time",
    "Compare the stock performance of Microsoft, Amazon, and IBM in a line graph",
    "Show the year-over-year performance comparison for each stock",
    "Visualize the stock volatility by plotting the month-to-month percentage change",
    "Create a bar chart showing the average annual price for each stock"
  ]
}
```