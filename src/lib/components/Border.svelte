<script lang="ts">
  export let classes: string;
  export let position: "top" | "left" = "top";
  export let invert = false;

  function adjust_position(position: "top" | "left"): string {
    if (position === "top")
      return "before:left-0 before:top-0 before:h-px before:w-6 after:left-8 after:right-0 after:top-0 after:h-px";
    if (position === "left")
      return "before:left-0 before:top-0 before:h-6 before:w-px after:bottom-0 after:left-0 after:top-8 after:w-px";
    return "";
  }
</script>

<div
  class="relative before:absolute after:absolute {classes} {invert
    ? 'before:bg-white after:bg-white/10'
    : 'before:bg-neutral-950 after:bg-neutral-950/10'} {adjust_position(
    position
  )}"
>
  <slot />
</div>
