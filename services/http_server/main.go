package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/papa", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Just for fun"))
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World from a container"))
	})

	http.ListenAndServe(":8000", mux)
}
