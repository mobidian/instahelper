package handlers

import (
	"html/template"
	"io"

	"github.com/socialplanner/instahelper/app/assets"
)

var funcs = template.FuncMap{
	"notifications": func() []Notification {
		// TODO. Replace with a method to actually get notifications
		return []Notification{
			{
				Text: "Test Notification",
				Link: "https://twitter.com/spieswithin",
			},
			{
				Text: "Test Notification 2",
				Link: "https://twitter.com/spieswithin",
			},
		}
	},
}

// Template will load the corresponding template with presets.
func Template(templateName string) *Page {
	if page, ok := Pages[templateName]; ok {
		return &page
	}
	return nil
}

var a = assets.MustAsset

// Creates a template with the default funcs. Panics on error.
func newTemplate(files ...string) *template.Template {
	tmpl := template.New("*").Funcs(funcs)

	for _, file := range files {
		// assets.Asset defaults to '/' as a separator
		file = "templates/" + file

		content := string(a(file))
		tmpl = template.Must(tmpl.Parse(content))
	}
	return tmpl
}

// Execute is shorthand for Page.Template.Execute(w, Page, data)
func (p *Page) Execute(w io.Writer, data ...map[string]interface{}) error {
	templateData := map[string]interface{}{
		"Pages": SortPages(Pages),
		"Title": p.Name,
		"Icon":  p.Icon,
		"Link":  p.Link,
	}

	if len(data) > 0 {
		for key, val := range data[0] {
			templateData[key] = val
		}
	}

	return p.Template.Execute(
		w,
		templateData,
	)
}
