import requests
import json

def fetch_country_latlng():
    url = "https://restcountries.com/v3.1/all"
    response = requests.get(url)
    countries = response.json()

    name_mapper = {
        'Cape Verde': 'Cabo Verde',
        'Ivory Coast': "Cote d'Ivoire",
        'Saint Kitts and Nevis': 'St. Kitts and Nevis',
        'Saint Lucia': 'St. Lucia',
        'Saint Vincent and the Grenadines': 'St. Vincent and Grenadines'
    }

    lat_lng_mapper = {}
    for country in countries:
        name = country['name']['common']
        if name in name_mapper:
            name = name_mapper[name]
        latlng = country['latlng']
        lat_lng_mapper[name] = {'lat': latlng[0], 'lng': latlng[1]}

    with open('lat_lng_mapper.json', 'w') as json_file:
        json.dump(lat_lng_mapper, json_file)
