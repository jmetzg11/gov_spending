<script>
	import { onMount } from 'svelte';
	let { data } = $props();
	let ApexCharts;
	let chart;
	let chartId = 'bar-graph';

	$inspect(data);

	onMount(async () => {
		ApexCharts = (await import('apexcharts')).default;

		return () => {
			if (chart) {
				chart.destroy();
			}
		};
	});

	$effect(() => {
		if (data && data.length > 0 && ApexCharts) {
			if (chart) {
				chart.destroy();
			}
			const options = {
				chart: {
					type: 'bar',
					height: '100%',
					width: '100%',
					toolbar: { show: false },
					margin: 0,
					padding: 0
				},
				title: {
					text: ''
				},
				series: [
					{
						name: 'Amount',
						data: data.map((item) => item.amount)
					}
				],
				xaxis: {
					categories: data.map((item) => item.year),
					labels: { style: { colors: '#000' } }
				},
				yaxis: { show: false },
				grid: { show: false, padding: { top: -30, right: 0, left: 0, bottom: 0 } },
				tooltip: { enabled: false },
				dataLabels: {
					enabled: true,
					formatter: function (val) {
						const result = val / 1000000;
						const formattedValue =
							result > 9999
								? `$${Number(result.toFixed(0)).toLocaleString('en')} mil`
								: `$${Number(result.toFixed(2)).toLocaleString('en')} mil`;
						return formattedValue;
					},
					style: { colors: ['#000'], fontWeight: '400', fontSize: '14px' }
				},
				plotOptions: {
					bar: {
						horizontal: false,
						dataLabels: { position: 'top' }
					}
				},
				responsive: [
					{
						breakpoint: 768, // Adjust for smaller screens
						options: {
							xaxis: {
								labels: {
									style: { fontSize: '12px' } // Smaller font size for axis labels
								}
							},
							dataLabels: {
								style: { fontSize: '10px' } // Smaller font size for data labels
							},
							plotOptions: {
								bar: {
									dataLabels: { position: 'top' }
								}
							}
						}
					},
					{
						breakpoint: 480, // Further adjustment for very small screens
						options: {
							xaxis: {
								labels: {
									style: { fontSize: '10px' } // Even smaller font size
								}
							},
							dataLabels: {
								style: { fontSize: '8px' } // Smaller font size for data labels
							}
						}
					}
				]
			};
			chart = new ApexCharts(document.getElementById(chartId), options);
			chart.render();
		}
	});
</script>

<div id={chartId} class="w-full h-full"></div>

<style>
	#bar-graph {
		height: 100%;
		width: 100%;
	}
</style>
