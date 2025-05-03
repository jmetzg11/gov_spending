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
