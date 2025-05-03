<script>
	import { onMount } from 'svelte';
	import { getData, getSelectors, getBarData } from './helpers';
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
	$effect(() => {
		console.log('something changed');
		console.log('Year:', selectedYear);
		console.log('Country:', selectedCountry);
		if (data.length > 0 && selectedCountry) {
			barData = getBarData(data, selectedCountry);
		}
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
				<div class="flex-1"><Info /></div>
			</div>
			<div class="h-[60%] w-full"><Bars data={barData} /></div>
		</div>
	</div>
	<div class="flex-grow">
		<Map />
	</div>
</div>
