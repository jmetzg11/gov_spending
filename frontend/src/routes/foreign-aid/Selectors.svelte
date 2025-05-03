<script>
	let { selectedYear = $bindable(), selectedCountry = $bindable(), years, countries } = $props();
</script>

<div class="flex flex-row gap-1 sm:gap-2 md:gap-4 p-1 sm:p-2 md:p-3 lg:p-4 h-full w-full">
	<div class="flex flex-col w-full lg:w-1/3">
		<label for="year-selected" class="text-sm font-semibold mb-1">Year</label>
		<select
			bind:value={selectedYear}
			class="p-0 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
		>
			{#each years as year}
				<option value={year}>{year}</option>
			{/each}
		</select>
	</div>

	<div class="flex flex-col w-full lg:w-2/3">
		<label for="country-select" class="text-sm font-semibold mb-1">Country</label>
		<select
			bind:value={selectedCountry}
			class="p-0 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
		>
			{#each countries as country}
				<option value={country}>{country}</option>
			{/each}
		</select>
	</div>
</div>
