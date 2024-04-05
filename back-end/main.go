package main

import (
	"html/template"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("../blog.html"))
		temp.Execute(w, nil)
	}
	http.HandleFunc("/blogs", h1)
	http.Post("losthost:5173/blogs", "application/json")
}
