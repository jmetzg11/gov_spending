<script>
	import { writable } from 'svelte/store';
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import '../app.css';
	import Agency from './agency/Agency.svelte';
	import ForeignAid from './foreign-aid/ForeignAid.svelte';
	import FunctionSpending from './function-spending/FunctionSpending.svelte';
	import Info from './info/Info.svelte';

	const activePage = writable('agency');

	function setPage(page) {
		$activePage = page;
	}
</script>

<div class="flex space-x-4 mb-6">
	<button
		class="px-4 py-2 rounded cursor-pointer {$activePage === 'agency'
			? 'bg-blue-500 text-white'
			: 'bg-gray-200'}"
		onclick={() => setPage('agency')}>Agency</button
	>
	<button
		class="px-4 py-2 rounded cursor-pointer {$activePage === 'function-spending'
			? 'bg-blue-500 text-white'
			: 'bg-gray-200'}"
		onclick={() => setPage('function-spending')}>Function Spending</button
	>
	<button
		class="px-4 py-2 rounded cursor-pointer {$activePage === 'foreign-aid'
			? 'bg-blue-500 text-white'
			: 'bg-gray-200'}"
		onclick={() => setPage('foreign-aid')}>Foreign Aid</button
	>
	<button
		class="px-4 py-2 rounded cursor-pointer {$activePage === 'info'
			? 'bg-blue-500 text-white'
			: 'bg-gray-200'}"
		onclick={() => setPage('info')}>Info</button
	>
</div>

{#if $activePage === 'agency'}
	<Agency />
{:else if $activePage === 'function-spending'}
	<FunctionSpending />
{:else if $activePage === 'foreign-aid'}
	<ForeignAid />
{:else if $activePage === 'info'}
	<Info />
{/if}
