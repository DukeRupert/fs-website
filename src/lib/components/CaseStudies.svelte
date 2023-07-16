<script lang="ts">
  import fts_logo from "$lib/assets/images/fts_logo.png?run";
  import kcc_logo from "$lib/assets/images/kcc_logo.png?run";
  import { formatDate } from "$lib/utils/formatDate";
  import Img from "@zerodevx/svelte-img";
  import Border from "./Border.svelte";
  import Container from "./Container.svelte";
  export let caseStudies = [
    {
      client: "Kagen Coffee and Crepes",
      title: "Crepes, coffee, and community",
      description:
        "A vacation led to love, which led to a dream, which became Kagen Coffee and Crepes in Richland, WA",
      href: "/portfolio/kagen-coffee-and-crepes",
      summary: [
        "Find love in the face of fear — Phobia is a dating app that matches users based on their mutual phobias so they can be scared together.",
        "We worked with Phobia to develop a new onboarding flow. A user is shown pictures of common phobias and we use the microphone to detect which ones make them scream, feeding the results into the matching algorithm.",
      ],
      logo: {
        src: kcc_logo,
        alt: "Kagen Coffee and Crepes logo",
      },
      date: "2018-06",
      service: "Website development",
      testimonial: {
        author: { name: "Kagen Cox", role: "Owner" },
        content: `Firefly is an absolutely must have company working for you in your
              corner! They are super approachable, helpful and friendly! If you
              are looking to take your company to the next level I highly
              recommend that you give Firefly a call!`,
      },
    },
    {
      client: "FtS Excavation",
      title: "Precision + Integrity",
      description:
        "A dedication to bring precision and integrity to every excavation job in the challenging Puget Sound region.",
      href: "/portfolio/fts-excavation",
      summary: [
        "Find love in the face of fear — Phobia is a dating app that matches users based on their mutual phobias so they can be scared together.",
        "We worked with Phobia to develop a new onboarding flow. A user is shown pictures of common phobias and we use the microphone to detect which ones make them scream, feeding the results into the matching algorithm.",
      ],
      logo: {
        src: fts_logo,
        alt: "FtS Excavation logo",
      },
      date: "2019-06",
      service: "Website development",
      testimonial: {
        author: { name: "Frank Sharp", role: "Owner" },
        content: `Firefly is absolutely the best way to get the most value from your website. Reasonably priced and headache free! I appreciate getting to speak with a real person who responds lightening fast, listens to what I want and can create anything and do it better than anyone. I highly recommend using Firefly software to create your website or any digital tool you need. We've decided to hire them for our app creation too.`,
      },
    },
  ];
</script>

<Container cls="mt-40">
  <h2 class="font-display text-2xl font-semibold text-neutral-950">
    Case studies
  </h2>
  <div class="mt-10 space-y-20 sm:space-y-24 lg:space-y-32">
    {#each caseStudies as caseStudy}
      <article>
        <Border cls="grid grid-cols-3 gap-x-8 gap-y-8 pt-16">
          <div
            class="col-span-full sm:flex sm:items-center sm:justify-between sm:gap-x-8 lg:col-span-1 lg:block"
          >
            <div class="sm:flex sm:items-center sm:gap-x-6 lg:block">
              <Img
                src={caseStudy.logo.src}
                alt={caseStudy.logo.alt}
                class="h-16 w-16 flex-none"
              />
              <h3
                class="mt-6 text-sm font-semibold text-neutral-950 sm:mt-0 lg:mt-8"
              >
                {caseStudy.client}
              </h3>
            </div>
            <div class="mt-1 flex gap-x-4 sm:mt-0 lg:block">
              <p
                class="text-sm tracking-tight text-neutral-950 after:ml-4 after:font-semibold after:text-neutral-300 after:content-['/'] lg:mt-2 lg:after:hidden"
              >
                {caseStudy.service}
              </p>
              <p class="text-sm text-neutral-950 lg:mt-2">
                <time dateTime={caseStudy.date}>
                  {formatDate(caseStudy.date)}
                </time>
              </p>
            </div>
          </div>
          <div class="col-span-full lg:col-span-2 lg:max-w-2xl">
            <p class="font-display text-4xl font-medium text-neutral-950">
              <a href={caseStudy.href}>{caseStudy.title}</a>
            </p>
            <div class="mt-6 space-y-6 text-base text-neutral-600">
              {#each caseStudy.summary as paragraph}
                <p>{paragraph}</p>
              {/each}
            </div>
            <div class="mt-8 flex">
              <a
                href={caseStudy.href}
                aria-label={`Read case study: ${caseStudy.client}`}
                class="rounded-md bg-primary-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
              >
                Read case study
              </a>
            </div>
            <Border position="left" cls="pl-8">
              <figure class="text-sm mt-12">
                <blockquote
                  class="text-neutral-600 [&>*]:relative [&>:first-child]:before:absolute [&>:first-child]:before:right-full [&>:first-child]:before:content-['“'] [&>:last-child]:after:content-['”']"
                >
                  <p>{caseStudy.testimonial.content}</p>
                </blockquote>
                <figcaption class="mt-6 font-semibold text-neutral-950">
                  {caseStudy.testimonial.author.name}, {caseStudy.testimonial
                    .author.role}
                </figcaption>
              </figure>
            </Border>
          </div></Border
        >
      </article>
    {/each}
  </div>
</Container>
