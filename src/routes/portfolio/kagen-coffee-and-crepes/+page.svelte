<script lang="ts">
  import type { CaseStudy } from "$lib/types/caseStudy";
  import type { PageLink } from "$lib/components/pagelinks/pageLink";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import PageLinks from "$lib/components/pagelinks/PageLinks.svelte";
  import ProjectShowcase from "../ProjectShowcase.svelte";
  import kcc_splash from "$lib/assets/images/kcc_splash_4_3.png?as=run";
  import kcc_logo from "$lib/assets/images/kcc_logo.png?as=run";
  import kcc_kagen from "$lib/assets/images/kagen_coffee_and_crepes_owner_873.webp?as=run";

  const caseStudy: CaseStudy = {
    client: "Kagen Coffee and Crepes",
    title: "Crepes, coffee, and community",
    description:
      "A vacation sparked a love that kindled a dream, ultimately giving birth to Kagen Coffee and Crepes in Richland, WA.",
    href: "/portfolio/kagen-coffee-and-crepes",
    link: "https://www.kagenandco.com/",
    summary: [
      "Like many entrepreneurs Kagen started his business with a basic do-it-yourself website. After his first year though he wisely realized that his business needed a better online presence in order to compete in the Tri-Cities area.",
      "When Kagen Coffee and Crepes approached us to optimize their online presence we delivered a highly performant website that highlights their special combination of delicious crepes and friendly atmosphere. It must have worked since now Kagen Coffee and Crepes is about to open their second location!",
    ],
    logo: {
      src: kcc_logo,
      alt: "Kagen Coffee and Crepes logo",
    },
    skills: [
      "Web Design",
      "Web Development",
      "Hosting",
      "CMS (Sanity)",
      "Backend (Shopify)",
      "Integration (Toast)",
    ],
    date: "2022-06",
    service: "Website development",
    testimonial: {
      author: { name: "Kagen Cox", role: "Owner" },
      src: kcc_kagen,
      content: `Firefly is an absolutely must have company working for you in your
              corner! They are super approachable, helpful and friendly! If you
              are looking to take your company to the next level I highly
              recommend that you give Firefly a call!`,
    },
    splash: {
      src: kcc_splash,
      alt: "Kagen Coffee and Crepes mobile website views"
    }
  };

  const pages: PageLink[] = [
    {
      client: "Ebony and Ivory",
      title: "A passion for music in Montana",
      description:
        "In the charming town of Helena, MT, there lived a music virtuoso whose unwavering passion for both music and children led her to open her very own enchanted studio for the sweet melodies of young hearts to flourish.",
      href: "/portfolio/ebony-and-ivory",
      date: "2023-07",
    },
    {
      href: "/portfolio/fts-excavation",
      client: "FtS Excavation",
      title: "Precision + Integrity",
      description:
        "A dedication to bring precision and integrity to every excavation job in the challenging Puget Sound region.",

      date: "2021-08",
    },
  ];
</script>

<svelte:head>
  <title>Firefly Software | Our Work | Kagen Coffee & Crepes</title>
</svelte:head>

<PageContainer centered={true}>
  <span slot="intro">
    <PageIntro eyebrow="Case Study" title={caseStudy.title} centered={true}>
      <p>
        {caseStudy.description}
      </p>
    </PageIntro>
  </span>

  <span slot="body">
    <div class="mx-auto max-w-5xl">
      <ProjectShowcase {caseStudy} />
    </div>
    <PageLinks title="More case studies" cls="mt-24 sm:mt-32 lg:mt-40" {pages} />
  </span>
</PageContainer>
