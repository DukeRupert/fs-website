package handlers

import (
	"net/http"
	"time"
)

const lomoBasePath = "/lo-mo-outfitters"

// LoMoData holds template data for Lo Mo Outfitting pages.
type LoMoData struct {
	Title       string
	Description string
	Slug        string
	BasePath    string
	BodyClass   string
	CSSFile     string
	Year        int
}

func newLoMoData(slug string, cssFile string) LoMoData {
	bodyClass := "bg-lt-deep"
	if slug == "amber" {
		bodyClass = "bg-la-dark"
	}
	return LoMoData{
		Title:       "Home | Lo Mo Outfitting",
		Description: "Missouri River fly fishing out of Craig, Montana. Local guides. No pretense.",
		Slug:        slug,
		BasePath:    lomoBasePath,
		BodyClass:   bodyClass,
		CSSFile:     cssFile,
		Year:        time.Now().Year(),
	}
}

// LoMoLandingHandler serves GET /lo-mo-outfitters — the comparison landing page.
func LoMoLandingHandler(tr *TemplateRenderer, cssFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := LoMoData{
			BasePath: lomoBasePath,
			CSSFile:  cssFile,
		}
		tr.Render(w, r, "lomo-landing", data)
	}
}

// LoMoHomeHandler serves GET /lo-mo-outfitters/{slug}/ — the variation homepage.
func LoMoHomeHandler(tr *TemplateRenderer, cssFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		if slug != "amber" && slug != "teal" {
			http.NotFound(w, r)
			return
		}
		data := newLoMoData(slug, cssFile)
		tr.Render(w, r, "lomo-"+slug, data)
	}
}

// LoMoContactHandler serves GET /lo-mo-outfitters/{slug}/contact.
func LoMoContactHandler(tr *TemplateRenderer, cssFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		if slug != "amber" && slug != "teal" {
			http.NotFound(w, r)
			return
		}
		data := newLoMoData(slug, cssFile)
		data.Title = "Contact | Lo Mo Outfitting"
		data.Description = "Get in touch with Lo Mo Outfitting."
		tr.Render(w, r, "lomo-contact-"+slug, data)
	}
}
