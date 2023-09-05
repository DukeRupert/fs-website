<script lang="ts">
    import PostCard from "./PostCard.svelte";
    interface Post {
        title: string;
    }

    export let posts: Post[];
</script>
<div
  class="mx-auto mt-16 grid max-w-2xl grid-cols-1 gap-x-8 gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-3"
>
  {#each posts as post}
    <PostCard {post} />
  {/each}
</div>