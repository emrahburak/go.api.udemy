package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	x := r.URL.Path[1:]
	data := ""

	if len(x) > 0 {
		data = " Path: " + x
	} else {
		data = "Index Page"
	}

	w.Write([]byte(data))

}

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("About Page"))
// }

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	// http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)

}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Merhaba Mars!"))
// })
