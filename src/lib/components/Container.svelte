<script lang="ts">
  export let classes = "";
  export let centered = false;
</script>

<!-- Content Container -->
<div
  class="mx-auto max-w-7xl px-6 lg:px-8 {classes} {centered
    ? 'text-center'
    : undefined}"
>
  <div class="mx-auto max-w-2xl lg:max-w-none">
    <slot />
  </div>
</div>
