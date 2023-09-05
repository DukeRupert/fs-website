import { createClient } from "@sanity/client";
import imageUrlBuilder from "@sanity/image-url";
import type { SanityAsset } from "@sanity/image-url/lib/types/types";

const client = createClient({
  projectId: "vpzagt04",
  dataset: "production",
  apiVersion: "2023-09-04",
  useCdn: false,
});

const builder = imageUrlBuilder(client);

export function urlFor(source: string | SanityAsset) {
  return builder.image(source);
}

export default client;
