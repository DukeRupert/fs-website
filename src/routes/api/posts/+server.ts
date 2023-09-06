import type { Post } from "$lib/types/sanity";
import type { RequestHandler, RequestEvent } from "./$types.js";
import client from "$lib/db";
import { json } from "@sveltejs/kit";

interface Query {
  posts: Post[];
  count: number;
}

export async function GET({ url }: RequestEvent) {
  const limit = Number(url.searchParams.get("limit") ?? 5);
  const page = Number(url.searchParams.get("page") ?? 0);
  const category = url.searchParams.get("category") ?? "";

  const projection = `{
			..., categories[]->{title}, author->{name, role, image}, body[]{
				..., 
				image{
					..., asset->
				}
			}
		}`;

  let filter = `*[_type == "post" && defined(slug)] | order(_createdAt desc)[$skip...$size]`;
  if (category !== "")
    filter = `*[_type == "post" && defined(slug) && $category in categories[]->title] | order(_createdAt desc)[$skip...$size]`;

  let count = `count(*[_type == "post" && defined(slug)])`;
  if (category !== "") {
    count = `count(*[_type == "post" && defined(slug) && $category in categories[]->title])`;
  }

  const params = {
    skip: page * limit,
    size: page * limit + limit,
    category,
  };

  const query = `{"posts" : ${filter + projection}, "count": ${count}}`;
  const data = (await client.fetch(query, params)) as Query;
  return json(data);
}
