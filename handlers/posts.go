package handlers

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

// Post represents a parsed blog post.
type Post struct {
	Slug        string
	Title       string
	Description string
	Date        string
	Author      string
	Tags        []string
	Content     template.HTML
}

// PostData holds template data for a single post page.
type PostData struct {
	PageData
	Post Post
}

// PostListData holds template data for the post listing page.
type PostListData struct {
	PageData
	Posts []Post
}

// PostHandler serves blog posts from markdown files in the content/posts directory.
type PostHandler struct {
	contentDir string
	renderer   *TemplateRenderer
	site       SiteData
	md         goldmark.Markdown
}

// NewPostHandler creates a handler for blog posts.
func NewPostHandler(contentDir string, renderer *TemplateRenderer, site SiteData) *PostHandler {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	return &PostHandler{
		contentDir: contentDir,
		renderer:   renderer,
		site:       site,
		md:         md,
	}
}

// HandlePost serves a single blog post by slug.
func (ph *PostHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	post, err := ph.loadPost(slug)
	if err != nil {
		log.Printf("[posts] error loading %s: %v", slug, err)
		http.NotFound(w, r)
		return
	}

	data := PostData{
		PageData: NewPageData(post.Title, post.Description, ph.site),
		Post:     post,
	}
	ph.renderer.Render(w, r, "post", data)
}

// HandlePostList serves the blog post listing.
func (ph *PostHandler) HandlePostList(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.loadAllPosts()
	if err != nil {
		log.Printf("[posts] error loading posts: %v", err)
		http.Error(w, "error loading posts", http.StatusInternalServerError)
		return
	}

	data := PostListData{
		PageData: NewPageData("Posts", "Articles about web development, small business, and technology.", ph.site),
		Posts:    posts,
	}
	ph.renderer.Render(w, r, "posts", data)
}

func (ph *PostHandler) loadPost(slug string) (Post, error) {
	filename := filepath.Join(ph.contentDir, slug+".md")
	return ph.parseMarkdownFile(filename, slug)
}

func (ph *PostHandler) loadAllPosts() ([]Post, error) {
	entries, err := os.ReadDir(ph.contentDir)
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") || entry.Name() == "_index.md" {
			continue
		}
		slug := strings.TrimSuffix(entry.Name(), ".md")
		post, err := ph.parseMarkdownFile(filepath.Join(ph.contentDir, entry.Name()), slug)
		if err != nil {
			log.Printf("[posts] skipping %s: %v", entry.Name(), err)
			continue
		}
		posts = append(posts, post)
	}

	// Sort by date descending
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	return posts, nil
}

func (ph *PostHandler) parseMarkdownFile(filename, slug string) (Post, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	post := Post{Slug: slug}

	// Parse YAML-style front matter (between --- delimiters)
	inFrontMatter := false
	var contentLines []string

	for scanner.Scan() {
		line := scanner.Text()

		if !inFrontMatter && strings.TrimSpace(line) == "---" {
			inFrontMatter = true
			continue
		}

		if inFrontMatter {
			if strings.TrimSpace(line) == "---" {
				inFrontMatter = false
				continue
			}
			parseFrontMatterLine(line, &post)
			continue
		}

		contentLines = append(contentLines, line)
	}

	if err := scanner.Err(); err != nil {
		return Post{}, err
	}

	// Render markdown to HTML
	source := []byte(strings.Join(contentLines, "\n"))
	var buf strings.Builder
	if err := ph.md.Convert(source, &buf); err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(buf.String())

	return post, nil
}

func parseFrontMatterLine(line string, post *Post) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	// Strip surrounding quotes
	value = strings.Trim(value, `"'`)

	switch key {
	case "title":
		post.Title = value
	case "description":
		post.Description = value
	case "date":
		post.Date = value
	case "author":
		post.Author = value
	case "tags":
		// Parse ["tag1", "tag2"] format
		value = strings.Trim(value, "[]")
		for _, tag := range strings.Split(value, ",") {
			tag = strings.TrimSpace(tag)
			tag = strings.Trim(tag, `"'`)
			if tag != "" {
				post.Tags = append(post.Tags, tag)
			}
		}
	}
}
