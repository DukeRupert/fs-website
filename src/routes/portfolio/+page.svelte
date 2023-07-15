<script lang="ts">
  import fts_service_hardscape from "$lib/assets/images/fts_service_hardscape.png?run";
  import fts_landing from "$lib/assets/images/fts_landing.png?run";
  import { FxReveal as Img } from "@zerodevx/svelte-img";
  import PageIntro from "$lib/components/PageIntro.svelte";
  import CaseStudies from "$lib/components/CaseStudies.svelte";

  const intro = {
    eyebrow: "Case Study",
    title: "Kagen Coffee & Crepes",
  };
</script>

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
<CaseStudies />
