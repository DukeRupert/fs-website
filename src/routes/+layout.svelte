<script>
  import "../app.postcss";
  import Navigation from "$lib/components/navigation/TopBar.svelte";
  import Footer from "$lib/components/Footer.svelte";

  // Vercel Analytics
  import { dev } from "$app/environment";
  import { inject } from "@vercel/analytics";
  inject({ mode: dev ? "development" : "production" });
</script>

<svelte:head>
  <script
    async
    src="https://www.googletagmanager.com/gtag/js?id=AW-10903277132"
  ></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() {
      dataLayer.push(arguments);
    }
    gtag("js", new Date());

    gtag("config", "AW-10903277132");
  </script>
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

<Navigation />
<slot />
<Footer />
