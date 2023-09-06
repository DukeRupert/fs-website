<script lang="ts">
	import PricingCards from './PricingCards.svelte';
  import FadeIn from "./FadeIn.svelte";
</script>

<section id="pricing" class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <FadeIn
      ><div class="mx-auto max-w-4xl text-center">
        <h2 class="text-base font-semibold leading-7 text-primary-600">
          Pricing
        </h2>
        <p
          class="mt-2 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl"
        >
          Affordable options for every business
        </p>
        <p
          class="mx-auto mt-6 max-w-2xl text-center text-lg leading-8 text-gray-600"
        >
          Choose from our most common services to take your business to the next
          level. We believe in transparent, no surprise pricing.
        </p>
      </div></FadeIn
    >
    <PricingCards />
  </div>
</section>
