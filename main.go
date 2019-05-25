package main

import (
	"log"
	"net/http"
	"os"
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

	doc.Find("#app").Each(func(i int, s *goquery.Selection) {
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func fileScrape() {
	f, err := os.Open(env.DemoURL)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Find("#app").Each(func(i int, s *goquery.Selection) {
		band := s.Find(".result").Text()
		fmt.Println(band)
	}))
}

func main() {
	//url := "https://example.com"
	//Scrape(url)

	fileScrape()
}
