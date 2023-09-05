<script lang="ts">
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import Blockquote from "$lib/components/Blockquote.svelte";
  import PageLinks from "$lib/components/pagelinks/PageLinks.svelte";
  import eandi_splash from "$lib/assets/images/ebony_and_ivory/splash_mobile.png?as=run";
  import eandi_logo from "$lib/assets/images/ebony_and_ivory/logo.png?as=run";
  import eandi_chante from "$lib/assets/images/ebony_and_ivory/chante.png?as=run";
  import Img from "@zerodevx/svelte-img";
  import GrayscaleTransitionImage from "$lib/components/GrayscaleTransitionImage.svelte";

  export let caseStudy = {
    client: "Ebony and Ivory",
    title: "A passion for music in Montana",
    description:
      "In the charming town of Helena, MT, there lived a music virtuoso whose unwavering passion for both music and children led her to open her very own enchanted studio for the sweet melodies of young hearts to flourish.",
    href: "/portfolio/ebony-and-ivory",
    link: "https://www.ebonyandivory.love/",
    summary: [
      "Working alongside the Ebony and Ivory music lessons team to bring their new website to life was a fantastic experience. From the initial discussions, it was clear that they had a clear vision and purpose for their online presence, which made the entire development process seamless and enjoyable. The end result is a visually stunning and highly user-friendly website that aligns perfectly with their goals. It's been incredibly satisfying to see the website's effectiveness in generating a consistent stream of new clients for their business.",
      "Since the website's launch, we've observed a significant increase in inquiries and bookings, which validates our approach to design and functionality. Our attention to detail and dedication to Ebony and Ivory's success have been key factors in achieving these results. It's rewarding to witness their music lessons reaching a broader audience, and we're excited to continue supporting the growth of Ebony and Ivory.",
    ],
    logo: {
      src: eandi_logo,
      alt: "Ebony and Ivory logo",
    },
    skills: ["Web Design", "Web Development", "Hosting", "Social Media"],
    date: "2023-07",
    service: "Website development",
    testimonial: {
      author: { name: "Chante' Williams", role: "Owner" },
      src: eandi_chante,
      content: `As a business owner, I'm all about efficiency, and Firefly Software totally delivered. Thanks to their work, our online presence has blown up, and we've got more students than ever.`,
    },
    splash: {
      src: eandi_chante,
      alt: "Fix me",
    },
  };

  export let pages = [
    {
      href: "/portfolio/kagen-coffee-and-crepes",
      client: "Kagen Coffee and Crepes",
      title: "Crepes, coffee, and community",
      description:
        "A vacation led to love, which led to a dream, which became Kagen Coffee and Crepes in Richland, WA",
      date: "2022-06",
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

<article class="mt-24 sm:mt-32 lg:mt-40">
  <div class="border-t border-neutral-200 bg-white/50">
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
  </div>

  <div class="border-y border-neutral-200 bg-neutral-100">
    <div class="-my-px mx-auto max-w-[76rem] bg-neutral-200">
      <a href={caseStudy.link} target="_blank" referrerpolicy="no-referrer">
        <GrayscaleTransitionImage>
          <Img
            src={caseStudy.splash.src}
            alt={caseStudy.splash.alt}
            class="w-full filter grayscale-[0]"
            loading="eager"
          />
        </GrayscaleTransitionImage>
      </a>
    </div>
  </div>

  <div class="mt-24 sm:mt-32 lg:mt-40">
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
  </div>
</article>
<PageLinks title="More case studies" cls="mt-24 sm:mt-32 lg:mt-40" {pages} />
