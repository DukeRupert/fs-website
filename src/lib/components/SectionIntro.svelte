<script lang="ts">
  import Container from "$lib/components/Container.svelte";
  export let title: string = "title";
  export let invert = false;
  export let smaller = false;
</script>

<Container>
  <div class="max-w-2xl">
    <h2>
      <span
        class="mt-6 text-xl {smaller
          ? 'text-2xl font-semibold'
          : 'text-4xl font-medium sm:text-5xl'} {invert
          ? 'text-neutral-300'
          : 'text-neutral-600'}"
      >
        {title}
      </span>
    </h2>
    <div
      class="mt-6 text-xl {invert ? 'text-neutral-300' : 'text-neutral-600'}"
    >
      <slot />
    </div>
  </div>
</Container>
