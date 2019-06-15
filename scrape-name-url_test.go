package main

import (
	"scraping-school/env"
	"testing"
)

func Test_scrapeUrl(t *testing.T) {

	tests := []struct {
		name       string
		schoolName string
		want       string
	}{
		// TODO: Add test cases.
		{"存在する高校名", env.HighSchoolName[0], env.HighSchoolURL[0]},
		{"海外の高校名", env.HighSchoolName[1], ""},
		{"存在しない高校名", env.HighSchoolName[2], ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrapeUrl(tt.schoolName); got != tt.want {
				t.Errorf("scrapeForUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
