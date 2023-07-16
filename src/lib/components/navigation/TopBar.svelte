<script lang="ts">
  import { quintOut } from "svelte/easing";
  import { fly, fade } from "svelte/transition";
  import type { FlyParams, FadeParams } from "svelte/transition";
  import Img from "@zerodevx/svelte-img";
  import small_logo from "$lib/assets/images/logo/logo_small.jpg?run";
  import large_logo from "$lib/assets/images/logo/logo_two_full.jpg?run";

  // Toggle mobile menu
  let is_mobile_open = false;
  function close(): void {
    is_mobile_open = false;
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
      <a href="/#services" class="text-sm font-semibold leading-6 text-gray-900"
        >Services</a
      >
      <a href="/#features" class="text-sm font-semibold leading-6 text-gray-900"
        >Features</a
      >
      <a href="/#pricing" class="text-sm font-semibold leading-6 text-gray-900"
        >Pricing</a
      >
      <a href="/portfolio" class="text-sm font-semibold leading-6 text-gray-900"
        >Portfolio</a
      >
      <a
        href="/#contact-us"
        class="text-sm font-semibold leading-6 text-gray-900">Contact Us</a
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
          <a href="/" class="-m-1.5 p-1.5">
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
                href="/#services"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Services</a
              >
              <a
                on:click={close}
                href="/#features"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Features</a
              >
              <a
                on:click={close}
                href="/#pricing"
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
                href="/#contact-us"
                class="-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50"
                >Contact Us</a
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  {/if}
</header>
