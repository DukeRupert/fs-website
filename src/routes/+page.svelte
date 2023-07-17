<script lang="ts">
  import Seo from "$lib/components/SEO.svelte";
  import Hero from "$lib/components/Hero.svelte";
  import Services from "$lib/components/Services.svelte";
  import Features from "$lib/components/Features.svelte";
  import Pricing from "$lib/components/Pricing.svelte";
  import Testimonial from "$lib/components/Testimonial.svelte";
  import Cta from "$lib/components/Cta.svelte";
  import ContactForm from "$lib/components/ContactForm.svelte";
  import Portfolio from "$lib/components/Portfolio.svelte";

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

<svelte:head>
  <script type="application/ld+json">
  {
    "@context": "http://schema.org",
    "@type": "LocalBusiness",
    "name": "Firefly Software",
    "image": "https://www.fireflysoftware.dev/images/firefly/logo_two_small.png",
    "telephone": "+14068719875",
    "email": "service@fireflysoftware.dev",
    "aggregateRating": {
      "@type": "AggregateRating",
      "ratingValue": "5.0",
      "reviewCount": "1"
    },
    "brand": {
      "@type": "Brand",
      "name": "Firefly Software",
      "logo": "https://www.fireflysoftware.dev/images/firefly/Firefly_Logo_Blue.png"
    },
    "address": {
      "@type": "PostalAddress",
      "streetAddress": "652 Jena Loop",
      "addressLocality": "Helena",
      "addressRegion": "MT",
      "addressCountry": "US",
      "postalCode": "59602"
    },
    "openingHoursSpecification": {
      "@type": "OpeningHoursSpecification",
      "dayOfWeek": {
        "@type": "DayOfWeek",
        "name": "Monday, Tuesday, Wednesday, Thursday, Friday"
      }
    },
    "url": "https://www.fireflysoftware.dev/",
    "hasOfferCatalog": {
      "@type": "OfferCatalog",
      "name": "Firefly Software Services",
      "itemListElement": [
        {
          "@type": "Offer",
          "itemOffered": {
            "@type": "Service",
            "name": "Web Design",
            "description": "We handcraft a clean and beautiful design tailored to your business that loads fast, really fast.",
            "url": "https://www.fireflysoftware.dev/#services"
          }
        },
        {
          "@type": "Offer",
          "itemOffered": {
            "@type": "Service",
            "name": "Website Maintenance",
            "description": "Never worry about updates, bugs or hosting. We handle everything so you can focus on your business.",
            "url": "https://www.fireflysoftware.dev/#services"
          }
        },
        {
          "@type": "Offer",
          "itemOffered": {
            "@type": "Service",
            "name": "SEO",
            "description": "No gimmicks or tricks, just well structured HTML combined with useful information for your customers leads to the best results.",
            "url": "https://www.fireflysoftware.dev/#services"
          }
        }
      ]
    }
  }
  </script>
</svelte:head>
<Seo data={seoData} />
<main>
  <Hero />
  <Services />
  <Testimonial />
  <Features />
  <Pricing />
  <Cta />
</main>
