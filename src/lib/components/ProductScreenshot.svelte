<script lang="ts">
  import kcc_landing from "$lib/assets/images/kcc_landing.png?as=run";
  import Img from "@zerodevx/svelte-img";
  import type { Feature, Image } from "./ProductScreenshot";
  export let left = false;
  export let eyebrow = "Deploy faster";
  export let title = "A better workflow";
  export let description = `Lorem ipsum, dolor sit amet consectetur adipisicing
                        elit. Maiores impedit perferendis suscipit eaque, iste
                        dolor cupiditate blanditiis ratione.`;
  export let image: Image = {
    src: kcc_landing,
    alt: "Kagen Coffee and Crepes mobile view screen capture",
  };

  export let features: Feature[] = [
    {
      title: "Enhanced User Experience",
      description:
        "Extremely fast web page load speeds ensure that visitors to your site have a smooth and enjoyable experience. Quick loading times reduce frustration, keep users engaged, and encourage them to explore your content, contact you or make a purchase.",
      icon: `<svg
                    class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                >
                  <path stroke="none" d="M0 0h24v24H0z"/>
                  <path d="M12 20l-7 -7a4 4 0 0 1 6.5 -6a.9 .9 0 0 0 1 0a4 4 0 0 1 6.5 6l-7 7" />
                </svg>`,
    },
    {
      title: "Improved Search Engine Ranking",
      description:
        "Search engines like Google consider page load speed as a ranking factor. Faster-loading pages are more likely to rank higher in search results, increasing your website's visibility and attracting more organic traffic.",
      icon: `<svg
                  class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  stroke-width="2"
                  stroke="currentColor"
                  fill="none"
                  stroke-linecap="round"
                  stroke-linejoin="round"
              >
              <path stroke="none" d="M0 0h24v24H0z"/>
              <polyline points="3 17 9 11 13 15 21 7" />
              <polyline points="14 7 21 7 21 14" />
              </svg>`,
    },
    {
      title: "Higher Conversion Rates",
      description:
        "Speedy web pages lead to higher conversion rates. Whether you're selling products, generating leads, or simply providing information, a fast-loading site increases the likelihood that visitors will take the desired action, such as making a purchase or filling out a contact form.",
      icon: `<svg
                  class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  stroke-width="2"
                  stroke="currentColor"
                  fill="none"
                  stroke-linecap="round"
                  stroke-linejoin="round"
              >
              <line x1="12" y1="1" x2="12" y2="23" />
              <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6" />
              </svg>`,
    },
  ];
</script>

