<script lang="ts">
  import type { Project } from "./../../routes/portfolio/projects";
  import fts_logo from "$lib/assets/images/fts_logo.png?as=run";
  import kcc_logo from "$lib/assets/images/kcc_logo.png?as=run";
  import { formatDate } from "$lib/utils/formatDate";
  import Img from "@zerodevx/svelte-img";
  import Border from "./Border.svelte";
  import Container from "./Container.svelte";
  export let projects: Project[] = [];
</script>

<Container cls="mt-40">
  <h2 class="font-display text-2xl font-semibold text-neutral-950">
    Case studies
  </h2>
  <div class="mt-10 space-y-20 sm:space-y-24 lg:space-y-32">
    {#each projects as project}
      <article>
        <Border cls="grid grid-cols-3 gap-x-8 gap-y-8 pt-16">
          <div
            class="col-span-full sm:flex sm:items-center sm:justify-between sm:gap-x-8 lg:col-span-1 lg:block"
          >
            <div class="sm:flex sm:items-center sm:gap-x-6 lg:block">
              <Img
                src={project.logo.src}
                alt={project.logo.alt}
                class="h-16 w-16 flex-none"
              />
              <h3
                class="mt-6 text-sm font-semibold text-neutral-950 sm:mt-0 lg:mt-8"
              >
                {project.client}
              </h3>
            </div>
            <div class="mt-1 flex gap-x-4 sm:mt-0 lg:block">
              <p
                class="text-sm tracking-tight text-neutral-950 after:ml-4 after:font-semibold after:text-neutral-300 after:content-['/'] lg:mt-2 lg:after:hidden"
              >
                {project.service}
              </p>
              <p class="text-sm text-neutral-950 lg:mt-2">
                <time dateTime={project.date}>
                  {formatDate(project.date)}
                </time>
              </p>
            </div>
          </div>
          <div class="col-span-full lg:col-span-2 lg:max-w-2xl">
            <p class="font-display text-4xl font-medium text-neutral-950">
              <a href={project.href}>{project.title}</a>
            </p>
            <div class="mt-6 space-y-6 text-base text-neutral-600">
              {#each project.summary as paragraph}
                <p>{paragraph}</p>
              {/each}
            </div>
            <div class="mt-8 flex">
              <a
                href={project.href}
                aria-label={`Read case study: ${project.client}`}
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
                  <p>{project.testimonial.content}</p>
                </blockquote>
                <figcaption class="mt-6 font-semibold text-neutral-950">
                  {project.testimonial.author.name}, {project.testimonial.author
                    .role}
                </figcaption>
              </figure>
            </Border>
          </div></Border
        >
      </article>
    {/each}
  </div>
</Container>
