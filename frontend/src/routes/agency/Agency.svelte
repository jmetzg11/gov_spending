<script>
	import { onMount } from 'svelte';
	let data = $state({ first_pie: [], secons_pie: [], table: [] });
	import PieGraph from './PieGraph.svelte';
	import Table from './Table.svelte';

	async function getData() {
		const apiUrl = import.meta.env.VITE_API_URL || '/api';
		const url = apiUrl + '/agency';
		const response = await fetch(url);
		data = await response.json();
	}

	onMount(() => {
		getData();
	});
</script>

<div class="w-full h-full">
	<div class="flex flex-col sm:flex-row gap-4">
		<PieGraph data={data.first_pie} title={'Top 9 Budgets'} />
		<PieGraph data={data.second_pie} title={'Next Top 9 Budgets'} />
	</div>
	<Table data={data.table} title={`Remaining ${data.table.length} Budgets`} />
</div>
