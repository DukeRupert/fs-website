<script lang="ts">
  import image_1 from "$lib/assets/images/kcc_ecommerce_coffee.png?as=run";
  import image_2 from "$lib/assets/images/fts_landing.png?as=run";
  import Seo from "$lib/components/SEO.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import type { ProductScreenshotData } from "$lib/components/ProductScreenshot";
  import FeatureProductScreenshot from "$lib/components/ProductScreenshot.svelte";
  import Cta from "$lib/components/Cta.svelte";

  const load_fast: ProductScreenshotData = {
    left: false,
    eyebrow: "Fast page load",
    title: "Why loading lightning-fast matters",
    description:
      "At Firefly Software, we don't just build websites; we craft digital experiences that leave a lasting impression. Our secret ingredient? Lightning-fast page load speeds. In a world where every second counts, your website's speed is your competitive edge. Our handcrafted websites are meticulously optimized for speed, ensuring that your visitors don't just browse; they zoom through your content. Why does this matter?",
    image: {
      src: image_1,
      alt: "Kagen Coffee and Crepes e-commerce mobile view",
    },
    features: [
      {
        title: "Enhanced User Experience",
        description:
          "Extremely fast web page load speeds ensure that visitors to your site have a smooth and enjoyable experience. Quick loading times reduce frustration, keep users engaged, and encourage them to explore your content, contact you or make a purchase.",
        icon: `<svg
                    class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                >
                  <path stroke="none" d="M0 0h24v24H0z"/>
                  <path d="M12 20l-7 -7a4 4 0 0 1 6.5 -6a.9 .9 0 0 0 1 0a4 4 0 0 1 6.5 6l-7 7" />
                </svg>`,
      },
      {
        title: "Improved Search Engine Ranking",
        description:
          "Search engines like Google consider page load speed as a ranking factor. Faster-loading pages are more likely to rank higher in search results, increasing your website's visibility and attracting more organic traffic.",
        icon: `<svg
                  class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  stroke-width="2"
                  stroke="currentColor"
                  fill="none"
                  stroke-linecap="round"
                  stroke-linejoin="round"
              >
              <path stroke="none" d="M0 0h24v24H0z"/>
              <polyline points="3 17 9 11 13 15 21 7" />
              <polyline points="14 7 21 7 21 14" />
              </svg>`,
      },
      {
        title: "Higher Conversion Rates",
        description:
          "Speedy web pages lead to higher conversion rates. Whether you're selling products, generating leads, or simply providing information, a fast-loading site increases the likelihood that visitors will take the desired action, such as making a purchase or filling out a contact form.",
        icon: `<svg
                  class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  stroke-width="2"
                  stroke="currentColor"
                  fill="none"
                  stroke-linecap="round"
                  stroke-linejoin="round"
              >
              <line x1="12" y1="1" x2="12" y2="23" />
              <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6" />
              </svg>`,
      },
    ],
  };

  const support: ProductScreenshotData = {
    left: true,
    eyebrow: "Mobile-first design",
    title: "Why your website design matters",
    description:
      "In today's fast-paced digital landscape, the way your website performs on mobile devices is mission-critical. That's where Firefly Software steps in with our cutting-edge mobile-first design approach.",
    image: {
      src: image_2,
      alt: "Ebony and Ivory landing page mobile view",
    },
    features: [
      {
        title: "Google's Favorite",
        description:
          "Did you know that Google predominantly uses the mobile version of your site for indexing and ranking? Don't worry; we've got it covered. Our mobile-first approach optimizes your site for SEO, ensuring that potential clients find you easily.",
        icon: `<svg
                    class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                >
                    <path stroke="none" d="M0 0h24v24H0z"/>
                    <path d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3" />
                    <path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3" />
                    <circle cx="15" cy="9" r="1"  />
                </svg>`,
      },
      {
        title: "Consistent Branding",
        description:
          "With Firefly Software, you don't have to compromise on design. We maintain a consistent brand image across all devices, so your audience recognizes and trusts your business, no matter how they access your site.",
        icon: `<svg
                    class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                >
                    <path stroke="none" d="M0 0h24v24H0z"/>
                    <rect x="7" y="4" width="10" height="16" rx="1" />
                    <line x1="11" y1="5" x2="13" y2="5" />
                    <line x1="12" y1="17" x2="12" y2="17.01" />
                </svg>`,
      },
      {
        title: "Future-Proofing",
        description:
          "Mobile-first isn't just a trend; it's the future. We build websites with tomorrow in mind, ensuring your online presence remains relevant and competitive as mobile technology continues to evolve.",
        icon: `<svg
                    class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"/>
                </svg>`,
      },
    ],
  };

  const title = "Website development";
  const description = `Firefly Software specializes in website development services that prioritize
    lightning-fast load times and mobile-first design. Our commitment to
    delivering exceptional user experiences ensures that our clients not only
    attract more visitors but also convert them into loyal customers, driving
    their online success.`;

  const seoData = {
    title: `Firefly Software | ${title}`,
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
<PageIntro
  eyebrow="Website development"
  title="Exceptional Web Developers: Elevate Your Online Presence"
  centered={false}
>
  <p>
    Firefly Software specializes in website development services that prioritize
    lightning-fast load times and mobile-first design. Our commitment to
    delivering exceptional user experiences ensures that our clients not only
    attract more visitors but also convert them into loyal customers, driving
    their online success.
  </p>
</PageIntro>
<FeatureProductScreenshot {...load_fast} />
<FeatureProductScreenshot {...support} />
<Cta />
