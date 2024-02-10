```
{
  "type": "JSON",
  "title": "National Parks Overview",
  "description": "Provides detailed information about various national parks, including activities, coordinates, and contact info.",
  "data_format": "This JSON file begins with a top-level object containing metadata fields 'total', 'limit', and 'start', followed by a 'data' array. Each element of the 'data' array is an object representing a national park, containing fields for unique ID, URL, park name, description, geographic coordinates, and nested objects and arrays for activities, topics, contact information, operating hours, addresses, images, weather info, etc. Each park object also includes fields for state codes, entrance fees, and relevant links. Special considerations for preprocessing include parsing nested objects and arrays, especially for 'activities', 'topics', 'contacts', and 'operatingHours'.",
  "fields": [
    "total",
    "limit",
    "start",
    "data[].id",
    "data[].url",
    "data[].fullName",
    "data[].parkCode",
    "data[].description",
    "data[].latitude",
    "data[].longitude",
    "data[].latLong",
    "data[].activities[]",
    "data[].topics[]",
    "data[].states",
    "data[].contacts.phoneNumbers[]",
    "data[].contacts.emailAddresses[]",
    "data[].entranceFees[]",
    "data[].entrancePasses[]",
    "data[].fees[]",
    "data[].directionsInfo",
    "data[].directionsUrl",
    "data[].operatingHours[]",
    "data[].addresses[]",
    "data[].images[]",
    "data[].weatherInfo",
    "data[].name",
    "data[].designation",
    "data[].relevanceScore"
  ],
  "suggested_visualizations": [
    "Map each park with pins, using latitude and longitude for location",
    "Create a timeline or calendar showing operating hours and exceptions for each park",
    "Draw a pie chart showing the distribution of parks by state",
    "Visualize the number and type of activities available at each park in a bar graph",
    "Show a gallery of images for each park with captions and credits"
  ]
}
```