import type { PageLoad } from "./$types";
import type { Post } from "$lib/types/sanity";
import { error } from "@sveltejs/kit";

export const load: PageLoad = async ({ params, fetch }) => {
  const { slug } = params;
  const res = await fetch(`/api/posts/${slug}`);
  if (!res.ok) {
    throw error(404, { message: "An error has occured" });
  }
  const post: Post = await res.json();
  return {
    post,
  };
};
