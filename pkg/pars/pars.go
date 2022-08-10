package pars

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nagapw09/parser2/pkg/bd"
	"github.com/nagapw09/parser2/pkg/logger"
	_ "github.com/nagapw09/parser2/pkg/logger"
	"log"
	"net/http"
	"strconv"
)

func ExampleScrape(request string, page int) {
	var urls string

	urls = "https://yandex.ru/search/?text=" + request + "&lr=213&p=" + strconv.Itoa(page)
	res, err := http.Get(urls)
	logger.ForError(err)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	logger.ForError(err)

	doc.Find(".serp-item").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		header := s.Find(".OrganicTitleContentSpan").Text()

		fmt.Printf("№ = %d Link - %s Header - %s\n", i, link, header)
		bd.Insert(link, header, request, page)
	})
}
