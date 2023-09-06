<script lang="ts">
  import type { PageData } from "./$types";
  import Seo from "$lib/components/SEO.svelte";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import PostGrid from "$lib/components/posts/PostGrid.svelte";

  export let data: PageData;
  $: ({ posts, count } = data);

  const eyebrow = "Resource hub";
  const title = "Posts";
  const description = `A collection of articles dedicated to clarifying web technologies and best practices, aiming to simplify the ever-changing web landscape for informed business decision-making.`;

  const seoData = {
    title: `${title} | Firefly Software`,
    description: description,
    url: "https://www.fireflysoftware.dev/posts",
    og: {
      src: "https://www.fireflysoftware.dev/images/facebook_link_image.png",
      alt: "Firefly Software landing page",
      mimeType: "image/png",
      width: 1200,
      height: 648,
    },
  };
</script>

<Seo data={seoData} />
<PageContainer centered>
  <span slot="intro">
    <PageIntro {eyebrow} {title}>
      <p>
        {description}
      </p>
    </PageIntro>
  </span>
  <span slot="body">
    <PostGrid {posts} />
  </span>
</PageContainer>
