package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"scraping-school/env"
	"scraping-school/prefectures"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func createCSVfile(fileName string) {
	if _, err := os.Stat("csv/" + fileName + ".csv"); err != nil {
		if _, err := os.Create("csv/" + fileName + ".csv"); err != nil {
			log.Fatal(err)
		}
	}
}

func writeCSV(deviValue string, schoolName string, course string, prefecture string) {
	file, err := os.OpenFile("csv/"+prefecture+".csv", os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{
		deviValue,
		schoolName,
		course,
	}

	writer := csv.NewWriter(file)
	writer.Write(info)

	writer.Flush()
}

func searchName(deviValue string, schoolInfo string, prefecture string) {

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

						writeCSV(deviValue, schoolName, course, prefecture)

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

func Scrape(url string, prefecture string) {
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
		searchName(deviValue, schoolInfo, prefecture)
	})
}

func main() {

	for _, prefecture := range prefectures.Prefectures {
		createCSVfile(prefecture)
		Scrape(env.SearchURL+prefecture+env.DeviationURL, prefecture)
		time.Sleep(5)
	}

}
