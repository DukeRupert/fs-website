<script lang="ts">
  import type { CaseStudy } from "$lib/types/caseStudy";
  import Blockquote from "$lib/components/Blockquote.svelte";
  import Img from "@zerodevx/svelte-img";
  import GrayscaleTransitionImage from "$lib/components/GrayscaleTransitionImage.svelte";

  export let caseStudy: CaseStudy;
</script>

<article class="mt-24 sm:mt-32 lg:mt-40">
  <div class="border-t border-neutral-200 bg-white/50">
    <div class="mx-auto max-w-5xl">
      <dl
        class="-mx-6 grid grid-cols-1 text-sm text-neutral-950 sm:mx-0 sm:grid-cols-3"
      >
        <div
          class="border-t border-neutral-200 px-6 py-4 first:border-t-0 sm:border-t-0"
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
