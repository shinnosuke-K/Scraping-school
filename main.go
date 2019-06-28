package main

import (
	"math/rand"
	"scraping-school/scrape"
	"time"
)

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

func ScrapeForUrl() {
	schoolInfos := scrape.ReadSchoolName()

	for _, schoolInfo := range schoolInfos {
		schoolInfo.SchoolUrl = scrape.ScrapeUrl(schoolInfo.Name, schoolInfo.FileName)
		scrape.WriteCSVForURL(schoolInfo)
		time.Sleep(time.Second * time.Duration(random(0.0, 5.0)))
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
