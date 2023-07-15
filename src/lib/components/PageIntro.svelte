<script lang="ts">
  import Container from "./Container.svelte";
  export let eyebrow = "Our Portfolio";
  export let title = "Don't just take our word on it";
</script>

<Container classes="mt-24 sm:mt-32 lg:mt-40">
  <h1>
    <span class="text-base font-semibold leading-7 text-primary-600"
      >{eyebrow}</span
    >
  </h1>
  <p class="mt-2 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
    {title}
  </p>
  <slot />
</Container>
