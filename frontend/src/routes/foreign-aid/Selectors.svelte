<script>
	let { selectedYear = $bindable(), selectedCountry = $bindable(), years, countries } = $props();
	$inspect(years);
</script>

<select bind:value={selectedYear}>
	{#each years as year}
		<option value={year}>{year}</option>
	{/each}
</select>

<select bind:value={selectedCountry}>
	{#each countries as country}
		<option value={country}>{country}</option>
	{/each}
</select>
