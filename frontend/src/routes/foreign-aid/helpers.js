export async function getData() {
	const apiUrl = import.meta.env.VITE_API_URL || '/api';
	const url = apiUrl + '/foreign_aid';
	const response = await fetch(url);
	const respJson = await response.json();
	return respJson.data;
}

export class ForeignAidDataProcessor {
	constructor(rawData) {
		this.rawData = rawData;
		this.dataByYear = {};
		this.dataByCountry = {};
		this.dataByCountryAndYear = {};

		this.rawData.forEach((item) => {
			// data by year
			if (!this.dataByYear[item.Year]) {
				this.dataByYear[item.Year] = [];
			}
			this.dataByYear[item.Year].push(item);

			// data by country
			if (!this.dataByCountry[item.Country]) {
				this.dataByCountry[item.Country] = [];
			}
			this.dataByCountry[item.Country].push(item);

			// data by country and year
			const key = `${item.Country}-${item.Year}`;
			if (!this.dataByCountryAndYear[key]) {
				this.dataByCountryAndYear[key] = [];
			}
			this.dataByCountryAndYear[key].push(item);
		});

		this.years = Object.keys(this.dataByYear).sort((a, b) => a - b);
		this.countries = Object.keys(this.dataByCountry).sort((a, b) => a.localeCompare(b));
	}

	getSelectors() {
		return {
			allYears: ['all', ...this.years],
			allCountries: ['all', ...this.countries]
		};
	}

	_getFilteredData(country, year) {
		if (country === 'all' && year === 'all') {
			return this.rawData;
		} else if (country === 'all') {
			return this.dataByYear[year] || [];
		} else if (year === 'all') {
			return this.dataByCountry[country] || [];
		} else {
			return this.dataByCountryAndYear[`${country}-${year}`] || [];
		}
	}

	_makeNumber = (amount) => {
		if (amount === undefined || amount === null || isNaN(amount)) {
			return '$0';
		}

		let val;
		if (amount > 1000000000000) {
			val = (amount / 1000000000000).toFixed(2);
			return `$${Number(val).toLocaleString('en')} Trillion`;
		} else if (amount > 1000000000) {
			val = (amount / 1000000000).toFixed(2);
			return `$${Number(val).toLocaleString('en')} Billion`;
		} else if (amount > 1000000) {
			val = (amount / 1000000).toFixed(2);
			return `$${Number(val).toLocaleString('en')} Million`;
		} else {
			return '$' + amount.toLocaleString('en');
		}
	};

	getInfoValues(country, year) {
		const filteredData = this._getFilteredData(country, year);
		const amount = filteredData.reduce((sum, item) => sum + item.Amount, 0);

		let sentence;
		if (year !== 'all') {
			sentence =
				country === 'all'
					? `Total foreign aid provided by the U.S. in ${year}`
					: `Total foreign aid provided by the U.S. to ${country} in ${year}`;
		} else {
			sentence =
				country === 'all'
					? 'Total foreign aid provided by the U.S. worldwide since 2017'
					: `Total foreign aid provided by the U.S. to ${country} since 2017`;
		}

		return { sentence, amount: this._makeNumber(amount) };
	}

	getBarData(selectedCountry) {
		const data =
			selectedCountry === 'all' ? this.rawData : this.dataByCountry[selectedCountry] || [];

		const cumulativeByYear = data.reduce((acc, item) => {
			acc[item.Year] = (acc[item.Year] || 0) + item.Amount;
			return acc;
		}, {});

		return Object.entries(cumulativeByYear)
			.map(([year, amount]) => ({ year: parseInt(year), amount }))
			.sort((a, b) => a.year - b.year);
	}

	_normalizeData(data) {
		const minValue = Math.min(...data.map((item) => item.Amount));
		const maxValue = Math.max(...data.map((item) => item.Amount));
		const minRange = 4;
		const maxRange = 40;

		return data.map((item) => {
			const normalizedAmount =
				((item.Amount - minValue) / (maxValue - minValue)) * (maxRange - minRange) + minRange;

			return {
				country: item.Country,
				amount: item.Amount,
				lat: item.Lat,
				lng: item.Lng,
				normalizedAmount
			};
		});
	}

	getMapData(selectedYear) {
		let mapData;
		if (selectedYear !== 'all') {
			mapData = this.dataByYear[selectedYear];
		} else {
			const aggregatedByCountr = {};

			this.rawData.forEach((item) => {
				if (!aggregatedByCountr[item.Country]) {
					aggregatedByCountr[item.Country] = {
						Country: item.Country,
						Amount: 0,
						Lat: item.Lat,
						Lng: item.Lng
					};
				}
				aggregatedByCountr[item.Country].Amount += item.Amount;
			});

			mapData = Object.values(aggregatedByCountr);
		}
		return this._normalizeData(mapData);
	}
}
