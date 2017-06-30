package models

import (
	"io/ioutil"
	"log"
)

func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read the file: %s", filename)
	}
	return string(b)
}
