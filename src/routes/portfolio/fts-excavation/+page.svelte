<script lang="ts">
  import type { CaseStudy } from "$lib/types/caseStudy";
  import type { PageLink } from "$lib/components/pagelinks/pageLink";
  import fts_splash from "$lib/assets/images/fts_splash_4_3.png?as=run";
  import fts_frank from "$lib/assets/images/fts_frank.jpg?as=run";
  import fts_logo from "$lib/assets/images/fts_logo.png?as=run";
  import Seo from "$lib/components/SEO.svelte";
  import ProjectShowcase from "../ProjectShowcase.svelte";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import PageLinks from "$lib/components/pagelinks/PageLinks.svelte";

  const caseStudy: CaseStudy = {
    client: "FtS Excavation",
    title: "Precision + Integrity",
    description:
      "A dedication to bring precision and integrity to every excavation job in the challenging Puget Sound region.",
    summary: [
      " Frank and Beverly founded FtS Excavation in 2019. The mission was to provide excavation services to the Puget Sound region in a manner that demonstrated the highest precision and integrity. We share their love of excellence and constant drive towards improvement.",
    ],
    logo: {
      src: fts_logo,
      alt: "FtS logo",
    },
    href: "/portfolio/fts-excavation",
    link: "https://www.fts-excavation.com/",
    date: "2021-08",
    skills: ["Web design", "Web development"],
    service: "Website development",
    testimonial: {
      author: { name: "Frank Sharp", role: "Owner" },
      content: `Firefly is absolutely the best way to get the most value from your website. Reasonably priced and headache free! I appreciate getting to speak with a real person who responds lightening fast, listens to what I want and can create anything and do it better than anyone. I highly recommend using Firefly software to create your website or any digital tool you need. We've decided to hire them for our app creation too.`,
      src: fts_frank,
    },
    splash: {
      src: fts_splash,
      alt: "FtS Excavation mobile website views",
    },
  };

  const pages: PageLink[] = [
    {
      href: "/portfolio/kagen-coffee-and-crepes",
      client: "Kagen Coffee and Crepes",
      title: "Crepes, coffee, and community",
      description:
        "A vacation sparked a love that kindled a dream, ultimately giving birth to Kagen Coffee and Crepes in Richland, WA.",
      date: "2022-06",
    },
    {
      client: "Ebony and Ivory",
      title: "A passion for music in Montana",
      description:
        "In the charming town of Helena, MT, there lived a music virtuoso whose unwavering passion for both music and children led her to open her very own enchanted studio for the sweet melodies of young hearts to flourish.",
      href: "/portfolio/ebony-and-ivory",
      date: "2023-07",
    },
  ];

  const eyebrow = "Case Study";
  const seoData = {
    title: `${caseStudy.title} | Firefly Software`,
    description: caseStudy.description,
    url: `https://www.fireflysoftware.dev${caseStudy.href}`,
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
    <PageIntro {eyebrow} title={caseStudy.title} centered={true}>
      <p>
        {caseStudy.description}
      </p>
    </PageIntro>
  </span>

  <span slot="body">
    <div class="mx-auto max-w-5xl">
      <ProjectShowcase {caseStudy} />
    </div>
    <PageLinks
      title="More case studies"
      cls="mt-24 sm:mt-32 lg:mt-40"
      {pages}
    />
  </span>
</PageContainer>
