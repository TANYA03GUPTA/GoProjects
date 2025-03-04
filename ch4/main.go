package main

import (
	"html/template"
	"net/http"
	"time"
)

type Page struct {
    Title   string
    Content string
    Date    string
}
func main() {
    http.HandleFunc("/", ServePage)
    http.ListenAndServe(":8080", nil)
}

func NewPage(title, content string) *Page {
    return &Page{
        Title:   title,
        Content: content,
        Date:    time.Now().Format("January 2, 2006"),
    }
}
func ServePage(w http.ResponseWriter, r *http.Request) {
    page := NewPage("My Blog Title", "This is the content of the blog.")

    tmpl, err := template.ParseFiles("templates/blog.html")
    if err != nil {
        http.Error(w, "Unable to load template", http.StatusInternalServerError)
        return
    }

    // Execute the template with the page data
    err = tmpl.Execute(w, page)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
}