<script lang="ts">
  import SanityImage from "./SanityImage.svelte";
  import type { Image } from "$lib/types/sanity";
  import type { CustomBlockComponentProps } from "@portabletext/svelte";

  export let portableText: CustomBlockComponentProps<Image>;
  $: ({ value } = portableText);
</script>

{#if value.asset}
  <figure>
    <SanityImage image={value} />
  </figure>
{/if}
