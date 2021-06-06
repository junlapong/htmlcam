package main

import (
	_ "embed"
	"log"
	"net/http"
)

//go:embed index.html
var html string

func main() {

	m := http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})

	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}

	log.Printf("http://localhost%s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
