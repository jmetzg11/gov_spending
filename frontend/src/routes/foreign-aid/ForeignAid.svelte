<script>
	import { onMount } from 'svelte';
	import { getData, ForeignAidDataProcessor } from './helpers';
	import Selectors from './Selectors.svelte';
	import Info from './Info.svelte';
	import Bars from './Bars.svelte';
	import Map from './Map.svelte';
	let data = $state([]);
	let selectedYear = $state('all');
	let allYears = $state([]);
	let selectedCountry = $state('all');
	let allCountries = $state([]);
	let barData = $state([]);
	let dataProcessor;

	let infoSentence = $state('');
	let infoAmount = $state('');

	let mapData = $state([]);

	$effect(() => {
		selectedCountry;
		selectedYear;

		if (dataProcessor) {
			// barData
			barData = dataProcessor.getBarData(selectedCountry);
			// info
			const { sentence, amount } = dataProcessor.getInfoValues(selectedCountry, selectedYear);
			infoSentence = sentence;
			infoAmount = amount;
			// mapData
			mapData = dataProcessor.getMapData(selectedYear);
		}
	});

	onMount(async () => {
		const rawData = await getData();
		dataProcessor = new ForeignAidDataProcessor(rawData);
		// selectors
		const selectors = dataProcessor.getSelectors();
		allYears = selectors.allYears;
		allCountries = selectors.allCountries;
		// barData
		barData = dataProcessor.getBarData(selectedCountry);
		// info
		const { sentence, amount } = dataProcessor.getInfoValues(selectedCountry, selectedYear);
		infoSentence = sentence;
		infoAmount = amount;
		// mapData
		mapData = dataProcessor.getMapData(selectedYear);
	});
</script>

<div class="flex flex-col w-full h-full overflow-hidden">
	<div class="w-full h-[30%]">
		<div class="container-top-part h-full w-full flex flex-col">
			<div class="flex h-[40%] w-full">
				<div class="flex-1">
					<Selectors
						bind:selectedYear
						bind:selectedCountry
						years={allYears}
						countries={allCountries}
					/>
				</div>
				<div class="flex-1"><Info {infoSentence} {infoAmount} /></div>
			</div>
			<div class="h-[60%] w-full">
				<Bars data={barData} />
			</div>
		</div>
	</div>
	<div class="flex-grow">
		<Map data={mapData} {selectedYear} />
	</div>
</div>
