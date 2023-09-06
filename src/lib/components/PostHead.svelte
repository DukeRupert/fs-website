<script lang="ts">
  import type { MainImage } from "$lib/types/sanity";
  import SanityImage from "./portableText/SanityImage.svelte";
  export let eyebrow = "Our Portfolio";
  export let title = "Don't just take our word on it";
  export let centered = true;
  export let image: MainImage;
</script>

<div class="mx-auto max-w-3xl text-base leading-7 text-gray-700">
  <p>
    <span class="text-base font-semibold leading-7 text-primary-600"
      >{eyebrow}</span
    >
  </p>
  <h1 class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
    {title}
  </h1>
  <div
    class="mt-10 max-w-2xl text-xl leading-8 text-gray-600 {centered
      ? 'mx-auto'
      : ''}"
  >
    <slot />
  </div>
  {#if image}
    <div class="mt-16">
      <SanityImage
        {image}
        maxWidth={768}
        alt={title}
        class="aspect-video rounded-xl bg-gray-50 object-cover"
      />
    </div>
  {/if}
</div>
