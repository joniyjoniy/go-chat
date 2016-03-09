package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"flag"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", PORT, "application's address")
	flag.Parse()
	gomniauth.SetSecurityKey(SECURITYKEY)
	gomniauth.WithProviders(
		facebook.New(FACEBOOKCLIENTID, FACEBOOKCLIENTSECRET, ADDR + PORT + "/auth/callback/facebook")
		github.New(GITHUBCLIENTID, GITHUBCLIENTSECRET, ADDR + PORT + "/auth/callback/github")
		google.New(GOOGLECLIENTID, GOOGLECLIENTSECRET, ADDR + PORT + "/auth/callback/google")
	)
	r := newRoom()
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバを開始 ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
