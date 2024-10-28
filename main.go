package main

import (
	"net/http"
)

func main() {
	i := InertiaStart()

	mux := http.NewServeMux()

	mux.Handle("/home", i.Middleware(homeRoute(i)))
	mux.Handle("/about", i.Middleware(aboutRoute(i)))
	mux.Handle("/data", i.Middleware(dataRoute(i)))
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	http.ListenAndServe(":3000", mux)
}
