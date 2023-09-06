<script lang="ts">
  import type { CtaData } from "$lib/components/Cta";
  import kcc_specials from "$lib/assets/images/kcc_specials.png?as=run";
  import Seo from "$lib/components/SEO.svelte";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import type { SimpleFeaturesListData } from "$lib/components/SimpleFeaturesList";
  import SimpleFeaturesList from "$lib/components/SimpleFeaturesList.svelte";
  import type { ProductScreenshotData } from "$lib/components/ProductScreenshot";
  import FeatureProductScreenshot from "$lib/components/ProductScreenshot.svelte";

  const maintenance: ProductScreenshotData = {
    left: false,
    eyebrow: "Our expertise",
    title: "Don't stress, call us",
    description:
      "At Firefly Software, our expert team is here to offer professional website repair and maintenance services that will keep your site running smoothly, securely, and at its best. Don't let common issues slow you down; let us light the way to a trouble-free website.",
    image: {
      src: kcc_specials,
      alt: "Kagen Coffee and Crepes crepe specials mobile view",
    },
    features: [
      {
        title: "Performance Optimization",
        description:
          "We understand the importance of lightning-fast load times. Our team will fine-tune your website for optimal speed, ensuring that visitors stay engaged and search engines give you a thumbs-up.",
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
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                </svg>`,
      },
      {
        title: "Bug Fixes",
        description:
          "Annoying bugs and glitches can turn away potential customers. We'll identify and fix issues, ensuring a seamless user experience.",
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
                <path d="M9 9v-1a3 3 0 0 1 6 0v1" />  
                <path d="M8 9h8a6 6 0 0 1 1 3v3a5 5 0 0 1 -10 0v-3a6 6 0 0 1 1 -3" />  
                <line x1="3" y1="13" x2="7" y2="13" />  
                <line x1="17" y1="13" x2="21" y2="13" />  
                <line x1="12" y1="20" x2="12" y2="14" />  
                <line x1="4" y1="19" x2="7.35" y2="17" />  
                <line x1="20" y1="19" x2="16.65" y2="17" />  
                <line x1="4" y1="7" x2="7.75" y2="9.4" />  
                <line x1="20" y1="7" x2="16.25" y2="9.4" />
              </svg>`,
      },
      {
        title: "Software Updates",
        description:
          "Staying up to date is crucial for security and functionality. We'll manage updates, plugins, and themes to keep your site running smoothly.",
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
    ],
  };

  const features: SimpleFeaturesListData = {
    left: true,
    eyebrow: "Your digital handyman",
    title: "Why Choose Firefly Software?",
    description:
      "A local web developer from Montana brings the advantage of personal attention and a deep understanding of your unique needs. With their expertise and commitment to swift solutions, your website problems will be resolved promptly, ensuring minimal downtime and maximum online success.",
    features: [
      {
        title: "Specialized Experience",
        description:
          "Our team knows the ins and outs of platforms such as Wix, Squarespace, and WordPress to address issues effectively.",
      },
      {
        title: "Customized Solutions",
        description:
          "Every website is unique, and we provide tailored solutions that match your specific needs and budget.",
      },
      {
        title: "Transparent Pricing",
        description:
          "No hidden fees or surprises. We offer clear, competitive quote and you only pay if we fix the problem.",
      },
      {
        title: "Timely Support",
        description: "When you need assistance, we're just a message away. Our responsive support team is ready to address your concerns promptly."
      }
    ],
  };

  const cta: CtaData = {
    title: "Ready to Get Started?",
    description: "Don't let performance and security issues hinder your online success. Contact Firefly Software today! We can illuminate your website's potential and keep it shining brightly!",
    primary_action: {
      name: "Get started",
      href: "/contact-us"
    },
    secondary_action: null
  }

  const eyebrow = "Digital TLC";
  const title = "Website Repair and Maintenance";
  const description = `Is your Wix, Squarespace, or WordPress website facing performance and security issues that are causing headaches? At Firefly Software, our expert team is here to offer professional website repair and maintenance services that will keep your site running smoothly, securely, and at its best. Don't let common issues slow you down; let us light the way to a trouble-free website.`;

  const seoData = {
    title: `Firefly Software | ${title}`,
    description: description,
    url: "https://www.fireflysoftware.dev/website-maintenance",
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
        Is your Wix, Squarespace, or WordPress website facing performance and security issues that are causing headaches?
      </p>
    </PageIntro>
  </span>
  <span slot="body">
    <FeatureProductScreenshot {...maintenance} />
    <SimpleFeaturesList {...features} />
  </span>
</PageContainer>
