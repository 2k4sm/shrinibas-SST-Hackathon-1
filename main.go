package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("PORT")

	if port != "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	ul, err := url.Parse(r.URL.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := ul.Query()

	searchQuerry := params.Get("squery")

	fmt.Println("Search querry is:", searchQuerry)

}
