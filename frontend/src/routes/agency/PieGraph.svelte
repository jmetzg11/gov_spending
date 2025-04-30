<script>
	import { onMount } from 'svelte';
	import { formatNumber } from './helper';
	let { data, title } = $props();
	let ApexCharts;
	let chart;

	onMount(async () => {
		ApexCharts = (await import('apexcharts')).default;

		return () => {
			chart.destroy();
		};
	});

	$effect(() => {
		if (data && data.length > 0) {
			const options = {
				chart: {
					type: 'pie'
				},
				title: {
					text: title,
					align: 'center',
					style: {
						fontSize: '30px',
						color: '#333'
					}
				},
				labels: data.map((d) => d.agency),
				series: data.map((d) => d.proportion),
				colors: data.map((d) => `rgb(${d.color})`),
				legend: {
					show: false
				},
				dataLabels: {
					formatter: (value, { seriesIndex, w }) => {
						return `${data[seriesIndex].agency}: ${(data[seriesIndex].proportion * 100).toFixed(2)}%`;
					},
					distributed: true,
					style: {
						fontSize: '20px',
						fontFamily: 'Helvetica, Arial, sans-serif',
						fontWeight: 'bold',
						colors: undefined
					},
					background: {
						enabled: true,
						foreColor: '#fff',
						padding: 4,
						borderRadius: 2,
						borderWidth: 1,
						borderColor: '#fff',
						opacity: 0.9,
						dropShadow: {
							enabled: false,
							top: 1,
							left: 1,
							blur: 1,
							color: '#000',
							opacity: 0.45
						}
					},
					dropShadow: {
						enabled: false,
						top: 1,
						left: 1,
						blur: 1,
						color: '#000',
						opacity: 0.45
					}
				},
				tooltip: {
					enabled: true,
					y: {
						formatter: (value, { seriesIndex, w }) => {
							return `$${data[seriesIndex].amount.toLocaleString('en-US')}`;
						}
					},
					style: {
						fontsize: '20px',
						color: '#ffffff'
					},
					fillSeriesColor: false,
					custom: function ({ series, seriesIndex, dataPointIndex, w }) {
						const label = w.config.labels[seriesIndex];
						const value = formatNumber(data[seriesIndex].amount);
						return `<div style="background-color: #333333; padding: 8px; border-radius: 4px; color: #ffffff;">
                      ${label}: $${value}
                    </div>`;
					}
				},
				plotOptions: {
					pie: {
						dataLabels: {
							style: {
								colors: data.map((d) => d.textColor)
							}
						},
						customScale: 1,
						expandOnClick: true
					}
				},
				stats: {
					hover: {
						filter: {
							type: 'darket',
							value: 0.8
						}
					}
				}
			};
			chart = new ApexCharts(document.getElementById(title), options);
			chart.render();
		}
	});
</script>

<div id={title} class="chart-container"></div>

<style>
	.chart-container {
		width: 100%;
		height: 100%;
	}
</style>
