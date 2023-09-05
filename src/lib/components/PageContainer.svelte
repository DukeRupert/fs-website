<script lang="ts">
  import Cta from "./Cta.svelte";
  export let cls = "";
  export let centered = false;
</script>

<!-- Content Container -->
<div class="bg-white mt-12 py-24 sm:py-32 px-6 lg:px-8 {cls}">
  <div class="mx-auto max-w-2xl {centered ? 'text-center' : undefined}">
    <slot name="intro" />
  </div>
  <slot name="body" />
  <Cta />
</div>
