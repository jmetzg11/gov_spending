from helpers import fetch_country_latlng
from get_budget import fetch_agency_codes, fetch_budget_resources, fetch_function_spending

if __name__ == '__main__':
    # fetch_country_latlng()
    # fetch_agency_codes()
    # fetch_budget_resources(year=2025)
    # fetch_function_spending(years=list(range(2017, 2026)))
    fetch_function_spending(years=list(range(2018, 2026)))
