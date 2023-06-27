<script lang="ts">
  import Seo from "$lib/components/SEO.svelte";
  import Hero from "$lib/components/Hero.svelte";
  import Services from "$lib/components/Services.svelte";
  import Features from "$lib/components/Features.svelte";
  import Pricing from "$lib/components/Pricing.svelte";
  import Testimonial from "$lib/components/Testimonial.svelte";
  import ContactForm from "$lib/components/ContactForm.svelte";

  const seoData = {
    title: "Firefly Software - Handcrafted Web Design & Development",
    description:
      "Don't do-it-yourself with page builders like Wordpress, Wix and Squarespace. We offer handcrafted websites that produce superior results starting at $150/mo.",
    url: "https://www.fireflysoftware.dev/",
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
<main>
  <Hero />
  <Services />
  <Testimonial />
  <Features />
  <Pricing />
  <ContactForm />
</main>
