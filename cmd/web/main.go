package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates *template.Template

//type PageData struct {
//	Title       string
//	NewsSection []NewsItem
//}
//type NewsItem struct {
//	Hashtag string
//	Image   string
//	Content string
//	Link    string
//}

//func loadTemplates() {
//	templates = template.Must(template.ParseGlob("ui/html/*.html"))
//}
//func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
//	err := templates.ExecuteTemplate(w, tmpl+".html", data)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//files := []string{
	//	"./ui/html/base.tmpl",
	//	"./ui/html/partials/header.tmpl",
	//	"./ui/html/pages/home.tmpl",
	//	"./ui/html/pages/footer.tmpl",
	//	"./ui/html/pages/content.tmpl",
	//}

	//pageData := PageData{
	//	Title: "News Website",
	//	NewsSection: []NewsItem{
	//		{
	//			Hashtag: "#AstanaITScience",
	//			Image:   "/static/img/news/NIRS.png",
	//			Content: "Intra-university stage of the student research competition (SRW) at Astana IT University...",
	//			Link:    "#",
	//		},
	//		// Add other news items here
	//	},
	//}
	//
	//renderTemplate(w, "base", pageData)
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
