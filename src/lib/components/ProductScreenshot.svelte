<script lang="ts">
    import kcc_landing from "$lib/assets/images/kcc_landing.png?as=run";
    import Img from "@zerodevx/svelte-img";
    import type { Feature, Image } from "./ProductScreenshot";
    export let left = false;
    export let eyebrow = "Deploy faster";
    export let title = "A better workflow";
    export let description = `Lorem ipsum, dolor sit amet consectetur adipisicing
                        elit. Maiores impedit perferendis suscipit eaque, iste
                        dolor cupiditate blanditiis ratione.`;
    export let image: Image = {
        src: kcc_landing,
        alt: "Kagen Coffee and Crepes mobile view screen capture"
    }

    export let features: Feature[] = [
        {
            title: "Push to deploy.",
            description:
                "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis ratione.",
            iconPath: [
                '<path fill-rule="evenodd" d="M5.5 17a4.5 4.5 0 01-1.44-8.765 4.5 4.5 0 018.302-3.046 3.5 3.5 0 014.504 4.272A4 4 0 0115 17H5.5zm3.75-2.75a.75.75 0 001.5 0V9.66l1.95 2.1a.75.75 0 101.1-1.02l-3.25-3.5a.75.75 0 00-1.1 0l-3.25 3.5a.75.75 0 101.1 1.02l1.95-2.1v4.59z" clip-rule="evenodd" />',
            ],
        },
        {
            title: "SSL certificates.",
            description:
                "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis ratione.",
            iconPath: [
                '<path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" />',
            ],
        },
        {
            title: "Database backups.",
            description:
                "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Maiores impedit perferendis suscipit eaque, iste dolor cupiditate blanditiis ratione.",
            iconPath: [
                '<path d="M4.632 3.533A2 2 0 016.577 2h6.846a2 2 0 011.945 1.533l1.976 8.234A3.489 3.489 0 0016 11.5H4c-.476 0-.93.095-1.344.267l1.976-8.234z" />',
                '<path fill-rule="evenodd" d="M4 13a2 2 0 100 4h12a2 2 0 100-4H4zm11.24 2a.75.75 0 01.75-.75H16a.75.75 0 01.75.75v.01a.75.75 0 01-.75.75h-.01a.75.75 0 01-.75-.75V15zm-2.25-.75a.75.75 0 00-.75.75v.01c0 .414.336.75.75.75H13a.75.75 0 00.75-.75V15a.75.75 0 00-.75-.75h-.01z" clip-rule="evenodd" />',
            ],
        },
    ];
</script>

<div class="overflow-hidden bg-white py-24 sm:py-32">
    <div class="mx-auto max-w-7xl px-6 lg:px-8">
        <div
            class="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 sm:gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-2"
        >
            <div class="lg:pr-8 lg:pt-4">
                <div class="lg:max-w-lg">
                    <h2
                        class="text-base font-semibold leading-7 text-primary-600"
                    >
                        {eyebrow}
                    </h2>
                    <p
                        class="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl"
                    >
                        {title}
                    </p>
                    <p class="mt-6 text-lg leading-8 text-gray-600">
                        {description}
                    </p>
                    <dl
                        class="mt-10 max-w-xl space-y-8 text-base leading-7 text-gray-600 lg:max-w-none"
                    >
                        {#each features as { title, description, iconPath }}
                            <div class="relative pl-9">
                                <dt class="inline font-semibold text-gray-900">
                                    <svg
                                        class="absolute left-1 top-1 h-5 w-5 text-primary-600"
                                        width="24"
                                        height="24"
                                        viewBox="0 0 24 24"
                                        stroke-width="2"
                                        stroke="currentColor"
                                        fill="none"
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                    >
                                        {#each iconPath as p}
                                            {@html p}
                                        {/each}
                                    </svg>
                                    {title}
                                </dt>
                                <dd class="inline">
                                    {description}
                                </dd>
                            </div>
                        {/each}
                    </dl>
                </div>
            </div>
            {#if left}
                <div
                    class="mt-16 sm:mt-24 lg:mt-0 lg:flex-shrink-0 lg:flex-grow"
                >
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
                                src={image.src}
                                alt={image.alt}
                                loading="eager"
                                decoding="sync"
                            />
                        </foreignObject>
                    </svg>
                </div>
            {:else}
                <div class="flex items-start justify-center lg:order-first">
                    <div
                        class="mt-16 sm:mt-24 lg:mt-0 lg:flex-shrink-0 lg:flex-grow"
                    >
                        <svg
                            viewBox="0 0 366 729"
                            role="img"
                            class="mx-auto w-[22.875rem] max-w-full drop-shadow-xl"
                        >
                            <title>App screenshot</title>
                            <defs>
                                <clipPath
                                    id="2ade4387-9c63-4fc4-b754-10e687a0d332"
                                >
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
                                    src={image.src}
                                    alt={image.alt}
                                    loading="eager"
                                    decoding="sync"
                                />
                            </foreignObject>
                        </svg>
                    </div>
                </div>
            {/if}
        </div>
    </div>
</div>
