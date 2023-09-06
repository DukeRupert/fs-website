<script lang="ts">
  import { urlFor } from "$lib/db";
  import type { Image, MainImage } from "$lib/types/sanity";
  import { onMount } from "svelte";
  export let image: Image | MainImage;
  export let maxWidth = 1200;
  export let alt = "";
  let className = "";
  export { className as class };

  $: ({ width, height, aspectRatio } = image?.asset?.metadata.dimensions);

  // Once loaded, the image will transition to full opacity
  let loaded = false;
  function onLoad() {
    loaded = true;
  }

  onMount(() => {
    // Fallback if load event doesn't fire
    setTimeout(() => (loaded = true), 250);
  });
</script>

{#if image}
  <img
    class={className
      ? className
      : "w-full h-full object-center object-cover sm:w-full sm:h-full"}
    loading="lazy"
    src={urlFor(image.asset)
      .width(maxWidth)
      .fit("fillmax")
      .format("webp")
      .url()}
    {alt}
    {width}
    {height}
    style="aspect-ratio: {aspectRatio}; opacity: {loaded
      ? 1
      : 0}; transition: .3s cubic-bezier(0.11, 0, 0.5, 0) opacity;"
    on:load={onLoad}
  />
{/if}
