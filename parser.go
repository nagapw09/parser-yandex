package main

import (
	"flag"
	"github.com/nagapw09/parser2/bd"
	"github.com/nagapw09/parser2/pars"
)

func main() {
	var arg1 = flag.String("1", "", "request")
	var arg2 = flag.String("2", "0", "pages")
	flag.Parse()

	bd.Base()
	pars.ExampleScrape(*arg1, *arg2)
}
