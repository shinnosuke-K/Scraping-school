package main

import (
	"fmt"
	"scraping-school/env"
	"scraping-school/prefectures"
	"time"
)

func scrapeForUrl() {
	schoolNames := readSchoolName()

	for _, schoolName := range schoolNames {
		fmt.Println(scrapeUrl(schoolName))
		time.Sleep(time.Second * 3)
	}
}

func scrapeForCourse() {
	for _, prefecture := range prefectures.Prefectures {
		CreateCSVfile("csv-name-course/" + prefecture + ".csv")
		scrapeCourse(env.SearchURL+prefecture+env.DeviationURL, "csv-name-course/"+prefecture+".csv")
		time.Sleep(time.Millisecond * 5)
	}
}

func main() {
	//scrapeForCourse()
	scrapeForUrl()
}
