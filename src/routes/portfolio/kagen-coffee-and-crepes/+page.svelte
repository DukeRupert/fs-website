<script lang="ts">
  import Container from "$lib/components/Container.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";

  import kcc_splash from "$lib/assets/images/kcc_splash_4_3.png?run";
  import kagen_kcc from "$lib/assets/images/kagen_coffee_and_crepes_owner_873.webp?run";
  import Img from "@zerodevx/svelte-img";
  import GrayscaleTransitionImage from "$lib/components/GrayscaleTransitionImage.svelte";

  export const caseStudy = {
    client: "Kagen Coffee and Crepes",
    title: "Crepes, coffee, and community",
    description:
      "A vacation led to love, which led to a dream, which became Kagen Coffee and Crepes in Richland, WA",
    summary: [
      "Find love in the face of fear — Phobia is a dating app that matches users based on their mutual phobias so they can be scared together.",
      "We worked with Phobia to develop a new onboarding flow. A user is shown pictures of common phobias and we use the microphone to detect which ones make them scream, feeding the results into the matching algorithm.",
    ],

    date: "2018-06",
    service: "Website development",
    testimonial: {
      author: { name: "Kagen Cox", role: "Owner of Kagen Coffee and Crepes" },
      content: `Firefly is an absolutely must have company working for you in your
              corner! They are super approachable, helpful and friendly! If you
              are looking to take your company to the next level I highly
              recommend that you give Firefly a call!`,
    },
  };
</script>

<article class="mt-24 sm:mt-32 lg:mt-40">
  <header>
    <!-- Page Intro -->
    <PageIntro eyebrow="Case Study" title={caseStudy.title}>
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
            src={kcc_splash}
            alt="Kagen Coffee and Crepes mobile website views"
            class="w-full filter grayscale-[0]"
            loading="eager"
          />
        </GrayscaleTransitionImage>
      </div>
    </div>
  </header>
  <Container classes="mt-24 sm:mt-32 lg:mt-40">
    <article class="mx-auto prose prose-lg text-left">
      <div class="typography">
        <h2>Overview</h2>
        <p>
          Like many entrepreneurs Kagen started his business with a basic
          do-it-yourself website. After his first year though he wisely realized
          that his business needed a better online presence in order to compete
          in the Tri-Cities area.
        </p>
        <p>
          When Kagen Coffee and Crepes approached us to optimize their online
          presence we delivered a highly performant website that highlights
          their special combination of delicious crepes and friendly atmosphere.
          It must have worked since now Kagen Coffee and Crepes is about to open
          their second location!
        </p>
        <p>
          A few of the features required by this site was the ability to rapidly
          promote limited-time specials via a CMS and integrating a third-party
          online ordering solution. We also created a custom storefront for
          Kagen's Coffee Club ( a monthly coffee subscription ) utilizing
          Shopify as the backend.
        </p>

        <h2>What we did</h2>
      </div>
      <ul role="list" class="list-none my-6 flex flex-wrap gap-4">
        <li
          class="rounded-full bg-neutral-100 px-4 py-1.5 text-base text-neutral-600"
        >
          Design
        </li>
        <li
          class="rounded-full bg-neutral-100 px-4 py-1.5 text-base text-neutral-600"
        >
          Development
        </li>
        <li
          class="rounded-full bg-neutral-100 px-4 py-1.5 text-base text-neutral-600"
        >
          Shopify
        </li>
        <li
          class="rounded-full bg-neutral-100 px-4 py-1.5 text-base text-neutral-600"
        >
          Toast
        </li>
        <li
          class="rounded-full bg-neutral-100 px-4 py-1.5 text-base text-neutral-600"
        >
          CMS
        </li>
      </ul>

      <figure
        class="grid grid-cols-[auto,1fr] items-center gap-x-4 gap-y-8 sm:grid-cols-12 sm:grid-rows-[1fr,auto,auto,1fr] sm:gap-x-10 lg:gap-x-16 my-32"
      >
        <blockquote
          class="col-span-2 text-xl/7 text-neutral-600 sm:col-span-7 sm:col-start-6 sm:row-start-2"
        >
          <p>
            {caseStudy.testimonial.content}
          </p>
        </blockquote>
        <div
          class="col-start-1 row-start-2 overflow-hidden sm:col-span-5 sm:row-span-full"
        >
          <Img
            src={kagen_kcc}
            alt="Kagen Cox owner of Kagen Coffee and Crepes"
            class="grayscale h-full w-full rounded-2xl object-cover sm:aspect-[7/9] sm:h-auto sm:w-full"
          />
        </div>
        <figcaption
          class="text-sm text-neutral-950 sm:col-span-7 sm:row-start-3 sm:text-base"
        >
          <span class="font-semibold">{caseStudy.testimonial.author.name}</span
          ><span class="hidden font-semibold sm:inline">, </span><br
            class="sm:hidden"
          /><span class="sm:font-semibold"
            >{caseStudy.testimonial.author.role}</span
          >
        </figcaption>
      </figure>
    </article>
  </Container>
</article>
