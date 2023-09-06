import type { PageLoad } from "./$types";
import type { Post } from "$lib/types/sanity";
interface Query {
  posts: Post[];
  count: number;
}

export const load: PageLoad = async ({ url, fetch }) => {
  // Get params
  const order = url.searchParams.get("order") ?? "desc";
  const limit = url.searchParams.get("limit") ?? "5";
  const page = url.searchParams.get("page") ?? "0";
  const category = url.searchParams.get("category") ?? "";

  // Build query
  const base = url.origin;
  let query = new URL("/api/posts", base);
  query.searchParams.append("order", order);
  query.searchParams.append("limit", limit);
  query.searchParams.append("page", page);
  query.searchParams.append("category", category);

  const response = await fetch(query.href);

  const data: Query = await response.json();
  const { posts, count } = data;

  return {
    posts,
    count,
  };
};
