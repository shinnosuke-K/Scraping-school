package main

import (
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"scraping-school/check"
	"scraping-school/env"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrapeUrl(schoolName string) string {
	// not client
	//res, err := http.Get("https://www.google.com/search?q=" + url.QueryEscape(schoolName[0]) + "&num=2")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// client

	client := &http.Client{}

	searchURL := "https://www.google.com/search?q=" + url.QueryEscape(schoolName) + "&num=2"

	res, err := http.NewRequest("GET", searchURL, nil)
	check.Error(err)

	res.Header.Set("User-Agent", env.MacOSOfChrome)

	resp, err := client.Do(res)
	check.Error(err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check.Error(err)

	var schoolURL string

	doc.Find(env.GoogleSelector).Each(func(i int, s *goquery.Selection) {
		if tmpURL, _ := s.Attr("href"); strings.Contains(tmpURL, "ed.jp/") || strings.Contains(tmpURL, "ac.jp/") {
			schoolURL = tmpURL
		}
	})

	return schoolURL
}

func readSchoolName() []string {
	var schoolName []string

	files, err := ioutil.ReadDir("csv-name-course/")
	check.Error(err)

	for _, file := range files {
		if fileName := file.Name(); fileName != ".DS_Store" {
			csvFile, err := os.Open("csv-name-course/" + fileName)
			check.Error(err)

			reader := csv.NewReader(csvFile)

			for {
				line, err := reader.Read()
				if err != nil {
					break
				}

				schoolName = append(schoolName, line[1])
			}
		}
	}

	return schoolName
}
