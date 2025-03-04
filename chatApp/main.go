package main

import (
	"chatApp/trace"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//why using once ?
	t.once.Do(func() {
		t.templ =
		template.Must(template.ParseFiles(filepath.Join("templates",t.filename)))})
		data := map[string]interface{}{
		"Host": r.Host,
		}
		if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
		}
		t.templ.Execute(w, data)
}
//tracing in go for testing ?

func main() {
	var addr = flag.String("addr", ":5301", "The addr of the application.")
	flag.Parse() // parse the flags

	gomniauth.SetSecurityKey("this is  a very obstacle full ay where god says tes and we pray ")
gomniauth.WithProviders(facebook.New("", "V",""),
github.New(),
google.New()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	//http.Handle("/", &templateHandler{filename: "chat.html"})
	//diff btw handle and handlefunc?
	http.Handle("/room", r)
	http.Handle("/chat", MustAuth(&templateHandler{filename:
		"chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})	
	http.HandleFunc("/auth/",loginHandler)

	// get the room going on defined server 
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
// set up gomniauth ?
