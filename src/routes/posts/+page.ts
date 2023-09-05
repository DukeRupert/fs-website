import type { PageLoad } from "./$types";
import type { Post } from "$lib/types/sanity";
import client from "$lib/db";
export const load: PageLoad = async () => {
  const filter =
    '*[_type == "post" && publishedAt < now()] | order(publishedAt desc)';
  const projection = `
		{ 
			...,
			categories[]->{...},
			mainImage{..., asset->},
	  	}`;
  const query = filter + projection;
  const data: Post[] = await client.fetch(query);

  return {
    posts: data,
  };
};
