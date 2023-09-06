<script lang="ts">
    import Seo from "$lib/components/SEO.svelte";
    import PageContainer from "$lib/components/PageContainer.svelte";
    import PageIntro from "$lib/components/PageIntro.svelte";
    import PricingCards from '$lib/components/PricingCards.svelte';
  
    const eyebrow = "Simple and transparent"
    const title = "Pricing";
    const description = `At Firefly Software, we're committed to transparency. Our clear and honest pricing structure enables you to make informed decisions and invest in your online success with confidence.`;
  
    const seoData = {
      title: `${title} | Firefly Software`,
      description: description,
      url: "https://www.fireflysoftware.dev/website-development",
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
      <PageIntro
        {eyebrow}
        {title}
      >
        <p>
          {description}
        </p>
      </PageIntro>
    </span>
    <span slot="body">
      <PricingCards />
    </span>
  </PageContainer>