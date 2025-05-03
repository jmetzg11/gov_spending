<script>
	import { onMount } from 'svelte';
	let xLabels = $state([]);
	let agencies = $state([]);
	let lineData = $state([]);
	let ApexCharts;
	let chart;
	let chartHeight = 0;
	let chartId = 'line-graph';

	const colors = [
		'rgba(0, 128, 128, 1)',
		'rgba(255, 99, 71, 1)',
		'rgba(124, 252, 0, 1)',
		'rgba(70, 130, 180, 1)',
		'rgba(0, 191, 255, 1)',
		'rgba(155, 69, 0, 1)',
		'rgba(138, 43, 226, 1)',
		'rgba(60, 179, 113, 1)'
	];

	function calculateChartHeight() {
		chartHeight = Math.max(window.innerHeight * 0.7, 300);
	}

	async function getData() {
		const apiUrl = import.meta.env.VITE_API_URL || '/api';
		const url = apiUrl + '/function-spending';
		const response = await fetch(url);
		const data = await response.json();
		console.log(data);
		lineData = data.data;
		xLabels = data.years;
		agencies = data.agencies;
	}

	onMount(async () => {
		ApexCharts = (await import('apexcharts')).default;
		getData();

		return () => {
			chart.destroy();
			window.removeEventListener('resize', calculateChartHeight);
		};
	});

	$effect(() => {
		if (xLabels.length > 0 && agencies.length > 0 && ApexCharts) {
			calculateChartHeight();
			window.addEventListener('resize', calculateChartHeight);
			renderChart();
		}
	});

	function renderChart() {
		const series = agencies.map((agency) => ({
			name: agency,
			data: lineData[agency].map((d) => d.amount)
		}));

		const options = {
			chart: {
				type: 'line',
				height: chartHeight,
				toolbar: { show: false }
			},
			series: series,
			xaxis: {
				categories: xLabels,
				title: { text: 'Years' }
			},
			yaxis: {
				title: { text: '$ Billions' },
				labels: {
					formatter: (value) =>
						`$${value.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}B`
				}
			},
			tooltip: {
				y: {
					formatter: (value) =>
						`$${value.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}B`
				},
				custom: function ({ series, seriesIndex, dataPointIndex, w }) {
					const yearValues = [];
					for (let i = 0; i < series.length; i++) {
						if (series[i][dataPointIndex] !== undefined) {
							yearValues.push({
								name: w.globals.seriesNames[i],
								value: series[i][dataPointIndex],
								color: w.globals.colors[i]
							});
						}
					}

					yearValues.sort((a, b) => b.value - a.value);

					let tooltipHTML = `<div class="apexcharts-tooltip-title">${xLabels[dataPointIndex]}</div><div>`;

					yearValues.forEach((item) => {
						const formattedValue = `$${item.value.toLocaleString(undefined, {
							minimumFractionDigits: 2,
							maximumFractionDigits: 2
						})}B`;

						tooltipHTML += `
							<div style="display: flex; align-items: center; margin: 5px 0;">
							<span style="background: ${item.color}; width: 10px; height: 10px; margin-right: 5px;"></span>
							<span style="flex: 1;">${item.name}:</span>
							<span style="font-weight: bold;">${formattedValue}</span>
							</div>
						`;
					});

					tooltipHTML += '</div>';

					return tooltipHTML;
				}
			},
			colors: agencies.map((_, index) => colors[index % colors.length]),
			stroke: {
				curve: 'smooth',
				width: 3
			},
			legend: {
				show: true,
				position: 'top'
			},
			responsive: [
				{
					breakpoint: 768,
					options: {
						chart: {
							height: '300px'
						},
						legend: {
							position: 'bottom'
						}
					}
				}
			]
		};

		chart = new ApexCharts(document.getElementById(chartId), options);
		chart.render();
	}
</script>

<div id={chartId} class="w-full"></div>
