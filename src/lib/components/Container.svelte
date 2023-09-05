<script lang="ts">
  export let cls = "";
  export let centered = false;
</script>

<!-- Content Container -->
<div class="{cls} {centered ? 'text-center' : undefined}">
  <div class="mx-auto max-w-2xl lg:max-w-none">
    <slot />
  </div>
</div>
