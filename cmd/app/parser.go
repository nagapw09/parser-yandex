package main

import (
	"flag"
	"github.com/nagapw09/parser2/pkg/bd"
	"github.com/nagapw09/parser2/pkg/pars"
)

func main() {
	var arg1 = flag.String("1", "", "request")
	var arg2 = flag.Int("2", 0, "page")
	flag.Parse()

	bd.CreateTable()
	pars.ExampleScrape(*arg1, *arg2)
}
