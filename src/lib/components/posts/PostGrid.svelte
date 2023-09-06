<script lang="ts">
  import type { Post } from "$lib/types/sanity";
  import PostCard from "./PostCard.svelte";

  export let posts: Post[];
</script>

<div class="mx-auto max-w-7xl sm:px-6 lg:px-8">
  <div
    class="mx-auto mt-16 grid max-w-2xl grid-cols-1 gap-x-8 gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-3"
  >
    {#each posts as post}
      <PostCard {post} />
    {/each}
  </div>
</div>
