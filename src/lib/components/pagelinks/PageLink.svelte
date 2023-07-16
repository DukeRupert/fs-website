<script lang="ts">
  import Border from "$lib/components/Border.svelte";
  import { formatDate } from "$lib/utils/formatDate";
  export let page = {
    href: "#",
    title: "Title",
    date: "2022-06",
    description:
      "The description goes here to briefly describe what the project is about.",
  };
</script>

<!-- PageLink -->
<article id={page.href}>
  <Border position="left" cls="relative flex flex-col items-start pl-8">
    <h3 class="mt-6 text-base font-semibold text-neutral-950">
      {page.title}
    </h3>
    <time dateTime={page.date} class="order-first text-sm text-neutral-600">
      {formatDate(page.date)}
    </time>
    <p class="mt-2.5 text-base text-neutral-600">{page.description}</p>
    <a
      href={page.href}
      class="mt-6 flex gap-x-3 text-base font-semibold text-neutral-950 transition hover:text-neutral-700"
      aria-label={`Read more: ${page.title}`}
    >
      Read more
      <!-- Arrow Icon -->
      <svg
        class="w-6 flex-none fill-current"
        viewBox="0 0 24 6"
        aria-hidden="true"
      >
        <path fill="evenodd" clip="evenodd" d="M24 3 18 .5v2H0v1h18v2L24 3Z" />
      </svg>

      <span class="absolute inset-0" />
    </a>
  </Border>
</article>
