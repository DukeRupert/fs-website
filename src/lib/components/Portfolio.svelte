<script lang="ts">
  import FadeIn from "./FadeIn.svelte";
  import CaseStudies from "./CaseStudies.svelte";
  import PROJECTS from "../../routes/portfolio/projects";
</script>

<section id="pricing" class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <FadeIn>
      <div class="mx-auto max-w-2xl lg:text-center">
        <h2 class="text-base font-semibold leading-7 text-primary-600">
          Our work
        </h2>
        <p
          class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl"
        >
          Creating success for our clients
        </p>
        <!-- <p class="mt-6 text-lg leading-8 text-gray-600">
          Learn a how we have helped our clients succeed.
        </p> -->
      </div>
    </FadeIn>
  </div>
  <CaseStudies projects={PROJECTS} />
</section>
