<script lang="ts">
  import type { CustomBlockComponentProps } from "@portabletext/svelte";
  export let portableText: CustomBlockComponentProps<{
    code: string;
    language: string;
  }>;

  $: ({ value } = portableText);
</script>

<pre>
  <!-- If you want syntax highlighting on you code values, load PrismJs.com in your __layout.svelte 😉 -->
  <code class="language-{value.language}">{value.code}</code>
</pre>
