package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"pribadi/assistantwife/webapps/controllers"
)

var templates map[string]*template.Template
var baseTemplate []string

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	baseTemplate = []string{"views/layouts/base.html", "views/layouts/head.html"}
}
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "the template does not exist.", http.StatusInternalServerError)
	}

	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	templateBase := baseTemplate
	templateBase = append(templateBase, "views/Shop/index.html")
	templates["index"] = template.Must(template.ParseFiles(templateBase...))

	renderTemplate(w, "index", "base", nil)
}

func debt(w http.ResponseWriter, r *http.Request) {
	templateBase := baseTemplate
	templateBase = append(templateBase, "views/Debt/index.html")

	log.Println(templateBase)
	templates["debt"] = template.Must(template.ParseFiles(templateBase...))

	renderTemplate(w, "debt", "base", nil)
}

func main() {
	mux := http.NewServeMux()

	pathBase, _ := os.Getwd()
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir(pathBase+"/assets")))

	mux.Handle("/assets/", fs)

	shopC := controllers.ShopController{}

	mux.HandleFunc("/", index)
	mux.HandleFunc("/shop/get", shopC.Get)
	mux.HandleFunc("/shop/save", shopC.Save)
	mux.HandleFunc("/shop/delete", shopC.Delete)

	mux.HandleFunc("/debt", debt)

	server := &http.Server{
		Addr:    ":8022",
		Handler: mux,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
