<script>
  import "../app.postcss";
  import Navigation from "$lib/components/navigation/TopBar.svelte";
  import Footer from "$lib/components/Footer.svelte";
</script>

<Navigation />
<slot />
<Footer />
