<script lang="ts">
  import type { Post } from "$lib/types/sanity";
  import { urlFor } from "$lib/db";
  export let post: Post;
  console.log(post);
  // request a weekday along with a long date
  const options: Intl.DateTimeFormatOptions = {
    weekday: undefined,
    year: "numeric",
    month: "short",
    day: "numeric",
  };
  $: date = new Date(post.publishedAt);
</script>

<article class="flex flex-col items-start justify-between">
  <div class="relative w-full">
    <img
      src={urlFor(post.mainImage.asset).width(340).format("webp").toString()}
      alt={post?.title ?? "A stock image"}
      class="aspect-[16/9] w-full rounded-2xl bg-gray-100 object-cover sm:aspect-[2/1] lg:aspect-[3/2]"
    />
    <div
      class="absolute inset-0 rounded-2xl ring-1 ring-inset ring-gray-900/10"
    />
  </div>
  <div class="max-w-xl">
    <div class="mt-8 flex items-center gap-x-4 text-xs">
      <time
        datetime={date.toLocaleDateString("en-US", options)}
        class="text-gray-500">{date.toLocaleDateString("en-US", options)}</time
      >
      {#if post.categories}
        {#each post.categories as c}
          <a
            href={`/posts?limit=5&page=0&category=${c.title}`}
            class="relative z-10 rounded-full bg-gray-50 px-3 py-1.5 font-medium text-gray-600 hover:bg-gray-100"
            >{c.title}</a
          >
        {/each}
      {/if}
    </div>
    <div class="group relative">
      <h3
        class="mt-3 text-lg font-semibold leading-6 text-gray-900 group-hover:text-gray-600"
      >
        <a href={`/posts/${post.slug.current}`}>
          <span class="absolute inset-0" />
          {post.title}
        </a>
      </h3>
      <p class="mt-5 line-clamp-3 text-sm leading-6 text-gray-600">
        {post.excerpt}
      </p>
    </div>
    <div class="relative mt-8 flex items-center gap-x-4">
      <img
        src={urlFor(post.author.image)
          .width(120)
          .height(120)
          .format("webp")
          .toString()}
        alt={post.author.name}
        class="h-10 w-10 rounded-full bg-gray-100"
      />
      <div class="text-sm leading-6">
        <p class="font-semibold text-gray-900">
          <span class="absolute inset-0" />
          {post.author.name}
        </p>
        <p class="text-gray-600">{post.author.role}</p>
      </div>
    </div>
  </div>
</article>
