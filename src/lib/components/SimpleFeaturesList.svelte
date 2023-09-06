<script lang="ts">
  import type { Feature } from "./SimpleFeaturesList";
  import FadeIn from "./FadeIn.svelte";
  export let eyebrow = "Simple pricing";
  export let title = "Per Month, 6 Month Minimum Contract";
  export let description = `Zero down for a standard 5 page small business website. If you need
            more than that then we will negotiate custom pricing based on the
            number of additional pages and additional features. You own your
            domain, content, listing, and profiles. Cancel anytime without any
            fees or hassle after the first 6 months.`
  export let features: Feature[] = []
</script>

<section id="features" class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div
      class="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 sm:gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-3"
    >
      <FadeIn
        ><div>
          <p class="text-base font-semibold leading-7 text-primary-600">
            {eyebrow}
          </p>
          <h2
            class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl"
          >
            {title}
          </h2>
          <p class="mt-6 text-base leading-7 text-gray-600">
            {description}
          </p>
        </div></FadeIn
      >
      <dl
        class="col-span-2 grid grid-cols-1 gap-x-8 gap-y-10 text-base leading-7 text-gray-600 sm:grid-cols-2 lg:gap-y-16"
      >
        {#each features as {title, description}}
          <FadeIn
            ><div class="group relative pl-9">
              <dt class="font-semibold text-gray-900">
                <svg
                  class="absolute left-0 top-1 h-5 w-5 text-primary-500 group-hover:text-tertiary-500 group-hover:animate-bounce"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
                    clip-rule="evenodd"
                  />
                </svg>

                {title}
              </dt>
              <dd class="mt-2">
                {description}
              </dd>
            </div></FadeIn
          >
        {/each}
      </dl>
    </div>
  </div>
</section>
