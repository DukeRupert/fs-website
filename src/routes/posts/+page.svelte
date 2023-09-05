<script lang="ts">
  import type { PageData } from "./$types";

  export let data: PageData;
  $: ({ posts } = data);
  $: console.log(posts);
</script>

<h1>Posts</h1>
