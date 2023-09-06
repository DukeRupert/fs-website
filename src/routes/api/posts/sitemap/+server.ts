import { error, json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import client from '$lib/db';

interface Post {
	title: string;
	slug: string;
	updatedAt: string;
}

function getDate(raw: string) {
	const split = raw.split('T');
	return split[0];
}

export const GET: RequestHandler = async () => {
	const filter = `*[_type == "post" && defined(slug)]`;
	const projection = `{title, "slug" : slug.current, "updatedAt": _updatedAt}`;
	const query = filter + projection;

	const posts: Post[] = await client.fetch(query);

	if (posts.length == 0) {
		throw error(404);
	}

	// format date
	const data = posts.map((el) => Object.assign({}, el, { updatedAt: getDate(el.updatedAt) }));

	return json(data);
};