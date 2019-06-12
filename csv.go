package main

import (
	"log"
	"os"
)

func CreateCSVfile(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Fatal(err)
		}
	}
}
