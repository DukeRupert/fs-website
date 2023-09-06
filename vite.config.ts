import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import { imagetools } from "@zerodevx/svelte-img/vite";

export default defineConfig({
  plugins: [sveltekit(), imagetools()],
  server: {
    // Set `host: true` if inside GitHub Codespaces to listen on all addresses,
    // see https://vitejs.dev/config/server-options.html#server-host
    host: !!process.env.CODESPACES,
  },
});
