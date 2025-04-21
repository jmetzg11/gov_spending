import csv
from collections import defaultdict
import json
from database import SessionLocal, ForeignAid

session = SessionLocal()

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
            if int(year) >= 2017:
                try:
                    if country != 'World' and 'Region' not in country:
                        data[country][year] += int(amount)
                except Exception as e:
                    print(country, e)

    entries = []
    with open('./data/lat_lng_mapper.json', 'r') as json_file:
        lat_lng_mapper = json.load(json_file)

    for country, years in data.items():
        for year, amount in years.items():
            entries.append()
            entries.append({
                'country': country,
                'year': year,
                'amount': amount,
                'lat': lat_lng_mapper[country]['lat'],
                'lng': lat_lng_mapper[country]['lng']
            })

    session.query(ForeignAid).delete()

    all_entries = []
