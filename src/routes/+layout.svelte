<script>
  import "../app.postcss";
  import Navigation from "$lib/components/navigation/TopBar.svelte";
  import Footer from "$lib/components/Footer.svelte";

  // Vercel Analytics
  import { dev } from "$app/environment";
  import { inject } from "@vercel/analytics";
  inject({ mode: dev ? "development" : "production" });
</script>

<svelte:head />

<Navigation />
<slot />
<Footer />
