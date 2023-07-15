<script lang="ts">
  import fts_service_hardscape from "$lib/assets/images/fts_service_hardscape.png?run";
  import fts_landing from "$lib/assets/images/fts_landing.png?run";
  import { FxReveal as Img } from "@zerodevx/svelte-img";
  import PageIntro from "$lib/components/PageIntro.svelte";

  const intro = {
    eyebrow: "Case Study",
    title: "Kagen Coffee & Crepes",
  };
</script>

<div class="mt-28">
  <div class="mx-auto mt-16 max-w-7xl px-6 sm:mt-32 lg:px-8">
    <!-- Page Intro -->
    <PageIntro eyebrow="Case Study" title="Kagen Coffe & Crepes">
      <p
        class="mx-auto mt-6 max-w-2xl text-center text-lg leading-8 text-gray-600"
      >
        Experience how handcrafted and optimized code results in lightning fast
        load times and a great user experience.
      </p>
    </PageIntro>

    <div
      class="bg-gradient-to-br from-[#efc5a4] to-[#D66E1C] py-16 flex justify-center gap-6"
    >
      <div class="">
        <svg
          viewBox="0 0 366 729"
          role="img"
          class="mx-auto w-[22.875rem] max-w-full drop-shadow-xl"
        >
          <title>App screenshot</title>
          <defs>
            <clipPath id="2ade4387-9c63-4fc4-b754-10e687a0d332">
              <rect width="316" height="684" rx="36" />
            </clipPath>
          </defs>
          <path
            fill="#4B5563"
            d="M363.315 64.213C363.315 22.99 341.312 1 300.092 1H66.751C25.53 1 3.528 22.99 3.528 64.213v44.68l-.857.143A2 2 0 0 0 1 111.009v24.611a2 2 0 0 0 1.671 1.973l.95.158a2.26 2.26 0 0 1-.093.236v26.173c.212.1.398.296.541.643l-1.398.233A2 2 0 0 0 1 167.009v47.611a2 2 0 0 0 1.671 1.973l1.368.228c-.139.319-.314.533-.511.653v16.637c.221.104.414.313.56.689l-1.417.236A2 2 0 0 0 1 237.009v47.611a2 2 0 0 0 1.671 1.973l1.347.225c-.135.294-.302.493-.49.607v377.681c0 41.213 22 63.208 63.223 63.208h95.074c.947-.504 2.717-.843 4.745-.843l.141.001h.194l.086-.001 33.704.005c1.849.043 3.442.37 4.323.838h95.074c41.222 0 63.223-21.999 63.223-63.212v-394.63c-.259-.275-.48-.796-.63-1.47l-.011-.133 1.655-.276A2 2 0 0 0 366 266.62v-77.611a2 2 0 0 0-1.671-1.973l-1.712-.285c.148-.839.396-1.491.698-1.811V64.213Z"
          />
          <path
            fill="#343E4E"
            d="M16 59c0-23.748 19.252-43 43-43h246c23.748 0 43 19.252 43 43v615c0 23.196-18.804 42-42 42H58c-23.196 0-42-18.804-42-42V59Z"
          />
          <foreignObject
            width="316"
            height="684"
            transform="translate(24 24)"
            clip-path="url(#2ade4387-9c63-4fc4-b754-10e687a0d332)"
          >
            <Img
              src={fts_landing}
              alt="Kagen Coffee and Crepes mobile website screen capture"
            />
          </foreignObject>
        </svg>
      </div>
      <div class="">
        <svg
          viewBox="0 0 366 729"
          role="img"
          class="mx-auto w-[22.875rem] max-w-full drop-shadow-xl"
        >
          <title>App screenshot</title>
          <defs>
            <clipPath id="2ade4387-9c63-4fc4-b754-10e687a0d332">
              <rect width="316" height="684" rx="36" />
            </clipPath>
          </defs>
          <path
            fill="#4B5563"
            d="M363.315 64.213C363.315 22.99 341.312 1 300.092 1H66.751C25.53 1 3.528 22.99 3.528 64.213v44.68l-.857.143A2 2 0 0 0 1 111.009v24.611a2 2 0 0 0 1.671 1.973l.95.158a2.26 2.26 0 0 1-.093.236v26.173c.212.1.398.296.541.643l-1.398.233A2 2 0 0 0 1 167.009v47.611a2 2 0 0 0 1.671 1.973l1.368.228c-.139.319-.314.533-.511.653v16.637c.221.104.414.313.56.689l-1.417.236A2 2 0 0 0 1 237.009v47.611a2 2 0 0 0 1.671 1.973l1.347.225c-.135.294-.302.493-.49.607v377.681c0 41.213 22 63.208 63.223 63.208h95.074c.947-.504 2.717-.843 4.745-.843l.141.001h.194l.086-.001 33.704.005c1.849.043 3.442.37 4.323.838h95.074c41.222 0 63.223-21.999 63.223-63.212v-394.63c-.259-.275-.48-.796-.63-1.47l-.011-.133 1.655-.276A2 2 0 0 0 366 266.62v-77.611a2 2 0 0 0-1.671-1.973l-1.712-.285c.148-.839.396-1.491.698-1.811V64.213Z"
          />
          <path
            fill="#343E4E"
            d="M16 59c0-23.748 19.252-43 43-43h246c23.748 0 43 19.252 43 43v615c0 23.196-18.804 42-42 42H58c-23.196 0-42-18.804-42-42V59Z"
          />
          <foreignObject
            width="316"
            height="684"
            transform="translate(24 24)"
            clip-path="url(#2ade4387-9c63-4fc4-b754-10e687a0d332)"
          >
            <Img
              src={fts_service_hardscape}
              alt="Kagen Coffee and Crepes mobile website screen capture"
            />
          </foreignObject>
        </svg>
      </div>
    </div>
  </div>

  <!-- Testimonial section -->
  <div class="mx-auto mt-24 max-w-7xl sm:mt-56 sm:px-6 lg:px-8">
    <div
      class="relative overflow-hidden bg-gray-900 px-6 py-20 shadow-xl sm:rounded-3xl sm:px-10 sm:py-24 md:px-12 lg:px-20"
    >
      <img
        class="absolute inset-0 h-full w-full object-cover brightness-150 saturate-0"
        src="https://images.unsplash.com/photo-1601381718415-a05fb0a261f3?ixid=MXwxMjA3fDB8MHxwcm9maWxlLXBhZ2V8ODl8fHxlbnwwfHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1216&q=80"
        alt=""
      />
      <div class="absolute inset-0 bg-gray-900/90 mix-blend-multiply" />
      <div
        class="absolute -left-80 -top-56 transform-gpu blur-3xl"
        aria-hidden="true"
      >
        <div
          class="aspect-[1097/845] w-[68.5625rem] bg-gradient-to-r from-[#ff4694] to-[#776fff] opacity-[0.45]"
          style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        />
      </div>
      <div
        class="hidden md:absolute md:bottom-16 md:left-[50rem] md:block md:transform-gpu md:blur-3xl"
        aria-hidden="true"
      >
        <div
          class="aspect-[1097/845] w-[68.5625rem] bg-gradient-to-r from-[#ff4694] to-[#776fff] opacity-25"
          style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        />
      </div>
      <div class="relative mx-auto max-w-2xl lg:mx-0">
        <img
          class="h-12 w-auto"
          src="https://tailwindui.com/img/logos/workcation-logo-white.svg"
          alt=""
        />
        <figure>
          <blockquote
            class="mt-6 text-lg font-semibold text-white sm:text-xl sm:leading-8"
          >
            <p>
              “Lorem ipsum dolor sit amet consectetur adipisicing elit. Nemo
              expedita voluptas culpa sapiente alias molestiae. Numquam corrupti
              in laborum sed rerum et corporis.”
            </p>
          </blockquote>
          <figcaption class="mt-6 text-base text-white">
            <div class="font-semibold">Judith Black</div>
            <div class="mt-1">CEO of Workcation</div>
          </figcaption>
        </figure>
      </div>
    </div>
  </div>

  <!-- FAQ section -->
  <div class="mx-auto my-24 max-w-7xl px-6 sm:my-56 lg:px-8">
    <div class="mx-auto max-w-4xl divide-y divide-gray-900/10">
      <h2 class="text-2xl font-bold leading-10 tracking-tight text-gray-900">
        Frequently asked questions
      </h2>
      <dl class="mt-10 space-y-6 divide-y divide-gray-900/10">
        <div class="pt-6">
          <dt>
            <!-- Expand/collapse question button -->
            <button
              type="button"
              class="flex w-full items-start justify-between text-left text-gray-900"
              aria-controls="faq-0"
              aria-expanded="false"
            >
              <span class="text-base font-semibold leading-7"
                >What&#039;s the best thing about Switzerland?</span
              >
              <span class="ml-6 flex h-7 items-center">
                <!--
                    Icon when question is collapsed.

                    Item expanded: "hidden", Item collapsed: ""
                  -->
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
                    d="M12 6v12m6-6H6"
                  />
                </svg>
                <!--
                    Icon when question is expanded.

                    Item expanded: "", Item collapsed: "hidden"
                  -->
                <svg
                  class="hidden h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M18 12H6"
                  />
                </svg>
              </span>
            </button>
          </dt>
          <dd class="mt-2 pr-12" id="faq-0">
            <p class="text-base leading-7 text-gray-600">
              I don&#039;t know, but the flag is a big plus. Lorem ipsum dolor
              sit amet consectetur adipisicing elit. Quas cupiditate laboriosam
              fugiat.
            </p>
          </dd>
        </div>

        <!-- More questions... -->
      </dl>
    </div>
  </div>
</div>
