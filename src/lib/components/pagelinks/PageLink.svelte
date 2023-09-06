<script lang="ts">
  import type { PageLink } from "./pageLink";
  import Border from "$lib/components/Border.svelte";
  import { formatDate } from "$lib/utils/formatDate";
  export let p: PageLink
</script>

<!-- PageLink -->
<article id={p.href}>
  <Border position="left" cls="relative flex flex-col items-start pl-8">
    <h3 class="mt-6 text-base font-semibold text-neutral-950">
      {p.title}
    </h3>
    <time dateTime={p.date} class="order-first text-sm text-neutral-600">
      {formatDate(p.date)}
    </time>
    <p class="mt-2.5 text-base text-neutral-600">{p.description}</p>
    <a
      href={p.href}
      class="mt-6 flex gap-x-3 text-base font-semibold text-neutral-950 transition hover:text-neutral-700"
      aria-label={`Read more: ${p.title}`}
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
