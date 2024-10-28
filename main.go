package main

import (
	"net/http"
	app "solidgo/app"
)

func main() {
	i := app.InertiaStart()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(app.HomeRoute(i)))
	mux.Handle("/about", i.Middleware(app.AboutRoute(i)))
	mux.Handle("/data", i.Middleware(app.DataRoute(i)))
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	http.ListenAndServe(":3000", mux)
}
