package main

import (
	"html/template"
	"log"
	"main/ascii"
	"net/http"
	"strings"
)

type Result struct {
	Text  string
	Style string
	Ascii string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		w.WriteHeader(404)
		renderTemplate(w, "error", "404: Page not found")
	} else {
		w.WriteHeader(200)
		renderTemplate(w, "index", nil)
	}
}

func artHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	style := r.FormValue("style")

	if len(text) <= 300 {

		// log.Println(text)
		ascii, err := ascii.GenerateAscii(text, style)
		result := &Result{Text: text, Style: style, Ascii: ascii}
		if err != nil {
			if err.Error() == "contains non ascii compatible characters" {
				// http.Error(w, "400: "+err.Error(), http.StatusBadRequest)
				w.WriteHeader(400)
				renderTemplate(w, "error", "400: "+err.Error())
			} else if strings.Contains(err.Error(), "no such file or directory") {
				// http.Error(w, "500: One of the files required is missing", http.StatusNotFound)
				w.WriteHeader(500)
				renderTemplate(w, "error", "500: One of the files required is missing")
			} else if strings.Contains(err.Error(), "bad Request") {
				w.WriteHeader(400)
				renderTemplate(w, "error", "400: "+err.Error())
			}
			return
		} else {
			// http.Redirect(w, r, "/ascii-art", http.StatusFound)
			w.WriteHeader(200)
			renderTemplate(w, "ascii-art", result)
		}
	} else {
		w.WriteHeader(400)
		renderTemplate(w, "error", "400: Input is too long")
		return
	}
}

func renderTemplate(w http.ResponseWriter, name string, data any) {
	tmp, err := template.ParseFiles("templates/" + name + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmp.Execute(w, data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", artHandler)

	log.Print("server running, open http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
