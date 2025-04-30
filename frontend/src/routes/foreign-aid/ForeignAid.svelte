<script>
	import { onMount } from 'svelte';
	import Selectors from './Selectors.svelte';
	import Info from './Info.svelte';
	import Bars from './Bars.svelte';
	import Map from './Map.svelte';

	let data = $state();
	$inspect(data);

	async function getData() {
		const apiUrl = import.meta.env.VITE_API_URL || '/api';
		const url = apiUrl + '/foreign_aid';
		const response = await fetch(url);
		data = await response.json();
	}

	onMount(() => {
		getData();
	});
</script>

<div class="flex flex-col w-full h-full overflow-hidden">
	<div class="w-full h-[30%]">
		<div class="container-top-part h-full w-full flex flex-col">
			<div class="flex h-[40%] w-full">
				<div class="flex-1"><Selectors /></div>
				<div class="flex-1"><Info /></div>
			</div>
			<div class="h-[60%] w-full"><Bars /></div>
		</div>
	</div>
	<div class="flex-grow">
		<Map />
	</div>
</div>
