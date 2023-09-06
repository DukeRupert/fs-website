<script lang="ts">
  export let eyebrow = "Don't just take our word on it";
  export let title = "Our Portfolio";
  export let centered = false;
</script>

<p>
  <span class="text-base font-semibold leading-7 text-primary-600"
    >{eyebrow}</span
  >
</p>
<p class="mt-2 text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
  {title}
</p>
<div
  class="mt-6 max-w-3xl text-xl text-neutral-600 {centered
    ? 'text-center'
    : undefined}"
>
  <slot />
</div>
