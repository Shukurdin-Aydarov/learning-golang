package main

import (
	"chitchat/data"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		return
	}
	_, err = session(writer, request)
	if err != nil {
		generateHtml(writer, threads, "layout", "public.navbar", "index")
	} else {
		generateHtml(writer, threads, "layout", "private.navbar", "index")
	}
}

func generateHtml(w http.ResponseWriter, data interface{}, templates ...string) {
	var files []string
	for _, file := range templates {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t := template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", data)
}

func err(writer http.ResponseWriter, request *http.Request) {

}

func login(writer http.ResponseWriter, request *http.Request) {

}

func logout(writer http.ResponseWriter, request *http.Request) {

}

func signup(writer http.ResponseWriter, request *http.Request) {

}

func signupAccount(writer http.ResponseWriter, request *http.Request) {

}

func newThread(writer http.ResponseWriter, request *http.Request) {

}

func createThread(writer http.ResponseWriter, request *http.Request) {

}

func postThread(writer http.ResponseWriter, request *http.Request) {

}

func readThread(writer http.ResponseWriter, request *http.Request) {

}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}
