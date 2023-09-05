<script lang="ts">
  import { PortableText } from "@portabletext/svelte";
  import type { InputValue } from "@portabletext/svelte/ptTypes";
  import Code from "./Code.svelte";
  import Link from "./Link.svelte";
  import ImageBlock from "./ImageBlock.svelte";
  import NormalList from "./NormalList.svelte";

  export let value: InputValue;
</script>

{#if value}
  <PortableText
    {value}
    components={{
      types: {
        code: Code,
        image: ImageBlock,
      },
      marks: {
        link: Link,
      },
    }}
  />
{/if}
