<script lang="ts">
	import { page } from '$app/stores';
	export let count = 0;
	$: limit = Number($page.url.searchParams.get('limit') ?? '5');
	$: pageNum = Number($page.url.searchParams.get('page') ?? '0');
	$: totalPages = Math.ceil(count / limit);
</script>

<nav class="flex mt-20 items-center justify-between border-t border-gray-200 px-4 sm:px-0">
	<div class="-mt-px flex w-0 flex-1">
		{#if pageNum !== 0}
			<a
				href="/posts?limit={limit}&page={pageNum - 1}"
				class="inline-flex items-center border-t-2 border-transparent pr-1 pt-4 text-sm font-medium text-gray-500 dark:text-gray-400 hover:border-gray-300 hover:text-gray-700 dark:hover:text-gray-300"
			>
				<svg
					class="mr-3 h-5 w-5 text-gray-400"
					viewBox="0 0 20 20"
					fill="currentColor"
					aria-hidden="true"
				>
					<path
						fill-rule="evenodd"
						d="M18 10a.75.75 0 01-.75.75H4.66l2.1 1.95a.75.75 0 11-1.02 1.1l-3.5-3.25a.75.75 0 010-1.1l3.5-3.25a.75.75 0 111.02 1.1l-2.1 1.95h12.59A.75.75 0 0118 10z"
						clip-rule="evenodd"
					/>
				</svg>
				<span class="hidden md:block">Previous</span>
			</a>
		{/if}
	</div>
	<div class="md:-mt-px md:flex">
		{#each Array(totalPages) as _, index}
			<a
				href="/posts?limit={limit}&page={index}"
				class="inline-flex items-center border-t-2 border-transparent px-4 pt-4 text-sm font-medium {pageNum ===
				index
					? 'border-primary-500 text-primary-600'
					: 'text-gray-500 hover:border-gray-300 hover:text-gray-700 dark:hover:text-gray-300'}"
			>
				{index + 1}
			</a>
		{/each}
	</div>
	<div class="-mt-px flex w-0 flex-1 justify-end">
		<a
			href="/posts?limit={limit}&page={pageNum + 1}"
			class="inline-flex items-center border-t-2 border-transparent pl-1 pt-4 text-sm font-medium text-gray-500 dark:text-gray-400 hover:border-gray-300 hover:text-gray-700 dark:hover:text-gray-300"
		>
			<span class="hidden md:block">Next</span>

			<svg
				class="ml-3 h-5 w-5 text-gray-400"
				viewBox="0 0 20 20"
				fill="currentColor"
				aria-hidden="true"
			>
				<path
					fill-rule="evenodd"
					d="M2 10a.75.75 0 01.75-.75h12.59l-2.1-1.95a.75.75 0 111.02-1.1l3.5 3.25a.75.75 0 010 1.1l-3.5 3.25a.75.75 0 11-1.02-1.1l2.1-1.95H2.75A.75.75 0 012 10z"
					clip-rule="evenodd"
				/>
			</svg>
		</a>
	</div>
</nav>