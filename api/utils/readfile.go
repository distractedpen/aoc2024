package utils

import (
	"log"
	"os"
	"strings"
)

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Readfile(filename string) string {
	data, err := os.ReadFile(filename)
	CheckFatal(err)
	parsedData := strings.Trim(string(data), "\n")
	return (parsedData)
}
