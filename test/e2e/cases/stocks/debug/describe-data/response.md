```json
{
  "type": "CSV",
  "title": "Historical Stock Prices",
  "description": "Monthly closing prices of Microsoft, Amazon, and IBM stocks from January 2000 to March 2010.",
  "data_format": "The data is in CSV format, where each row represents a stock's monthly closing price at the end of the given month. Each row contains three fields: the stock's symbol (`symbol`), the date (`date`), and the price (`price`) at which the stock closed at the end of that month. The `date` field is formatted as `Mon d yyyy`, where `Mon` is the three-letter abbreviation of the month. Before analysis, you may need to convert the `date` field into a date object and the `price` field into a numeric type.",
  "fields": ["symbol", "date", "price"],
  "suggested_visualizations": [
    "Plot the monthly closing price of each stock over time.",
    "Compare the monthly closing prices of the three stocks in a line graph.",
    "Create a bar chart to show the distribution of closing prices for each stock.",
    "Show the yearly average closing price for each stock in a grouped bar chart."
  ]
}
```