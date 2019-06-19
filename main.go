package main

import (
	"scraping-school/scrape"
	"time"
)

func ScrapeForUrl() {
	schoolInfos := scrape.ReadSchoolName()

	for _, schoolInfo := range schoolInfos {
		schoolInfo.SchoolUrl = scrape.ScrapeUrl(schoolInfo.Name)
		scrape.WriteCSVForURL(schoolInfo)
		time.Sleep(time.Second * 2)
	}
}

//func scrapeForCourse() {
//	for _, prefecture := range prefectures.Prefectures {
//		CreateCSVfile("csv-name-course/" + prefecture + ".csv")
//		scrapeCourse(env.SearchURL+prefecture+env.DeviationURL, "csv-name-course/"+prefecture+".csv")
//		time.Sleep(time.Millisecond * 5)
//	}
//}

func main() {
	//scrapeForCourse()
	ScrapeForUrl()
}
