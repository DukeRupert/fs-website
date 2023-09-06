<script lang="ts">
  import { quintOut } from "svelte/easing";
  import { fly, fade } from "svelte/transition";
  import type { FlyParams, FadeParams } from "svelte/transition";
  import Img from "@zerodevx/svelte-img";
  import small_logo from "$lib/assets/images/logo/logo_small.jpg?as=run";
  import large_logo from "$lib/assets/images/logo/logo_two_full.jpg?as=run";
  import Socials from "$lib/components/Socials.svelte";
  import { socials } from "$lib/fs";

  // Toggle mobile menu
  let is_mobile_open = false;
  function close(): void {
    is_mobile_open = false;
  }

  // Toggle services flyout menu
  let is_services_open = false;
  function close_services(): void {
    is_services_open = false;
  }

  // Animation params
  let DURATION = 300;
  const fade_backdrop: FadeParams = {
    duration: DURATION,
    easing: quintOut,
  };
  const fly_in: FlyParams = {
    duration: DURATION,
    x: 300,
    opacity: 0,
    easing: quintOut,
  };
</script>

<header class="absolute inset-x-0 top-0 z-50">
  <nav
    class="flex items-center justify-between p-6 lg:px-8"
    aria-label="Global"
  >
    <div class="flex lg:flex-1">
      <a href="/" class="-m-1.5 p-1.5">
        <span class="sr-only">Firefly Software</span>
        <Img class="h-14 w-auto" src={large_logo} alt="Firefly Software" />
      </a>
    </div>
    <div class="flex lg:hidden">
      <button
        type="button"
        on:click={() => (is_mobile_open = true)}
        class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
      >
        <span class="sr-only">Open main menu</span>
        <svg
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
          />
        </svg>
      </button>
    </div>
    <div class="hidden lg:flex lg:gap-x-12">
      <div class="relative">
        <button on:click={() => is_services_open = !is_services_open} type="button" class="flex items-center gap-x-1 text-sm font-semibold leading-6 text-gray-900" aria-expanded="false">
          Services
          <svg class="h-5 w-5 flex-none text-gray-400 transition-transform duration-150 ease-out {is_services_open ? undefined : "rotate-180"}" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
          </svg>
        </button>

        <!--'Services' flyout menu, show/hide based on flyout menu state.-->
        <div class="absolute -left-8 top-full z-10 mt-3 w-screen max-w-md overflow-hidden rounded-3xl bg-white shadow-lg ring-1 ring-gray-900/5 transition {is_services_open ? "ease-out duration-200 opacity-100 translate-y-0" : "ease-in duration-150 opacity-0 translate-y-1"}">
          <div class="p-4">
            <div class="group relative flex items-center gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50">
              <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gradient-to-br from-tertiary-500 via-primary-500 to-secondary-500 group-hover:from-tertiary-400 group-hover:via-tertiary-500 group-hover:to-primary-500">
                <svg class="h-6 w-6 text-gray-100 group-hover:text-white transition duration-150 ease-out" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" aria-hidden="true">
                  <path stroke="none" d="M0 0h24v24H0z"/>  <polyline points="7 8 3 12 7 16" />  <polyline points="17 8 21 12 17 16" />  <line x1="14" y1="4" x2="10" y2="20" />
                </svg>
              </div>
              <div class="flex-auto">
                <a on:click={close_services} href="/website-development" class="block font-semibold text-gray-900">
                  Website Development
                  <span class="absolute inset-0"></span>
                </a>
                <p class="mt-1 text-gray-600">Elevate your online presence.</p>
              </div>
            </div>
            <div class="group relative flex items-center gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50">
              <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gradient-to-br from-tertiary-500 via-primary-500 to-secondary-500 group-hover:from-tertiary-400 group-hover:via-tertiary-500 group-hover:to-primary-500">
                <svg class="h-6 w-6 text-gray-100 group-hover:text-white" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" aria-hidden="true">
                  <path stroke="none" d="M0 0h24v24H0z"/>  <polyline points="21 12 17 12 14 20 10 4 7 12 3 12" />
                </svg>
              </div>
              <div class="flex-auto">
                <a on:click={close_services} href="/website-maintenance" class="block font-semibold text-gray-900">
                  Website Maintenance
                  <span class="absolute inset-0"></span>
                </a>
                <p class="mt-1 text-gray-600">Keep your website running smoothly.</p>
              </div>
            </div>
          </div>
          <div class="grid grid-cols-2 divide-x divide-gray-900/5 bg-gray-50">
            <a href="tel:+14068719875" class="flex items-center justify-center gap-x-2.5 p-3 text-sm font-semibold leading-6 text-gray-900 hover:bg-gray-100">
              <svg class="h-5 w-5 flex-none text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M2 3.5A1.5 1.5 0 013.5 2h1.148a1.5 1.5 0 011.465 1.175l.716 3.223a1.5 1.5 0 01-1.052 1.767l-.933.267c-.41.117-.643.555-.48.95a11.542 11.542 0 006.254 6.254c.395.163.833-.07.95-.48l.267-.933a1.5 1.5 0 011.767-1.052l3.223.716A1.5 1.5 0 0118 15.352V16.5a1.5 1.5 0 01-1.5 1.5H15c-1.149 0-2.263-.15-3.326-.43A13.022 13.022 0 012.43 8.326 13.019 13.019 0 012 5V3.5z" clip-rule="evenodd" />
              </svg>
              Call us
            </a>
            <a href="/contact-us" class="flex items-center justify-center gap-x-2.5 p-3 text-sm font-semibold leading-6 text-gray-900 hover:bg-gray-100">
              <svg class="h-5 w-5 flex-none text-gray-400" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
                <path stroke="none" d="M0 0h24v24H0z"/>  <line x1="10" y1="14" x2="21" y2="3" />  <path d="M21 3L14.5 21a.55 .55 0 0 1 -1 0L10 14L3 10.5a.55 .55 0 0 1 0 -1L21 3" />
              </svg>
              Contact us
            </a>
          </div>
        </div>
      </div>
      <a href="/pricing" class="text-sm font-semibold leading-6 text-gray-900"
        >Pricing</a
      >
      <a href="/portfolio" class="text-sm font-semibold leading-6 text-gray-900"
        >Portfolio</a
      >
      <a href="/posts" class="text-sm font-semibold leading-6 text-gray-900"
        >Posts</a
      >
      <a
        href="/contact-us"
        class="rounded bg-primary-600 px-2 py-1 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
        >Contact Us</a
      >
    </div>
    <!-- <div class="hidden lg:flex lg:flex-1 lg:justify-end">
			<a href="#" class="text-sm font-semibold leading-6 text-gray-900"
				>Log in <span aria-hidden="true">&rarr;</span></a
			>
		</div> -->
  </nav>
  <!-- Mobile menu, show/hide based on menu open state. -->
  {#if is_mobile_open}
    <div class="lg:hidden" role="dialog" aria-modal="true">
      <!-- Background backdrop, show/hide based on slide-over state. -->
      <div
        transition:fade={fade_backdrop}
        class="fixed inset-0 z-50 bg-gray-950/20"
      />
      <div
        transition:fly={fly_in}
        class="fixed inset-y-0 right-0 z-50 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10"
      >
        <div class="flex items-center justify-between">
          <a on:click={close} href="/" class="-m-1.5 p-1.5">
            <span class="sr-only">Firefly Software</span>
            <Img
              class="h14 w-14 object-cover"
              src={small_logo}
              alt="Firefly Software"
            />
          </a>
          <button
            type="button"
            on:click={() => (is_mobile_open = false)}
            class="-m-2.5 rounded-md p-2.5 text-gray-700"
          >
            <span class="sr-only">Close menu</span>
            <svg
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              aria-hidden="true"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
        <div class="mt-6 flow-root">
          <div class="-my-6 divide-y divide-gray-500/10">
            <div class="space-y-2 py-6">
              <a
                on:click={close}
                href="/website-development"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Website development</a
              >
              <a
                on:click={close}
                href="/website-maintenance"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Website maintenance</a
              >
              <a
                on:click={close}
                href="/pricing"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Pricing</a
              >
              <a
                on:click={close}
                href="/portfolio"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Portfolio</a
              >
              <a
                on:click={close}
                href="/posts"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Posts</a
              >
              <div class="py-6 border-t-2">
                <a
                  on:click={close}
                  href="/contact-us"
                  class="rounded-md bg-primary-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
                  >Contact us</a
                >
              </div>
              <div class="py-6 flex w-full justify-center gap-4">
                <Socials {...socials} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  {/if}
</header>
