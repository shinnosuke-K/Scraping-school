package scrape

import (
	"scraping-school/env"
	"testing"
)

func Test_scrapeUrl(t *testing.T) {

	tests := []struct {
		name       string
		schoolName string
		fileName   string
		want       string
	}{
		// TODO: Add test cases.
		{"存在する高校名", env.HighSchoolName[0], env.FileName[0], env.HighSchoolURL[0]},
		{"存在する高校名", env.HighSchoolName[1], env.FileName[1], env.HighSchoolURL[1]},
		{"同じ名前の高校名", env.HighSchoolName[2], env.FileName[2], env.HighSchoolURL[2]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScrapeUrl(tt.schoolName, tt.fileName); got != tt.want {
				t.Errorf("scrapeForUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
