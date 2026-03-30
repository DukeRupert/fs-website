package handlers

import (
	"net/http"
)

// ProjectQuote holds a client testimonial.
type ProjectQuote struct {
	Text        string
	Attribution string
}

// ProjectResult holds a single metric.
type ProjectResult struct {
	Value string
	Unit  string
	Label string
}

// Project holds all data for a portfolio page.
type Project struct {
	Slug             string
	Client           string
	Title            string
	Subtitle         string
	Location         string
	Tags             []string
	LiveURL          string
	Screenshot       string
	ScreenshotMobile string
	Story            []string
	Quote            *ProjectQuote
	Results          []ProjectResult
	Deliverables     []string
	NextSlug         string
	NextClient       string
}

// PortfolioData extends PageData with project-specific fields.
type PortfolioData struct {
	PageData
	Project Project
}

// projects is the canonical list of portfolio entries.
var projects = map[string]Project{
	"nautilus-group": {
		Slug:     "nautilus-group",
		Client:   "Nautilus Group Cleaning Services",
		Title:    "No web presence. Now they have a pipeline.",
		Subtitle: "Custom site, brand, and SEO for a cleaning company starting from zero.",
		Location: "Commercial cleaning · Kennewick, WA",
		Tags:     []string{"Marketing site", "Brand", "SEO"},
		LiveURL:          "https://nautiluscleaning.com",
		Screenshot:       "/static/images/portfolio/nautilus_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/nautilus_screenshot_mobile.webp",
		NextSlug:         "a-team-asphalt",
		NextClient:       "A-Team Asphalt",
		Story: []string{
			"Nautilus had nothing — no site, no listings, no digital footprint. A husband-and-wife team with real credentials and real work ethic, but completely invisible online. They needed everything: a brand identity, a website, a Google Business Profile, and an SEO foundation that would start generating leads.",
			"We built it all from scratch. The site was designed to communicate trustworthiness and professionalism — insured, bonded, licensed — without looking corporate. The copy speaks directly to property managers and business owners who need reliable commercial cleaning.",
			"Within four weeks of launch they had received a commercial contract and five new inquiries. They had nothing before. Now they have a pipeline.",
		},
		Quote: &ProjectQuote{
			Text:        "We went from invisible to getting calls. That's what we needed.",
			Attribution: "Nautilus Group Cleaning Services",
		},
		Results: []ProjectResult{
			{Value: "0", Unit: "→1", Label: "Web presence built from nothing"},
			{Value: "4", Unit: "wks", Label: "Launch to first commercial contract"},
			{Value: "5", Unit: "leads", Label: "Inquiries in month one"},
		},
		Deliverables: []string{
			"Custom-designed marketing site — no templates, no page builders",
			"Brand identity — logo, color palette, typography",
			"Full copywriting — every page written by Firefly",
			"Google Business Profile setup and optimization",
			"SEO foundation — meta tags, schema markup, sitemap",
			"Hosting and SSL on Firefly infrastructure",
			"Ongoing maintenance and content updates",
		},
	},
	"a-team-asphalt": {
		Slug:     "a-team-asphalt",
		Client:   "A-Team Asphalt",
		Title:    "Twenty years of reputation. A website that finally matched it.",
		Subtitle: "Full custom site with brand identity, copy, and SEO for a Pacific Northwest paving company.",
		Location: "Asphalt & paving · Sumner, WA",
		Tags:     []string{"Marketing site", "SEO"},
		LiveURL:          "https://ateamasphalt.com",
		Screenshot:       "/static/images/portfolio/ateam_screenshot.webp",
		ScreenshotMobile: "/static/images/portfolio/ateam_screenshot_mobile.webp",
		NextSlug:         "a-team-gutters",
		NextClient:       "A-Team Gutters",
		Story: []string{
			"Don and Gary had built a real business the hard way — word of mouth, quality work, repeat clients across the Puyallup Valley. Their website hadn't kept up. It didn't reflect the scale of the operation or the quality of the work.",
			"We rebuilt it from the ground up — brand, copy, and code — and fixed an NAP inconsistency that was quietly suppressing their local search ranking. The new site was designed to convert: clear service pages, strong calls to action, and copy that speaks to both commercial property managers and residential homeowners.",
			"The site launched in six weeks. Every line of code was written for them. Every word of copy was written about their specific business, their specific market, their specific strengths.",
		},
		Quote: &ProjectQuote{
			Text:        "They understood what we do and built something we're actually proud to hand out.",
			Attribution: "Don — A-Team Asphalt",
		},
		Results: []ProjectResult{
			{Value: "6", Unit: "wks", Label: "Discovery to launch"},
			{Value: "0", Unit: "templates", Label: "Every line custom"},
			{Value: "1", Unit: "call", Label: "To reach the builder"},
		},
		Deliverables: []string{
			"Custom-designed marketing site — dark, industrial aesthetic matching the trade",
			"Brand identity — logo refinement, color system, typography",
			"Full copywriting — service pages, about, contact",
			"SEO strategy — NAP correction, local search optimization",
			"Google Business Profile setup",
			"Email infrastructure",
			"Hosting, SSL, and ongoing maintenance",
		},
	},
	"a-team-gutters": {
		Slug:     "a-team-gutters",
		Client:   "A-Team Gutters",
		Title:    "New trade, new brand, built to generate leads from day one.",
		Subtitle: "Brand identity and custom site for a new gutter installation company.",
		Location: "Gutter installation · Bonney Lake, WA",
		Tags:       []string{"Marketing site", "Brand"},
		NextSlug:   "western-skies",
		NextClient: "Western Skies Contracting",
		Story: []string{
			"A-Team Gutters launched in 2024 with strong trade credentials and no web presence. They needed a brand and a site that would generate residential leads in a competitive Pacific Northwest market from day one.",
			"We built the brand from the ground up — identity, copy, and a custom site designed around the specific services they offer and the specific areas they serve. No templates, no generic contractor design. A site built for this business, in this market.",
		},
		Results: []ProjectResult{
			{Value: "0", Unit: "→1", Label: "Brand created from scratch"},
			{Value: "1", Unit: "site", Label: "Custom built"},
			{Value: "2024", Unit: "", Label: "Launch year"},
		},
		Deliverables: []string{
			"Brand identity — logo, colors, typography",
			"Custom-designed marketing site",
			"Copywriting — all pages",
			"SEO foundation",
			"Hosting and SSL",
		},
	},
	"western-skies": {
		Slug:     "western-skies",
		Client:   "Western Skies Contracting",
		Title:    "A quoting tool built around how a contractor actually thinks.",
		Subtitle: "Custom estimating software replacing spreadsheets with a purpose-built workflow.",
		Location: "Construction · Hamilton, MT",
		Tags:       []string{"Web app", "Estimating"},
		NextSlug:   "rockabilly-roasting",
		NextClient: "Rockabilly Roasting",
		Story: []string{
			"Western Skies was running estimates out of spreadsheets — workable, but slow and error-prone. Generic estimating software meant paying for features that didn't fit and working around the ones that did.",
			"Skalkaho is purpose-built for Western Skies — their line items, their workflow, their clients. Line-item pricing with automated updates from local manufacturer price lists, e-signature support, version history, and a clean interface that gets out of the way. What used to take an afternoon now takes minutes.",
			"No subscription. No bloat. Just the tool they needed, built to spec.",
		},
		Results: []ProjectResult{
			{Value: "2", Unit: "phases", Label: "Milestone delivery"},
			{Value: "$4.5", Unit: "k", Label: "Fixed price, no surprises"},
			{Value: "0", Unit: "bloat", Label: "Built to spec, nothing extra"},
		},
		Deliverables: []string{
			"Custom web application — Go backend, htmx frontend",
			"Estimating workflow — line items, pricing, totals",
			"Automated manufacturer price list integration",
			"E-signature support",
			"Version history and draft management",
			"Client management",
			"Hosted on Firefly infrastructure",
		},
	},
	"rockabilly-roasting": {
		Slug:     "rockabilly-roasting",
		Client:   "Rockabilly Roasting",
		Title:    "Two buying experiences, two pricing structures, one platform.",
		Subtitle: "Custom e-commerce handling retail, wholesale, and subscriptions without a platform dependency.",
		Location: "Coffee roaster · Kennewick, WA",
		Tags:       []string{"E-commerce", "B2B + B2C"},
		NextSlug:   "nautilus-group",
		NextClient: "Nautilus Group",
		Story: []string{
			"Rockabilly needed more than a storefront. They sell retail direct to consumers and wholesale to cafes and restaurants — two different buying experiences, two different pricing structures, one platform.",
			"We built a custom e-commerce solution handling B2C sales, B2B wholesale accounts, and subscriptions. No Shopify, no plugin stack, no monthly ransom to a platform they don't own. The business owns the code, the data, and the customer relationships.",
		},
		Results: []ProjectResult{
			{Value: "2", Unit: "channels", Label: "B2C + wholesale"},
			{Value: "0", Unit: "platforms", Label: "Fully custom owned"},
			{Value: "1", Unit: "codebase", Label: "Both experiences unified"},
		},
		Deliverables: []string{
			"Custom e-commerce platform — Svelte frontend, Go backend",
			"B2C retail storefront with cart and checkout",
			"B2B wholesale portal with account-based pricing",
			"Subscription management",
			"Inventory and order management",
			"Payment processing integration",
		},
	},
}

// PortfolioHandler returns an http.HandlerFunc that serves individual project pages.
func PortfolioHandler(tr *TemplateRenderer, site SiteData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		project, ok := projects[slug]
		if !ok {
			http.NotFound(w, r)
			return
		}

		data := PortfolioData{
			PageData: NewPageData(
				project.Client+" — Firefly Software",
				project.Subtitle,
				site,
			).WithNav("work"),
			Project: project,
		}

		tr.Render(w, r, "portfolio", data)
	}
}