<div class="overflow-hidden bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl">
    <div
      class="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 sm:gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-2"
    >
      <div class="lg:pr-8 lg:pt-4">
        <div class="lg:max-w-lg">
          <h2 class="text-base font-semibold leading-7 text-primary-600">
            {eyebrow}
          </h2>
          <p
            class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl"
          >
            {title}
          </p>
          <p class="mt-6 text-lg leading-8 text-gray-600">
            {description}
          </p>
          <dl
            class="mt-10 max-w-xl space-y-8 text-base leading-7 text-gray-600 lg:max-w-none"
          >
            {#each features as { title, description, icon }}
              <div class="relative pl-9">
                <dt class="inline font-semibold text-gray-900">
                  {@html icon}
                  {title}
                </dt>
                <dd class="inline">
                  {description}
                </dd>
              </div>
            {/each}
          </dl>
        </div>
      </div>
      {#if left}
        <div class="mt-16 sm:mt-24 lg:mt-0 lg:flex-shrink-0 lg:flex-grow">
          <svg
            viewBox="0 0 366 729"
            role="img"
            class="mx-auto w-[22.875rem] max-w-full drop-shadow-xl"
          >
            <title>App screenshot</title>
            <defs>
              <clipPath id="2ade4387-9c63-4fc4-b754-10e687a0d332">
                <rect width="316" height="684" rx="36" />
              </clipPath>
            </defs>
            <path
              fill="#4B5563"
              d="M363.315 64.213C363.315 22.99 341.312 1 300.092 1H66.751C25.53 1 3.528 22.99 3.528 64.213v44.68l-.857.143A2 2 0 0 0 1 111.009v24.611a2 2 0 0 0 1.671 1.973l.95.158a2.26 2.26 0 0 1-.093.236v26.173c.212.1.398.296.541.643l-1.398.233A2 2 0 0 0 1 167.009v47.611a2 2 0 0 0 1.671 1.973l1.368.228c-.139.319-.314.533-.511.653v16.637c.221.104.414.313.56.689l-1.417.236A2 2 0 0 0 1 237.009v47.611a2 2 0 0 0 1.671 1.973l1.347.225c-.135.294-.302.493-.49.607v377.681c0 41.213 22 63.208 63.223 63.208h95.074c.947-.504 2.717-.843 4.745-.843l.141.001h.194l.086-.001 33.704.005c1.849.043 3.442.37 4.323.838h95.074c41.222 0 63.223-21.999 63.223-63.212v-394.63c-.259-.275-.48-.796-.63-1.47l-.011-.133 1.655-.276A2 2 0 0 0 366 266.62v-77.611a2 2 0 0 0-1.671-1.973l-1.712-.285c.148-.839.396-1.491.698-1.811V64.213Z"
            />
            <path
              fill="#343E4E"
              d="M16 59c0-23.748 19.252-43 43-43h246c23.748 0 43 19.252 43 43v615c0 23.196-18.804 42-42 42H58c-23.196 0-42-18.804-42-42V59Z"
            />
            <foreignObject
              width="316"
              height="684"
              transform="translate(24 24)"
              clip-path="url(#2ade4387-9c63-4fc4-b754-10e687a0d332)"
            >
              <Img
                src={image.src}
                alt={image.alt}
                loading="eager"
                decoding="sync"
              />
            </foreignObject>
          </svg>
        </div>
      {:else}
        <div class="flex items-start justify-center lg:order-first">
          <div class="mt-16 sm:mt-24 lg:mt-0 lg:flex-shrink-0 lg:flex-grow">
            <svg
              viewBox="0 0 366 729"
              role="img"
              class="mx-auto w-[22.875rem] max-w-full drop-shadow-xl"
            >
              <title>App screenshot</title>
              <defs>
                <clipPath id="2ade4387-9c63-4fc4-b754-10e687a0d332">
                  <rect width="316" height="684" rx="36" />
                </clipPath>
              </defs>
              <path
                fill="#4B5563"
                d="M363.315 64.213C363.315 22.99 341.312 1 300.092 1H66.751C25.53 1 3.528 22.99 3.528 64.213v44.68l-.857.143A2 2 0 0 0 1 111.009v24.611a2 2 0 0 0 1.671 1.973l.95.158a2.26 2.26 0 0 1-.093.236v26.173c.212.1.398.296.541.643l-1.398.233A2 2 0 0 0 1 167.009v47.611a2 2 0 0 0 1.671 1.973l1.368.228c-.139.319-.314.533-.511.653v16.637c.221.104.414.313.56.689l-1.417.236A2 2 0 0 0 1 237.009v47.611a2 2 0 0 0 1.671 1.973l1.347.225c-.135.294-.302.493-.49.607v377.681c0 41.213 22 63.208 63.223 63.208h95.074c.947-.504 2.717-.843 4.745-.843l.141.001h.194l.086-.001 33.704.005c1.849.043 3.442.37 4.323.838h95.074c41.222 0 63.223-21.999 63.223-63.212v-394.63c-.259-.275-.48-.796-.63-1.47l-.011-.133 1.655-.276A2 2 0 0 0 366 266.62v-77.611a2 2 0 0 0-1.671-1.973l-1.712-.285c.148-.839.396-1.491.698-1.811V64.213Z"
              />
              <path
                fill="#343E4E"
                d="M16 59c0-23.748 19.252-43 43-43h246c23.748 0 43 19.252 43 43v615c0 23.196-18.804 42-42 42H58c-23.196 0-42-18.804-42-42V59Z"
              />
              <foreignObject
                width="316"
                height="684"
                transform="translate(24 24)"
                clip-path="url(#2ade4387-9c63-4fc4-b754-10e687a0d332)"
              >
                <Img
                  src={image.src}
                  alt={image.alt}
                  loading="eager"
                  decoding="sync"
                />
              </foreignObject>
            </svg>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
