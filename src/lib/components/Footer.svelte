<script lang="ts">
  import logo from "$lib/assets/images/logo/logo_two_full.jpg?run";
  import Img from "@zerodevx/svelte-img";
</script>

<footer class="bg-white" aria-labelledby="footer-heading">
  <h2 id="footer-heading" class="sr-only">Footer</h2>
  <div class="mx-auto max-w-7xl px-6 pb-8 pt-16 sm:pt-24 lg:px-8 lg:pt-32">
    <div class="xl:grid xl:grid-cols-3 xl:gap-8">
      <div class="space-y-8">
        <Img class="h-14 w-auto" src={logo} alt="Firefly Software" />
        <p class="text-sm leading-6 text-gray-600">
          Creating a better online experience through meticulous craftsmanship.
        </p>
      </div>
      <div class="mt-16 grid grid-cols-2 gap-8 xl:col-span-2 xl:mt-0">
        <div class="md:grid md:grid-cols-2 md:gap-8">
          <div>
            <h3 class="text-sm font-semibold leading-6 text-gray-900">
              Navigation
            </h3>
            <ul class="mt-6 space-y-4">
              <li>
                <a
                  href="/"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Home</a
                >
              </li>
              <li>
                <a
                  href="#services"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Services</a
                >
              </li>
              <li>
                <a
                  href="/portfolio"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Portfolio</a
                >
              </li>
              <li>
                <a
                  href="#features"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Features</a
                >
              </li>
              <li>
                <a
                  href="#pricing"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Pricing</a
                >
              </li>
            </ul>
          </div>
          <div class="mt-10 md:mt-0">
            <h3 class="text-sm font-semibold leading-6 text-gray-900">
              Services
            </h3>
            <ul class="mt-6 space-y-4">
              <li>
                <a
                  href="#"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Web Design</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Website Maintenance</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >SEO</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-sm leading-6 text-gray-600 hover:text-gray-900"
                  >Content Creation</a
                >
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <div class="mt-16 border-t border-gray-900/10 pt-8 sm:mt-20 lg:mt-24">
      <p class="text-xs leading-5 text-gray-500">
        &copy; 2020 Firefly Software, LLC. All rights reserved.
      </p>
    </div>
  </div>
</footer>
