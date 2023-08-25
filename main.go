package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/2k4sm/shrinibas-SST-Hackathon-1/moviedb"
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type Search struct {
	Query  string
	Movies *moviedb.Movies
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("PORT")

	if port != "" {
		port = "3000"
	}
	apikey := os.Getenv("MOVIE_API_KEY")

	if apikey == "" {
		log.Fatal("Env: apikey must be set")
	}
	myClient := &http.Client{Timeout: 10 * time.Second}
	movieapi := moviedb.NewClient(myClient, apikey)

	mux := http.NewServeMux()
	//fs := http.FileServer(http.Dir("assets"))
	//mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", searchHandler(movieapi))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func searchHandler(movieapi *moviedb.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ul, err := url.Parse(r.URL.String())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := ul.Query()

		searchQuerry := params.Get("squery")

		results, err := movieapi.FetchMovie(searchQuerry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(results)

	}

}
