package main

import (
	"Web/cfg"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type IndexData struct {
	Title   string
	Content string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/post2b", b)

	fmt.Println(cfg.Host)
	fmt.Println(Hoo)
	fmt.Println(Hbb)

	err := http.ListenAndServe(":8964", nil)
	if err != nil {
		log.Fatal("LAS: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./a.html"))
	data := new(IndexData)
	data.Title = "GoWeb"
	data.Content = "GoWebContent"
	template.Execute(w, data)
}

func b(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	template := template.Must(template.ParseFiles("./b.html"))
	data := new(IndexData)
	var s string = ""

	for k, v := range r.Form {

		data.Title = k
		data.Content = strings.Join(v, s)
	}
	template.Execute(w, data)
}
