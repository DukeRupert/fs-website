<script lang="ts">
  import Container from "$lib/components/Container.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import Blockquote from "$lib/components/Blockquote.svelte";
  import PageLinks from "$lib/components/pagelinks/PageLinks.svelte";
  import Cta from "$lib/components/Cta.svelte";
  import kcc_splash from "$lib/assets/images/kcc_splash_4_3.png?as=run";
  import kcc_logo from "$lib/assets/images/kcc_logo.png?as=run";
  import kcc_kagen from "$lib/assets/images/kagen_coffee_and_crepes_owner_873.webp?as=run";
  import Img from "@zerodevx/svelte-img";
  import GrayscaleTransitionImage from "$lib/components/GrayscaleTransitionImage.svelte";
  import type { PageData } from "./$types";

  const caseStudy = {
    client: "Kagen Coffee and Crepes",
    title: "Crepes, coffee, and community",
    description:
      "A vacation led to love, which led to a dream, which became Kagen Coffee and Crepes in Richland, WA",
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
  };

  const pages = [
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

<article class="mt-24 sm:mt-32 lg:mt-40">
  <header>
    <!-- Page Intro -->
    <PageIntro eyebrow="Case Study" title={caseStudy.title} centered>
      <p
        class="mx-auto mt-6 max-w-2xl text-center text-lg leading-8 text-gray-600"
      >
        {caseStudy.description}
      </p>
    </PageIntro>

    <div
      class="mt-24 border-t border-neutral-200 bg-white/50 sm:mt-32 lg:mt-40"
    >
      <Container>
        <div class="mx-auto max-w-5xl">
          <dl
            class="-mx-6 grid grid-cols-1 text-sm text-neutral-950 sm:mx-0 sm:grid-cols-3"
          >
            <div
              class="border-t border-neutral-200 px-6 py-4 first:border-t-0 sm:border-l sm:border-t-0"
            >
              <dt class="font-semibold">Client</dt>
              <dd>{caseStudy.client}</dd>
            </div>
            <div
              class="border-t border-neutral-200 px-6 py-4 first:border-t-0 sm:border-l sm:border-t-0"
            >
              <dt class="font-semibold">Year</dt>
              <dd>
                <time dateTime={caseStudy.date.split("-")[0]}>
                  {caseStudy.date.split("-")[0]}
                </time>
              </dd>
            </div>
            <div
              class="border-t border-neutral-200 px-6 py-4 first:border-t-0 sm:border-l sm:border-t-0"
            >
              <dt class="font-semibold">Service</dt>
              <dd>{caseStudy.service}</dd>
            </div>
          </dl>
        </div>
      </Container>
    </div>

    <div class="border-y border-neutral-200 bg-neutral-100">
      <div class="-my-px mx-auto max-w-[76rem] bg-neutral-200">
        <a href={caseStudy.link} target="_blank" referrerpolicy="no-referrer">
          <GrayscaleTransitionImage>
            <Img
              src={kcc_splash}
              alt="Kagen Coffee and Crepes mobile website views"
              class="w-full filter grayscale-[0]"
              loading="eager"
            />
          </GrayscaleTransitionImage>
        </a>
      </div>
    </div>
  </header>
  <Container cls="mt-24 sm:mt-32 lg:mt-40">
    <article class="mx-auto max-w-prose text-left">
      <div class="prose prose-lg">
        <h2>Overview</h2>
        {#each caseStudy.summary as item}
          <p>{item}</p>
        {/each}

        <h2>What we did</h2>
        <ul role="list" class="list-none my-6 flex flex-wrap gap-4">
          {#each caseStudy.skills as skill}
            <li
              class="rounded-full bg-neutral-100 px-4 py-1.5 text-base text-neutral-600"
            >
              {skill}
            </li>
          {/each}
        </ul>
      </div>

      <Blockquote testimonial={caseStudy.testimonial} />
    </article>
  </Container>
  <PageLinks title="More case studies" cls="mt-24 sm:mt-32 lg:mt-40" {pages} />
  <Cta />
</article>
