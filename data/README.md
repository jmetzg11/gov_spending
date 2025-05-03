# Gov Spending

https://api.usaspending.gov/api/v2/

# Geting foreign aid data

go to https://www.foreignassistance.gov/data

click API in Complete Dataset

click Actions => Query data

bottom left click column manager and just keep Country Name, Fiscal Year, constant_dollar_amount then click appy

bottom left click filters to put years 2015-2024 (enter each year manually then press enter, site is lame and it's easier to add like (fiscal year = 2015 OR fiscal year = 2016 OR ...))

Then click export

**Update April 24, 2025**
click CSV in Complete Dataset. API is down because of DOGE

# Code in main.py

You can't run all functions at once because of rating limits. Also you have to manually download the foreign aid data as a csv file from the website first.

- fetch_country_latlng() - helper to get country lat and lng for the map
- fetch_agency_codes() - Each agency has a code which is used later to look up their budget and function spending
- fetch_budget_resources() - gets the budget resources for each year
- fetch_function_spending() - Agency spending can be categorized. I focused on Energy', 'Net Interest', 'Commerce and Housing Credit', 'Transportation', 'Agriculture', 'Health', 'Education, Training, Employment, and Social Services', 'National Defense'
- add_foreign_aid() - adds the foreign aid data to the database that comes from a large csv file that was downloaded before
- add_agency_budget() - adds the agency budget data to the database
- add_function_spending() - adds the function spending data to the database
