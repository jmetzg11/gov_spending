<script>
	let { infoSentence, infoAmount } = $props();
</script>

<div
	class="flex flex-col justify-center items-center h-full w-full text-xs sm:text-sm md:text-md lg:text-lg"
>
	<div
		class="flex justify-center items-center w-full p-0 sm:p-1 md:p-2 lg:p-3 text-[10px] sm:text-xs md:text-sm lg:text-base"
	>
		{infoSentence}
	</div>
	<div class="font-semibold text-sm sm:text-base md:text-lg lg:text-xl">
		{infoAmount}
	</div>
</div>
