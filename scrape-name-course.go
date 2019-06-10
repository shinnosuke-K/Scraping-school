package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"scraping-school/env"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func createCSVfile(fileName string) {
	if _, err := os.Stat("csv-name-course/" + fileName + ".csv-name-course"); err != nil {
		if _, err := os.Create("csv-name-course/" + fileName + ".csv-name-course"); err != nil {
			log.Fatal(err)
		}
	}
}

func writeCSV(deviValue string, schoolName string, course string, prefecture string) {
	file, err := os.OpenFile("csv-name-course/"+prefecture+".csv-name-course", os.O_WRONLY|os.O_CREATE, 0600)
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

	//for _, prefecture := range prefectures.Prefectures {
	//	createCSVfile(prefecture)
	//	Scrape(env.SearchURL+prefecture+env.DeviationURL, prefecture)
	//	time.Sleep(5)
	//}

	files, err := ioutil.ReadDir("csv-name-course/")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.Name() == ".DS_Store" {
			continue
		}

		fmt.Println(file.Name())
		csvFile, err := os.Open("csv-name-course/" + file.Name())
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(csvFile)

		line, _ := reader.ReadAll()

		for k, v := range line {
			fmt.Println(k, v[2])
		}
		//fmt.Println(line)
	}
}
