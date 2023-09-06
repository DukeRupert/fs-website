<script lang="ts">
  import Seo from "$lib/components/SEO.svelte";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import CaseStudies from "$lib/components/CaseStudies.svelte";
  import PROJECTS from "./projects";

  const eyebrow = "Proven solutions for real-world problems."
  const title = "Our Portfolio"
  const description = " We take great pride in highlighting our past web development work for clients, where our expertise and dedication have consistently resulted in stunning and effective web presences."
  const seoData = {
    title: `${title} | Firefly Software`,
    description: description,
    url: "https://www.fireflysoftware.dev/portfolio",
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
<PageContainer centered={true}>
  <span slot="intro">
    <PageIntro
      {eyebrow}
      {title}
      centered={true}
    >
      <p>
       {description}
      </p>
    </PageIntro>
  </span>

  <span slot="body">
    <div class="mx-auto max-w-5xl">
      <CaseStudies projects={PROJECTS} />
    </div>
  </span>
</PageContainer>
