<script>
	import { onMount } from 'svelte';
	import '../app.css';
	import Agency from './agency/Agency.svelte';
	import ForeignAid from './foreign-aid/ForeignAid.svelte';
	import FunctionSpending from './function-spending/FunctionSpending.svelte';
	import Info from './info/Info.svelte';

	let activePage = $state('foreign-aid');

	function setPage(page) {
		activePage = page;
	}

	const infos = [
		{ text: 'Foreign Aid', value: 'foreign-aid' },
		{ text: 'Agency', value: 'agency' },
		{ text: 'Function Spending', value: 'function-spending' },
		{ text: 'Info', value: 'info' }
	];
</script>

<div class="h-screen flex flex-col">
	<div class="flex flex-wrap lg:flex-nowrap w-full justify-around items-center p-4 gap-2">
		{#each infos as info}
			<button
				class="px-4 py-2 rounded cursor-pointer {activePage === info.value
					? 'bg-blue-500 text-white'
					: 'bg-gray-200'}"
				onclick={() => setPage(info.value)}>{info.text}</button
			>
		{/each}
	</div>

	<div class="flex-grow">
		{#if activePage === 'agency'}
			<Agency />
		{:else if activePage === 'function-spending'}
			<FunctionSpending />
		{:else if activePage === 'foreign-aid'}
			<ForeignAid />
		{:else if activePage === 'info'}
			<Info />
		{/if}
	</div>
</div>
