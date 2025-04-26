import csv
from collections import defaultdict
import json
from database import SessionLocal, ForeignAid, AgencyBudget, FunctionSpending

session = SessionLocal()

# download from https://www.foreignassistance.gov/data
def add_foreign_aid():
    data = defaultdict(lambda: defaultdict(int))
    replacements = {
        'China (Hong Kong, S.A.R., P.R.C.)': 'China',
        'China (P.R.C.)': 'China',
        'China (Tibet)': 'China',
        'Burma (Myanmar)': 'Myanmar',
        'Congo (Brazzaville)': 'Republic of the Congo',
        'Congo (Kinshasa)': 'DR Congo',
        'Korea, Democratic Republic of': 'North Korea',
        'Korea, Republic of': 'South Korea',
        'Sao Tome and Principe': 'São Tomé and Príncipe',
        'Slovak Republic': 'Slovakia',
        'Sudan (former)': 'Sudan',
        'West Bank and Gaza': 'Palestine',
        'Micronesia, Federated States of': 'Micronesia',
        'Kiribati, Republic of': 'Kiribati',
        'Curacao': 'Curaçao',
        'Serbia and Montenegro (former)': 'Serbia'
    }
    with open('./data/foreign_aid.csv', 'r') as csv_file:
        csv_reader = csv.DictReader(csv_file)
        for row in csv_reader:
            country = row['Country Name']
            year = row['Fiscal Year']
            amount = row['Constant Dollar Amount']
            country = replacements.get(country, country)
            if year.isdigit() and int(year) >= 2017:
                try:
                    if country != 'World' and 'Region' not in country:
                        data[country][year] += int(amount)
                except Exception as e:
                    print(country, e)

    with open('./data/lat_lng_mapper.json', 'r') as json_file:
        lat_lng_mapper = json.load(json_file)

    entries = []
    for country, years in data.items():
        for year, amount in years.items():
            try:
                entry = ForeignAid(
                    country=country,
                    year=int(year),
                    amount=float(amount),
                    lat=lat_lng_mapper[country]['lat'],
                    lng=lat_lng_mapper[country]['lng']
                )
                entries.append(entry)
            except KeyError:
                print(f"Missing lat/lng data for {country}")
            except Exception as e:
                print(f"Error creating entry for {country}, {year}: {e}")

    session.query(ForeignAid).delete()
    session.add_all(entries)
    session.commit()
    print(f"Added {len(entries)} foreign aid entries to the database")

def add_agency_budget():
    entries = []
    with open('./data/agency_resources.csv') as csv_file:
        csv_reader = csv.DictReader(csv_file)
        for row in csv_reader:
            try:
                entries.append(AgencyBudget(
                    agency=row['Agency'],
                    budget=row['Budget']
                ))
            except Exception as e:
                print(f"Error with agency budget: {e}")

    session.query(AgencyBudget).delete()
    session.add_all(entries)
    session.commit()
    print(f"Added {len(entries)} agency budget entries to the database")

def add_function_spending():
    entries = []
    for year in range(2017, 2026):
        print(f'on year {year}')
        with open(f'./data/functions_{year}.csv') as csv_file:
            csv_reader = csv.DictReader(csv_file)
            for row in csv_reader:
                try:
                    entries.append(FunctionSpending(
                        year=year,
                        name=row['Function'],
                        amount=row['Amount']
                    ))
                except Exception as e:
                    print(f"Error with function spending: {e}")

    session.query(FunctionSpending).delete()
    session.add_all(entries)
    session.commit()
    print(f"Added {len(entries)} function spending entries to the database")
