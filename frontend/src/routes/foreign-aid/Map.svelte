<script>
	import { onMount } from 'svelte';
	let { data, selectedYear } = $props();
	let map;
	let markers = [];

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

	$effect(() => {
		selectedYear;
		data.length;
		if (map && data.length > 0) {
			clearMarkers();
			addMarkers(data);
		}
	});

	function initializeMap() {
		const corner1 = L.latLng(-90, -180);
		const corner2 = L.latLng(90, 180);
		const bounds = L.latLngBounds(corner1, corner2);
		map = L.map('map', {
			center: [0, 0],
			zoom: 3,
			maxBounds: bounds,
			maxBoundsViscosity: 0.8
		});

		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution:
				'&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
			noWrap: true
		}).addTo(map);

		addMarkers(data);
	}

	function addMarkers(data) {
		data.forEach((item) => {
			const { country, amount, normalizedAmount, lat, lng } = item;

			const marker = L.circleMarker([lat, lng], {
				radius: normalizedAmount,
				color: '#007BFF',
				fillColor: '#007BFF',
				fillOpacity: 0.5
			})
				.addTo(map)
				.bindTooltip(
					`<div style="font-size: 14px">${country}: $${parseFloat(amount).toLocaleString('en-US')}</div>`,
					{
						permanent: false,
						direction: 'top'
					}
				);
			markers.push(marker);
		});
	}

	function clearMarkers() {
		markers.forEach((marker) => {
			map.removeLayer(marker);
		});
		markers = [];
	}
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
