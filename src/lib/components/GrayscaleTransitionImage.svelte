<script lang="ts">
  import IntersectionObserver from "svelte-intersection-observer";

  let element: HTMLElement;
  let intersecting: boolean;
  const threshold = 0.8;
</script>

<IntersectionObserver {element} bind:intersecting {threshold}>
  <div
    bind:this={element}
    class="group relative transition duration-300 filter {intersecting
      ? 'grayscale-0'
      : 'grayscale-[0.7]'} hover:grayscale-0"
  >
    <slot />
  </div>
</IntersectionObserver>
