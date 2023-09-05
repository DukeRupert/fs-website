<script lang="ts">
  import type { PageData } from "./$types";
  import Seo from "$lib/components/SEO.svelte";
  import PageContainer from "$lib/components/PageContainer.svelte";
  import PageIntro from "$lib/components/PageIntro.svelte";

  export let data: PageData;
  $: ({ posts } = data);
  $: console.log(posts);
  
  const title = "Website development";
  const description = `Firefly Software specializes in website development services that prioritize
    lightning-fast load times and mobile-first design. Our commitment to
    delivering exceptional user experiences ensures that our clients not only
    attract more visitors but also convert them into loyal customers, driving
    their online success.`;

  const seoData = {
    title: `${title} | Firefly Software`,
    description: description,
    url: "https://www.fireflysoftware.dev/posts",
    og: {
      src: "https://www.fireflysoftware.dev/images/facebook_link_image.png",
      alt: "Firefly Software landing page",
      mimeType: "image/png",
      width: 1200,
      height: 648,
    },
  };
</script>

<Seo data={seoData} />
<PageIntro
  eyebrow="Posts"
  title="Demystifying the web and exploring best practices"
  centered={false}
>
  <p>
    This resource hub is your one-stop destination to harness the digital
    landscape's full potential. Stay informed, make informed decisions, and
    transform your online business with expert insights that ensure your website
    not only looks exceptional but also drives growth and customer engagement.
  </p>
</PageIntro>
<div
  class="mx-auto mt-16 grid max-w-2xl grid-cols-1 gap-x-8 gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-3"
>
  <article class="flex flex-col items-start justify-between">
    <div class="relative w-full">
      <img
        src="https://images.unsplash.com/photo-1496128858413-b36217c2ce36?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=3603&q=80"
        alt=""
        class="aspect-[16/9] w-full rounded-2xl bg-gray-100 object-cover sm:aspect-[2/1] lg:aspect-[3/2]"
      />
      <div
        class="absolute inset-0 rounded-2xl ring-1 ring-inset ring-gray-900/10"
      />
    </div>
    <div class="max-w-xl">
      <div class="mt-8 flex items-center gap-x-4 text-xs">
        <time datetime="2020-03-16" class="text-gray-500">Mar 16, 2020</time>
        <a
          href="#"
          class="relative z-10 rounded-full bg-gray-50 px-3 py-1.5 font-medium text-gray-600 hover:bg-gray-100"
          >Marketing</a
        >
      </div>
      <div class="group relative">
        <h3
          class="mt-3 text-lg font-semibold leading-6 text-gray-900 group-hover:text-gray-600"
        >
          <a href="#">
            <span class="absolute inset-0" />
            Boost your conversion rate
          </a>
        </h3>
        <p class="mt-5 line-clamp-3 text-sm leading-6 text-gray-600">
          Illo sint voluptas. Error voluptates culpa eligendi. Hic vel totam
          vitae illo. Non aliquid explicabo necessitatibus unde. Sed
          exercitationem placeat consectetur nulla deserunt vel. Iusto corrupti
          dicta.
        </p>
      </div>
      <div class="relative mt-8 flex items-center gap-x-4">
        <img
          src="https://images.unsplash.com/photo-1519244703995-f4e0f30006d5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
          alt=""
          class="h-10 w-10 rounded-full bg-gray-100"
        />
        <div class="text-sm leading-6">
          <p class="font-semibold text-gray-900">
            <a href="#">
              <span class="absolute inset-0" />
              Michael Foster
            </a>
          </p>
          <p class="text-gray-600">Co-Founder / CTO</p>
        </div>
      </div>
    </div>
  </article>

  <!-- More posts... -->
</div>
