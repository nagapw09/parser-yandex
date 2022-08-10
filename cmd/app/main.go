package main

import (
	"fmt"
	"github.com/nagapw09/parser2/pkg/bd"
	"github.com/nagapw09/parser2/pkg/pars"
)

func main() {
	err := bd.CreateTable()
	if err != nil {
		fmt.Printf("%s", err)
	}
	pars.ExampleScrape("jkl", 0)
}
