<script>
	import { onMount } from 'svelte';
	let data = $state(null);
	$inspect(data);

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

<div>agency</div>
