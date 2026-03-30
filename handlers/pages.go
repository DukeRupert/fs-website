package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// TemplateRenderer manages parsed templates and renders pages.
type TemplateRenderer struct {
	templates map[string]*template.Template
	mu        sync.RWMutex
	devMode   bool
	baseDir   string
}

// NewTemplateRenderer parses all page templates at startup.
// Each page template is combined with the shared base, nav, and footer partials.
func NewTemplateRenderer(baseDir string, devMode bool) (*TemplateRenderer, error) {
	tr := &TemplateRenderer{
		templates: make(map[string]*template.Template),
		devMode:   devMode,
		baseDir:   baseDir,
	}

	if err := tr.parseAll(); err != nil {
		return nil, err
	}

	return tr, nil
}

func (tr *TemplateRenderer) parseAll() error {
	shared := []string{
		tr.baseDir + "/base.html",
		tr.baseDir + "/nav.html",
		tr.baseDir + "/footer.html",
	}

	pages := map[string]string{
		"home":               tr.baseDir + "/home.html",
		"work":               tr.baseDir + "/work.html",
		"process":            tr.baseDir + "/process.html",
		"about":              tr.baseDir + "/about.html",
		"pricing":            tr.baseDir + "/pricing.html",
		"privacy":            tr.baseDir + "/privacy.html",
		"contact":            tr.baseDir + "/contact.html",
		"services-websites":  tr.baseDir + "/services-websites.html",
		"services-software":  tr.baseDir + "/services-software.html",
		"post":               tr.baseDir + "/post.html",
		"posts":              tr.baseDir + "/posts.html",
		"404":                tr.baseDir + "/404.html",
	}

	tr.mu.Lock()
	defer tr.mu.Unlock()

	for name, pageFile := range pages {
		files := append([]string{pageFile}, shared...)
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return err
		}
		tr.templates[name] = tmpl
	}

	return nil
}

// Render executes a named template and writes the result to w.
func (tr *TemplateRenderer) Render(w http.ResponseWriter, r *http.Request, name string, data any) {
	// In dev mode, reparse templates on every request for hot reload.
	if tr.devMode {
		if err := tr.parseAll(); err != nil {
			log.Printf("[template] reparse error: %v", err)
			http.Error(w, "template error", http.StatusInternalServerError)
			return
		}
	}

	tr.mu.RLock()
	tmpl, ok := tr.templates[name]
	tr.mu.RUnlock()

	if !ok {
		log.Printf("[template] not found: %s", name)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("[template] render error (%s): %v", name, err)
	}
}

// PageHandler returns an http.HandlerFunc that renders the named template with the given data.
func PageHandler(tr *TemplateRenderer, name string, data PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tr.Render(w, r, name, data)
	}
}
