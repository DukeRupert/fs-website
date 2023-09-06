<script lang="ts">
  import type { CustomBlockComponentProps } from "@portabletext/svelte";

  interface NormalBullet {
    _key: string;
    listItem: string;
    markDefs: any[];
    children: Child[];
    level: number;
    _type: string;
    style: string;
  }

  interface Child {
    _key: string;
    _type: string;
    marks: string[];
    text: string;
  }

  export let portableText: CustomBlockComponentProps<NormalBullet>;
  $: ({ value } = portableText);
  $: console.log(value);
</script>

{#each value.children as item}
  <li>{item.text}</li>
{/each}
