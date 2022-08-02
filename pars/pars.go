package pars

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nagapw09/parser2/bd"
	"log"
	"net/http"
)

func ExampleScrape(w http.ResponseWriter) {
	var request, urls, pages string
	fmt.Println("Что ищем?")
	fmt.Scanf("%s\n", &request)
	fmt.Println("Какую страниц парсим?")
	fmt.Scanf("%s\n", &pages)
	urls = ("https://yandex.ru/search/?text=" + request + "&lr=213&p=" + pages)
	res, err := http.Get(urls)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".serp-item").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		header := s.Find(".OrganicTitleContentSpan").Text()

		fmt.Fprintf(w, "№ = %d Link - %s Header - %s\n", i, link, header)
		bd.Insert(link, header, pages)
		fmt.Println("Закончил парсить")
	})

}
