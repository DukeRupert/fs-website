<script lang="ts">
  import fts_splash from "$lib/assets/images/fts_splash_4_3.png?run";
  import fts_frank from "$lib/assets/images/fts_frank.jpg?run";
  import Img from "@zerodevx/svelte-img";
  import Container from "$lib/components/Container.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import GrayscaleTransitionImage from "$lib/components/GrayscaleTransitionImage.svelte";
  import PageLinks from "$lib/components/pagelinks/PageLinks.svelte";
  import Blockquote from "$lib/components/Blockquote.svelte";
  import Cta from "$lib/components/Cta.svelte";

  export const caseStudy = {
    client: "FtS Excavation",
    title: "Precision + Integrity",
    description:
      "A dedication to bring precision and integrity to every excavation job in the challenging Puget Sound region.",
    summary: [
      " Frank and Beverly founded FtS Excavation in 2019. The mission was to provide excavation services to the Puget Sound region in a manner that demonstrated the highest precision and integrity. We share their love of excellence and constant drive towards improvement.",
    ],

    date: "2021-08",
    skills: ["Web design", "Web development"],
    service: "Website development",
    testimonial: {
      author: { name: "Frank Sharp", role: "Owner" },
      content: `Firefly is absolutely the best way to get the most value from your website. Reasonably priced and headache free! I appreciate getting to speak with a real person who responds lightening fast, listens to what I want and can create anything and do it better than anyone. I highly recommend using Firefly software to create your website or any digital tool you need. We've decided to hire them for our app creation too.`,
      src: fts_frank,
    },
  };
</script>

<svelte:head>
  <title>Firefly Software | Our Work | FtS Excavation</title>
</svelte:head>

<div class="mt-24 sm:mt-32 lg:mt-40">
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
        <GrayscaleTransitionImage>
          <Img
            src={fts_splash}
            alt="FtS Excavation mobile website views"
            class="w-full filter grayscale-[0]"
            loading="eager"
          />
        </GrayscaleTransitionImage>
      </div>
    </div>
  </header>
  <Container cls="mt-24 sm:mt-32 lg:mt-40">
    <article class="mx-auto text-left">
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
  <PageLinks title="More case studies" cls="mt-24 sm:mt-32 lg:mt-40" />
  <Cta />
</div>
