<script lang="ts">
  import type { PageData } from "./$types";
  import { page } from "$app/stores";
  import { urlFor } from "$lib/db";
  import Seo from "$lib/components/SEO.svelte";
  import PostHead from "$lib/components/PostHead.svelte";
  import PortableText from "$lib/components/portableText/PortableText.svelte";
  import Cta from "$lib/components/Cta.svelte";

  export let data: PageData;
  $: ({ post } = data);
  const eyebrow = "Post";
  $: seoData = {
    title: `${post.title} | Firefly Software`,
    description: post?.excerpt ?? "An article from Firefly Software",
    url: $page.url.href,
    og: {
      src: urlFor(post.mainImage.asset)
        .width(1200)
        .height(675)
        .format("webp")
        .crop("focalpoint")
        .quality(80)
        .toString(),
      alt: post.mainImage.alt,
      mimeType: "webp",
      width: 1200,
      height: 675,
    },
    published_at: post.publishedAt,
  };
</script>

<Seo data={seoData} type="article" />

<div class="relative py-24 lg:py-32 overflow-hidden bg-white">
  <div class="relative px-4 sm:px-6 lg:px-8">
    <article class="mt-6 prose md:prose-lg dark:prose-light mx-auto">
      <PostHead
        eyebrow="Article"
        title={post.title}
        image={post.mainImage}
        centered
      />

      <div class="prose break-words">
        <PortableText value={post.body} />
      </div>
    </article>
    <Cta />
  </div>
</div>
