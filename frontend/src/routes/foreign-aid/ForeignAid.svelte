<script>
	import { onMount } from 'svelte';
	import { getData, getSelectors, getBarData, makeInfoValues, makeNumber } from './helpers';
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

	let infoSentence = $state('');
	let infoAmount = $state('');

	$effect(() => {
		barData = getBarData(data, selectedCountry);
		const { sentence, amount } = makeInfoValues(data, selectedCountry, selectedYear);
		infoSentence = sentence;
		infoAmount = makeNumber(amount);
	});

	onMount(async () => {
		data = await getData();
		const selectors = getSelectors(data);
		allYears = selectors.allYears;
		allCountries = selectors.allCountries;
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
		<Map />
	</div>
</div>
