<script lang="ts">
  import PageIntro from "$lib/components/PageIntro.svelte";
  import CaseStudies from "$lib/components/CaseStudies.svelte";
  import Cta from "$lib/components/Cta.svelte";
  import PROJECTS from "./projects";
</script>

<svelte:head>
  <title>Firefly Software | Our Work</title>
</svelte:head>
<PageIntro
  eyebrow="Our work"
  title="Proven solutions for real-world problems."
  centered={false}
>
  <p>
    We believe in efficiency and maximizing our resources to provide the best
    value to our clients.
  </p>
</PageIntro>
<CaseStudies projects={PROJECTS} />
<Cta />
