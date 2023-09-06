import { error, json } from "@sveltejs/kit";
import client from "$lib/db";
import type { Post } from "$lib/types/sanity";

export async function GET({ params }) {
  const { slug } = params;
  const filter = `*[_type == "post" && publishedAt < now() && slug.current == "${slug}"] | order(publishedAt desc)[0]`;
  const projection = `
		{ 
			...,
      author->{name, role, image},
			categories[]->{...},
			mainImage{..., asset->},
	  	}`;
  const query = filter + projection;
  const data: Post = await client.fetch(query);

  if (!data) {
    throw error(404, { message: "Page not found" });
  }

  return json(data);
}
