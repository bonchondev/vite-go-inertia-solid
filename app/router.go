package app

import (
	"encoding/json"
	"log"
	"net/http"

	inertia "github.com/romsar/gonertia"
)

type Data struct {
	Dollars int `json:"dollars"`
}

func DataRoute(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := Data{
			Dollars: 100,
		}
		json.NewEncoder(w).Encode(data)
	}

	return http.HandlerFunc(fn)
}

func AboutRoute(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "About", inertia.Props{
			"amount": "$7,000",
		})
		if err != nil {
			handleServerErr(w, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func HomeRoute(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w,r)
			return
		}
		err := i.Render(w, r, "Index")
		if err != nil {
			handleServerErr(w, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func handleServerErr(w http.ResponseWriter, err error) {
	log.Printf("http error: %s\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}
