<script lang="ts">
	import { quintOut } from 'svelte/easing';
	import { fly } from 'svelte/transition';
	import type { FlyParams } from 'svelte/transition';

	let is_subscription = true;

	const FLY: FlyParams = {
		x: -25,
		delay: 100,
		duration: 300,
		opacity: 0,
		easing: quintOut
	};
</script>

<section id="pricing" class="bg-white py-24 sm:py-32">
	<div class="mx-auto max-w-7xl px-6 lg:px-8">
		<div class="mx-auto max-w-4xl text-center">
			<h2 class="text-base font-semibold leading-7 text-primary-600">Pricing</h2>
			<p class="mt-2 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
				Affordable options for every small business
			</p>
		</div>
		<p class="mx-auto mt-6 max-w-2xl text-center text-lg leading-8 text-gray-600">
			Choose from our most common services to take your business to the next level. We believe in
			transparent, no surprise pricing.
		</p>
		<div class="mt-16 flex justify-center">
			<fieldset
				class="grid grid-cols-2 gap-x-1 rounded-full p-1 text-center text-xs font-semibold leading-5 ring-1 ring-inset ring-gray-200"
			>
				<legend class="sr-only">Payment frequency</legend>
				<label
					class="{is_subscription
						? 'bg-primary-600 text-white'
						: 'text-gray-500'} cursor-pointer rounded-full px-2.5 py-1"
				>
					<input
						on:click={() => (is_subscription = true)}
						type="radio"
						name="frequency"
						value="monthly"
						class="sr-only"
					/>
					<span>Subscribe</span>
				</label>
				<label
					class="{is_subscription
						? 'text-gray-500'
						: 'bg-primary-600 text-white'} cursor-pointer rounded-full px-2.5 py-1"
				>
					<input
						on:click={() => (is_subscription = false)}
						type="radio"
						name="frequency"
						value="annually"
						class="sr-only"
					/>
					<span>Buy</span>
				</label>
			</fieldset>
		</div>
		<div
			class="isolate mx-auto mt-10 grid max-w-md grid-cols-1 gap-8 lg:mx-0 lg:max-w-none lg:grid-cols-3"
		>
			<div class="rounded-3xl p-8 xl:p-10 ring-2 ring-primary-600">
				<div class="flex items-center justify-between gap-x-4">
					<h3 id="tier-startup" class="text-lg font-semibold leading-8 text-primary-600">
						Standard Marketing
					</h3>
					<p
						class="rounded-full bg-primary-600/10 px-2.5 py-1 text-xs font-semibold leading-5 text-primary-600"
					>
						Most popular
					</p>
				</div>
				<p class="mt-4 text-sm leading-6 text-gray-600">
					Convert visitors into clients by putting your best foot forward!
				</p>
				<p class="mt-6 flex items-baseline gap-x-1">
					{#if is_subscription}
						<!-- Price, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-4xl font-bold tracking-tight text-gray-900">$150</span>
						<!-- Payment frequency, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-sm font-semibold leading-6 text-gray-600">/month</span>
					{:else}
						<!-- Price, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-4xl font-bold tracking-tight text-gray-900">$3,500</span>
					{/if}
				</p>
				<a
					href="#contact-us"
					aria-describedby="tier-marketing"
					class="mt-6 block rounded-md py-2 px-3 text-center text-sm font-semibold leading-6 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600 bg-primary-600 text-white shadow-sm hover:bg-primary-500"
					>Contact us</a
				>
				<ul role="list" class="mt-8 space-y-3 text-sm leading-6 text-gray-600 xl:mt-10">
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						5 pages
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Unlimited edits
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Advanced analytics
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						24/7 customer service
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Includes hosting
					</li>
				</ul>
			</div>
			<div class="rounded-3xl p-8 xl:p-10 ring-1 ring-gray-200">
				<div class="flex items-center justify-between gap-x-4">
					<h3 id="tier-enterprise" class="text-lg font-semibold leading-8 text-gray-900">
						Advanced Marketing
					</h3>
				</div>
				<p class="mt-4 text-sm leading-6 text-gray-600">
					Increase engagement with your customers through blog posts, instagram integration and
					custom form flows.
				</p>
				<p class="mt-6 flex items-baseline gap-x-1">
					{#if is_subscription}
						<!-- Price, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-4xl font-bold tracking-tight text-gray-900">$250</span>
						<!-- Payment frequency, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-sm font-semibold leading-6 text-gray-600">/month</span>
					{:else}
						<!-- Price, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-4xl font-bold tracking-tight text-gray-900">$5,000</span>
					{/if}
				</p>
				<a
					href="#contact-us"
					aria-describedby="tier-enterprise"
					class="mt-6 block rounded-md py-2 px-3 text-center text-sm font-semibold leading-6 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600 text-primary-600 ring-1 ring-inset ring-primary-200 hover:ring-primary-300"
					>Contact us</a
				>
				<ul role="list" class="mt-8 space-y-3 text-sm leading-6 text-gray-600 xl:mt-10">
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						7 pages
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						CRM integration (aka. blog)
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Includes 3 custom forms
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Unlimited edits
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Advanced analytics
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						24/7 customer service
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Includes hosting
					</li>
				</ul>
			</div>
			<div class="group rounded-3xl p-8 xl:p-10 ring-1 ring-gray-200 hover:ring-primary-600">
				<div class="flex items-center justify-between gap-x-4">
					<h3 id="tier-ecommerce" class="text-lg font-semibold leading-8 text-gray-900">
						E-commerce
					</h3>
				</div>
				<p class="mt-4 text-sm leading-6 text-gray-600">
					Sell more product through compelling storytelling. We create custom e-commerce storefront
					using Shopify as the backend.
				</p>
				<p class="mt-6 flex items-baseline gap-x-1">
					{#if is_subscription}
						<!-- Price, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-4xl font-bold tracking-tight text-gray-900">$500</span>
						<!-- Payment frequency, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-sm font-semibold leading-6 text-gray-600">/month</span>
					{:else}
						<!-- Price, update based on frequency toggle state -->
						<span in:fly={FLY} class="text-4xl font-bold tracking-tight text-gray-900">$12,000</span
						>
					{/if}
				</p>
				<a
					href="#contact-us"
					aria-describedby="tier-freelancer"
					class="mt-6 block rounded-md py-2 px-3 text-center text-sm font-semibold leading-6 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600 text-primary-600 ring-1 ring-inset ring-primary-200 hover:ring-primary-300"
					>Contact us</a
				>
				<ul role="list" class="mt-8 space-y-3 text-sm leading-6 text-gray-600 xl:mt-10">
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						100 products
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						5 pages
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Advanced analytics
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						24/7 customer support
					</li>
					<li class="flex gap-x-3">
						<svg
							class="h-6 w-5 flex-none text-primary-600"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
								clip-rule="evenodd"
							/>
						</svg>
						Includes hosting
					</li>
				</ul>
			</div>
		</div>
	</div>
</section>
