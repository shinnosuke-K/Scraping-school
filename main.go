package main

import (
	"log"
	"net/http"
	"scraping-school/env"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(url string) {
	res, err := http.Get(url)
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

	doc.Find(env.Selector).Each(func(i int, s *goquery.Selection) {
		deviValue := s.Find(env.DeviValueSelector).Text()
		schoolName := s.Find(env.SchoolNameSelector).Text()
		course := s.Find(env.CourseSelector).Text()
		fmt.Println(deviValue, schoolName, course)
	})
}

func main() {
	Scrape(env.SearchURL)
}
