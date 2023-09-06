<script lang="ts">
  import type { CtaData } from "./Cta";
  import Cta from "./Cta.svelte";
  export let cls = "";
  export let centered = false;
  export let CtaData: CtaData = {
    title: "Let's get started",
    description:
      "Find out today how we can improve the performance of your online presence.",
    primary_action: {
      name: "Get started",
      href: "/contact-us",
    },
    secondary_action: null,
  };
</script>

<!-- Content Container -->
<div class="bg-white w-full mt-12 py-24 sm:py-32 px-6 lg:px-8 {cls}">
  <div class="mx-auto max-w-2xl {centered ? 'text-center' : undefined}">
    <slot name="intro" />
  </div>
  <slot name="body" />
  <Cta {...CtaData} />
</div>
