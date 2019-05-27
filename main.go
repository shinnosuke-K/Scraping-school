package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"scraping-school/env"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var schoolName string
var course string

func createCSVfile() {
	if _, err := os.Stat(env.CSVFileName); err != nil {
		if _, err := os.Create(env.CSVFileName); err != nil {
			log.Fatal(err)
		}
	}
}

func writeCSV(deviValue string, schoolInfo string) {
	file, err := os.OpenFile(env.CSVFileName, os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	schoolInfo = strings.ReplaceAll(schoolInfo, "\n", "")
	schoolInfo = strings.ReplaceAll(schoolInfo, " ", "")

	for schoolInfo != "" {
		for _, value := range schoolInfo {
			if schoolChar := string([]rune{value}); schoolChar == "（" {
				schoolInfo = strings.Replace(schoolInfo, schoolName, "", 1)

				for _, value := range schoolInfo {
					if courseChar := string([]rune{value}); courseChar == "）" && strings.Contains(course, "/") {
						course += courseChar
						schoolInfo = strings.Replace(schoolInfo, course, "", 1)

						info := []string{
							deviValue,
							schoolName,
							course,
						}

						writer := csv.NewWriter(file)
						writer.Write(info)

						writer.Flush()

						schoolName = ""
						course = ""
						break

					} else {
						course += courseChar
					}
				}
				break
			} else {
				schoolName += schoolChar
			}
		}
	}

}

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
		schoolInfo := s.Find("td > ul > li").Text()
		writeCSV(deviValue, schoolInfo)
	})
}

func main() {
	createCSVfile()
	Scrape(env.SearchURL)
}
