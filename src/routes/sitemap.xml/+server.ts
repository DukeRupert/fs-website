import type { RequestHandler } from './$types';

interface Post {
	title: string;
	slug: string;
	updatedAt: string;
}

const pages: string[] = ['website-development', 'website-maintenance', 'pricing', 'posts', 'portfolio', 'contact-us']; //list of pages as a string ex. ["about", "blog", "contact"]

const site = 'https://www.fireflysoftware.dev';

export const GET: RequestHandler = async ({ fetch }) => {
	const postsData = await fetch('/api/posts/sitemap');
	const posts: Post[] = await postsData.json();
	const body = sitemap(posts, pages);
	const response = new Response(body);
	response.headers.set('Cache-Control', 'max-age=0, s-maxage=3600');
	response.headers.set('Content-Type', 'application/xml');
	return response;
};

const sitemap = (posts: Post[], pages: string[]) => `<?xml version="1.0" encoding="UTF-8" ?>
<urlset
  xmlns="https://www.sitemaps.org/schemas/sitemap/0.9"
  xmlns:news="https://www.google.com/schemas/sitemap-news/0.9"
  xmlns:xhtml="https://www.w3.org/1999/xhtml"
  xmlns:mobile="https://www.google.com/schemas/sitemap-mobile/1.0"
  xmlns:image="https://www.google.com/schemas/sitemap-image/1.1"
  xmlns:video="https://www.google.com/schemas/sitemap-video/1.1"
>
  <url>
    <loc>${site}</loc>
  </url>
  ${pages
		.map(
			(page) => `
  <url>
    <loc>${site}/${page}</loc>
  </url>
  `
		)
		.join('')}
  ${posts
		.map(
			(post) => `
  <url>
    <loc>${site}/posts/${post.slug}</loc>
    <lastmod>${post.updatedAt}</lastmod>
  </url>
  `
		)
		.join('')}
</urlset>`;