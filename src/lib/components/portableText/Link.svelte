<script lang="ts">
  import type { CustomBlockComponentProps } from "@portabletext/svelte";

  export let portableText: CustomBlockComponentProps<{
    href: string;
  }>;
  $: ({ value } = portableText);
</script>

<a href={value.href} target="_blank" rel="noreferrer">
  <slot />
</a>
