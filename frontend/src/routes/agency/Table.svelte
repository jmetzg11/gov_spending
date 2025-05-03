<script>
	let { data, title } = $props();
	import { formatNumber } from './helper';
</script>

<div class="text-center text-3xl font-bold mt-5 text-[#333]">{title}</div>
<div class="flex justify-center mb-5 overflow-hidden">
	<div class="max-w-4xl overflow-x-auto rounded-lg shadow-lg p-3">
		<table class="min-w-full bg-white">
			<thead>
				<tr class="bg-blue-500 text-white uppercase text-xs sm:text-sm md:text-base leading-normal">
					<th class="py-2 px-6 text-left">#</th>
					<th class="py-2 px-6 text-left">Agency</th>
					<th class="py-2 px-6 text-left">Amount</th>
					<th class="py-2 px-6 text-left">Proportion</th>
				</tr>
			</thead>
			<tbody class="text-gray-700 text-xs sm:text-sm md:text-base font-light">
				{#each data as item, index}
					<tr
						class="border-b border-gray-200 hover:bg-gray-100 transition duration-200 ease-in-out"
					>
						<td class="py-2 px-6 text-left">{index + 1}</td>
						<td class="py-2 px-6 text-left">{item.agency}</td>
						<td class="py-2 px-6 text-left">${formatNumber(item.amount)}</td>
						<td class="py-2 px-6 text-left">{(item.proportion * 100).toFixed(2)}%</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
