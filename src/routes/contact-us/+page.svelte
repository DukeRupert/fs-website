<script lang="ts">
  import firefly_jar from "$lib/assets/images/firefly_jar.jpeg?as=run";
  import Img from "@zerodevx/svelte-img";
  import Seo from "$lib/components/SEO.svelte";

  const seoData = {
    title: "Contact Us | Firefly Software",
    description:
      "Excited to leave behind tedious page builders and create the handcrafted web experience your small business deserves? We are!",
    url: "https://www.fireflysoftware.dev/contact-us",
    og: {
      src: "https://www.fireflysoftware.dev/images/facebook_link_image.png",
      alt: "Firefly Software landing page",
      mimeType: "image/png",
      width: 1200,
      height: 648,
    },
  };
</script>

<svelte:head>
  <script src="https://web3forms.com/client/script.js" async defer></script>
</svelte:head>

<Seo data={seoData} />
<section id="contact-us" class="relative bg-white">
  <div class="mt-24 lg:mt-0 lg:absolute lg:inset-0 lg:left-1/2">
    <Img
      class="h-64 w-full bg-gray-50 object-cover sm:h-80 lg:absolute lg:h-full lg:rounded-l-md"
      src={firefly_jar}
      alt="Fireflies captures in a jar on a dark wood background at night"
    />
  </div>
  <div
    class="pb-24 pt-16 sm:pb-32 sm:pt-24 lg:mx-auto lg:grid lg:max-w-7xl lg:grid-cols-2 lg:pt-32"
  >
    <div class="px-6 lg:px-8">
      <div class="mx-auto max-w-xl lg:mx-0 lg:max-w-lg">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900">
          Let's work <span class="gradient-heading">together</span>
        </h2>
        <p class="mt-2 text-lg leading-8 text-gray-600">
          Excited to leave behind tedious page builders and create the <span
            class="text-primary-600">handcrafted</span
          > web experience your small business deserves? We are!
        </p>

        <form
          action="https://api.web3forms.com/submit"
          method="POST"
          class="mt-16"
        >
          <input
            type="hidden"
            name="access_key"
            value="9d7cb285-459d-4211-81a2-642691f2cb33"
          />
          <input type="hidden" name="subject" value="Form submission" />
          <input type="hidden" name="from_name" value="fireflysoftware.dev" />
          <input
            type="hidden"
            name="redirect"
            value="https://fireflysoftware.dev/success"
          />
          <input
            type="checkbox"
            name="botcheck"
            class="hidden"
            style="display: none;"
          />
          <div class="grid grid-cols-1 gap-x-8 gap-y-6 sm:grid-cols-2">
            <div>
              <label
                for="first"
                class="block text-sm font-semibold leading-6 text-gray-900"
                >First name</label
              >
              <div class="mt-2.5">
                <input
                  type="text"
                  name="first"
                  id="first"
                  autocomplete="given-name"
                  placeholder="Malcom"
                  class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
            <div>
              <label
                for="last"
                class="block text-sm font-semibold leading-6 text-gray-900"
                >Last name</label
              >
              <div class="mt-2.5">
                <input
                  type="text"
                  name="last"
                  id="last"
                  autocomplete="family-name"
                  placeholder="Reynolds"
                  class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
            <div class="sm:col-span-2">
              <label
                for="email"
                class="block text-sm font-semibold leading-6 text-gray-900"
                >Email</label
              >
              <div class="mt-2.5">
                <input
                  id="email"
                  name="email"
                  type="email"
                  autocomplete="email"
                  placeholder="browncoatforever@serenity.com"
                  class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
            <div>
              <label
                for="company"
                class="block text-sm font-semibold leading-6 text-gray-900"
                >Company</label
              >
              <div class="mt-2.5">
                <input
                  type="text"
                  name="company"
                  id="company"
                  autocomplete="organization"
                  placeholder="Serenity Shipping"
                  class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
            <div>
              <div class="flex justify-between text-sm leading-6">
                <label for="phone" class="block font-semibold text-gray-900"
                  >Phone</label
                >
                <p id="phone-description" class="text-gray-400">Optional</p>
              </div>
              <div class="mt-2.5">
                <input
                  type="tel"
                  name="phone"
                  id="phone"
                  autocomplete="tel"
                  aria-describedby="phone-description"
                  placeholder="123-456-7890"
                  class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
            <label
              for="service"
              class="block text-sm font-medium leading-6 text-gray-900"
              >Which service are you interested in?</label
            >
            <div class="sm:col-span-2">
              <select
                id="service"
                name="service"
                class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
              >
                <option selected>Standard Marketing</option>
                <option>Advanced Marketing</option>
                <option>E-commerce</option>
                <option>Other</option>
              </select>
            </div>
            <div class="sm:col-span-2">
              <div class="flex justify-between text-sm leading-6">
                <label
                  for="message"
                  class="block text-sm font-semibold leading-6 text-gray-900"
                  >Anything else you would like to tell us?</label
                >
              </div>
              <div class="mt-2.5">
                <textarea
                  id="message"
                  name="message"
                  rows="4"
                  aria-describedby="message-description"
                  placeholder="Please include the url of your current website."
                  class="block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-primary-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
          </div>
          <div class="mt-10 flex justify-end border-t border-gray-900/10 pt-8">
            <button
              type="submit"
              class="rounded-md bg-primary-600 px-3.5 py-2.5 text-center text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
              >Send message</button
            >
          </div>
        </form>
      </div>
    </div>
  </div>
</section>
