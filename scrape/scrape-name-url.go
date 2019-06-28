package scrape

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"scraping-school/check"
	"scraping-school/env"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type school struct {
	Devi      string
	Name      string
	Course    string
	SchoolUrl string
	FileName  string
}

func WriteCSVForURL(schoolInfo school) {
	file, err := os.OpenFile("csv-name-url/"+schoolInfo.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	check.Error(err)

	info := []string{
		schoolInfo.Devi,
		schoolInfo.Name,
		schoolInfo.Course,
		schoolInfo.SchoolUrl,
	}

	writer := csv.NewWriter(file)
	err = writer.Write(info)
	check.Error(err)

	writer.Flush()
}

func ScrapeUrl(schoolName, fileName string) string {
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
	// not client

	// client
	client := &http.Client{}

	searchURL := "https://www.google.com/search?q=" + url.QueryEscape(schoolName) + "+-wiki+-" + url.QueryEscape(env.DeleteSearchWord) + "+" + strings.Replace(fileName, ".csv", "", 1) + "&num=1"

	fmt.Println(searchURL)

	res, err := http.NewRequest("GET", searchURL, nil)
	check.Error(err)

	res.Header.Set("User-Agent", env.MacOSOfChrome)

	resp, err := client.Do(res)
	check.Error(err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check.Error(err)
	// client

	var schoolURL string

	doc.Find(env.GoogleSelector).Each(func(i int, s *goquery.Selection) {
		schoolURL, _ = s.Attr("href")
	})

	return schoolURL
}

func ReadSchoolName() []school {

	schoolInfo := make([]school, 0)

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

				schoolInfo = append(schoolInfo, school{line[0], line[1], line[2], "", fileName})
			}
		}
	}

	return schoolInfo
}
