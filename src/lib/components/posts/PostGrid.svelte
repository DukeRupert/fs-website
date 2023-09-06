<script lang="ts">
  import type { Post } from "$lib/types/sanity";
  import { page } from "$app/stores";
  import PostCard from "./PostCard.svelte";
  import Pagination from "./Pagination.svelte";

  export let posts: Post[];
  export let count = 0;
  export let pagination = false;

  $: category = $page.url.searchParams.get("category") ?? "";
</script>

<div class="mx-auto max-w-7xl sm:px-6 lg:px-8">
  {#if category}
    <div class="mx-auto mt-8 pt-8 flex">
      <a
        href={`/posts?limit=5&page=0`}
        class="flex items-center gap-x-2 z-10 rounded-full bg-gray-50 px-3 py-1.5 font-medium text-gray-600 hover:bg-gray-100"
        ><span>{category}</span><span>
          <svg
            class="h-4 w-4"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg></span
        ></a
      >
    </div>
  {/if}
  <div
    class="mx-auto mt-16 grid max-w-2xl grid-cols-1 gap-x-8 gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-3"
  >
    {#each posts as post}
      <PostCard {post} />
    {/each}
  </div>
  {#if pagination}
    <Pagination {count} />
  {/if}
</div>
