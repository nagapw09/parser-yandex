package server

import (
	"github.com/nagapw09/parser2/pars"
	"log"
	"net/http"
)

func Servers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pars.ExampleScrape(w)
	})
	err := http.ListenAndServe(":4444", nil)

	log.Fatal("ListenAndServe: ", err)
}
