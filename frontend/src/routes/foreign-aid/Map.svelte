<script>
	import { onMount } from 'svelte';

	onMount(() => {
		const leafletCss = document.createElement('link');
		leafletCss.rel = 'stylesheet';
		leafletCss.href = 'https://unpkg.com/leaflet/dist/leaflet.css';
		document.head.appendChild(leafletCss);

		const leafletScript = document.createElement('script');
		leafletScript.src = 'https://unpkg.com/leaflet/dist/leaflet.js';
		leafletScript.onload = initializeMap;
		document.head.appendChild(leafletScript);
	});
</script>

<div id="map"></div>

<style>
	#map {
		height: 100%;
		width: 100%;
	}

	.leaflet-tooltip {
		font-size: 22px;
	}
</style>
