<script lang="ts">
  import image_1 from "$lib/assets/images/kcc_ecommerce_coffee.png?as=run";
  import Seo from "$lib/components/SEO.svelte";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import type { SimpleFeaturesListData } from "$lib/components/SimpleFeaturesList";
  import SimpleFeaturesList from "$lib/components/SimpleFeaturesList.svelte";
  import type { ProductScreenshotData } from "$lib/components/ProductScreenshot";
  import FeatureProductScreenshot from "$lib/components/ProductScreenshot.svelte";
  import type { CtaData } from "$lib/components/Cta";

  const ecommerce: ProductScreenshotData = {
    left: false,
    eyebrow: "Our expertise",
    title: "Custom e-commerce front-ends",
    description:
      "Our team specializes in crafting bespoke front-end solutions that prioritize enhanced product storytelling, seamless site-navigation, and exceptional site performance. If you're ready to redefine your online store, you've come to the right place.",
    image: {
      src: image_1,
      alt: "Kagen Coffee and Crepes e-commerce mobile view",
    },
    features: [
      {
        title: "Enhanced Product Storytelling",
        description:
          "Beyond showcasing your products, we're experts at creating narratives that captivate your customers. Our front-end designs infuse each product with a compelling story, engaging your audience on a deeper level and driving sales.",
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
        title: "Intuitive Site Navigation",
        description:
          "Navigating your online store should be a breeze. We ensure that your customers can easily find what they're looking for with intuitive menus, filters, and search functionality, making their shopping experience smooth and enjoyable.",
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
        title: "Optimized Site Performance",
        description:
          "Slow websites drive customers away. Our front-end solutions are finely tuned for speed, ensuring rapid load times and a responsive browsing experience on all devices. Say goodbye to bounce rates and hello to higher conversion rates.",
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

  const features: SimpleFeaturesListData = {
    left: true,
    eyebrow: "Better results",
    title: "Why choose Firefly Software?",
    description:
      "Choosing an experienced web developer means unlocking the power of expertise. Let us handle the technicalities, so you can focus on what you do best – growing your business. Elevate your e-commerce with a seasoned pro today!",
    features: [
      {
        title: "Experience Matters",
        description:
          "Our seasoned web developers bring years of experience and specialized e-commerce knowledge to the table, ensuring your front-end is built with precision.",
      },
      {
        title: "Responsive Support",
        description:
          "Our dedicated support team is always a message away, ready to address your questions and concerns promptly.",
      },
      {
        title: "Time Savings",
        description:
          "Why spend countless hours learning and building when you can have a professional front-end in a fraction of the time?",
      },
      {
        title: "Transparent Pricing",
        description: "No hidden fees or surprises. Our clear, competitive pricing allows you to budget confidently."
      }
    ],
  };

  const cta: CtaData = {
    title: "Ready to elevate your e-commerce?",
    description: "Illuminate your products, streamline your navigation, and supercharge your performance. Contact us today, and let's embark on this journey to redefine your online shopping experience!",
    primary_action: {
      name: "Get started",
      href: "/contact-us"
    },
    secondary_action: null
  }

  const eyebrow = "Elevate your storefront";
  const title = "E-commerce";
  const description = `Discover our custom e-commerce front-ends, expertly designed for top-notch performance and captivating product storytelling. Transform your online store with Firefly Software and watch your sales soar.`;

  const seoData = {
    title: `Firefly Software | ${title}`,
    description: description,
    url: "https://www.fireflysoftware.dev/e-commerce",
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
<PageContainer CtaData={cta} centered>
  <span slot="intro">
    <PageIntro {eyebrow} {title}>
      <p>
        Discover our custom e-commerce front-ends, expertly designed for top-notch performance and captivating product storytelling.
      </p>
    </PageIntro>
  </span>
  <span slot="body">
    <FeatureProductScreenshot {...ecommerce} />
    <SimpleFeaturesList {...features} />
  </span>
</PageContainer>
