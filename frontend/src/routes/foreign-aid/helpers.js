export async function getData() {
	const apiUrl = import.meta.env.VITE_API_URL || '/api';
	const url = apiUrl + '/foreign_aid';
	const response = await fetch(url);
	const respJson = await response.json();
	return respJson.data;
}

export function getSelectors(data) {
	const allYears = ['all', ...new Set(data.map((item) => item.Year))].sort((a, b) => {
		if (a === 'all') return -1;
		if (b === 'all') return 1;
		return a - b;
	});

	const allCountries = ['all', ...new Set(data.map((item) => item.Country))].sort((a, b) => {
		if (a === 'all') return -1;
		if (b === 'all') return 1;
		return a.localeCompare(b);
	});

	return { allYears, allCountries };
}

export function getBarData(data, selectedCountry) {
	let cumulativeData = {};
	if (selectedCountry === 'all') {
		data.forEach((item) => {
			if (item.Year in cumulativeData) {
				cumulativeData[item.Year] += item.Amount;
			} else {
				cumulativeData[item.Year] = item.Amount;
			}
		});
	} else {
		const countryData = data.filter((item) => item.Country === selectedCountry);
		countryData.forEach((item) => {
			if (item.Year in cumulativeData) {
				cumulativeData[item.Year] += item.Amount;
			} else {
				cumulativeData[item.Year] = item.Amount;
			}
		});
	}

	const dataArray = [];
	Object.entries(cumulativeData).forEach(([year, amount]) => {
		dataArray.push({
			year: parseInt(year),
			amount: amount
		});
	});

	dataArray.sort((a, b) => a.year - b.year);
	return dataArray;
}

export const makeInfoValues = (data, country, year) => {
	let amount = '';
	let sentence = '';
	if (country !== 'all') {
		data = data.filter((d) => d.Country === country);
	}
	if (year !== 'all') {
		data = data.filter((d) => d.Year === year);

		amount = data.reduce((acc, val) => {
			acc += val.Amount;
			return acc;
		}, 0);
		if (country === 'all') {
			sentence = `Total foreign aid provided by the U.S. in ${year}`;
			return { sentence, amount };
		} else {
			sentence = `Total foreign aid provided by the U.S. to ${country} in ${year}`;
			return { sentence, amount };
		}
	} else {
		amount = data.reduce((acc, val) => {
			acc += val.Amount;
			return acc;
		}, 0);
		if (country === 'all') {
			sentence = 'Total foreign aid provided by the U.S. worldwide since 2017';
			return { sentence, amount };
		} else {
			sentence = `Total foreign aid provided by the U.S. to ${country} since 2017`;
			return { sentence, amount };
		}
	}
	return { sentence: 'loading', amount: 0 };
};

export const makeNumber = (amount) => {
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
