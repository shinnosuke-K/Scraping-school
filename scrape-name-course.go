package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"scraping-school/check"
	"scraping-school/env"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func writeCSVForCourse(deviValue string, schoolName string, course string, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	check.Error(err)

	info := []string{
		deviValue,
		schoolName,
		course,
	}

	writer := csv.NewWriter(file)
	writer.Write(info)

	writer.Flush()
}

func searchName(deviValue string, schoolInfo string, filename string) {

	var schoolName string
	var course string

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

						writeCSVForCourse(deviValue, schoolName, course, filename)

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

func scrapeCourse(url string, prefecture string) {
	res, err := http.Get(url)
	check.Error(err)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	check.Error(err)

	doc.Find(env.Selector).Each(func(i int, s *goquery.Selection) {
		deviValue := s.Find(env.DeviValueSelector).Text()
		schoolInfo := s.Find("td > ul > li").Text()
		searchName(deviValue, schoolInfo, prefecture)
	})
}
