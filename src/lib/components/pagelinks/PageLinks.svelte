<script lang="ts">
  import Container from "$lib/components/Container.svelte";
  import SectionIntro from "$lib/components/SectionIntro.svelte";
  import PageLink from "./PageLink.svelte";

  export let cls: string;
  export let title: string = "Title";
  export let intro: string = "A short introduction string";
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

<!-- PageLinks -->
<div class="relative pt-24 sm:pt-32 lg:pt-40 {cls}">
  <SectionIntro {title} smaller />

  <Container cls={intro ? "mt-24" : "mt-16"}>
    <div class="grid grid-cols-1 gap-x-8 gap-y-16 lg:grid-cols-2">
      {#each pages as page}
        <PageLink {page} />
      {/each}
    </div>
  </Container>
</div>
