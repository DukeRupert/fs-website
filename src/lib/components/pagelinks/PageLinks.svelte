<script lang="ts">
  import type { PageLink as PageLinkData } from "./pageLink";
  import Container from "$lib/components/Container.svelte";
  import SectionIntro from "$lib/components/SectionIntro.svelte";
  import PageLink from "./PageLink.svelte";

  export let cls: string;
  export let title: string = "Title";
  export let intro: string = "A short introduction string";
  export let pages: PageLinkData[] = []
</script>

<!-- PageLinks -->
<div class="relative pt-24 sm:pt-32 lg:pt-40 {cls}">
  <SectionIntro {title} smaller />

  <Container cls={intro ? "mt-24" : "mt-16"}>
    <div class="grid grid-cols-1 gap-x-8 gap-y-16 lg:grid-cols-2">
      {#each pages as page}
        <PageLink p={page} />
      {/each}
    </div>
  </Container>
</div>
