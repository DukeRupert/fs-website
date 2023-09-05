<script lang="ts">
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import CaseStudies from "$lib/components/CaseStudies.svelte";
  import PROJECTS from "./projects";
</script>

<svelte:head>
  <title>Firefly Software | Our Work</title>
</svelte:head>
<PageContainer centered={true}>
  <span slot="intro">
    <PageIntro
      eyebrow="Proven solutions for real-world problems."
      title="Our Portfolio"
      centered={true}
    >
      <p>
        We take great pride in highlighting our past web development work for
        clients, where our expertise and dedication have consistently resulted
        in stunning and effective web presences.
      </p>
    </PageIntro>
  </span>

  <span slot="body">
    <div class="mx-auto max-w-5xl">
      <CaseStudies projects={PROJECTS} />
    </div>
  </span>
</PageContainer>
