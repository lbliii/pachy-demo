package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	errorCount   int
	warningCount int
)

func main() {
	files, err := ioutil.ReadDir("logs")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		content := readFile("logs/" + file.Name())
		//fmt.Println(content)
		countWarningsAndErrors(content)
	}
	fmt.Println("errorCount:", errorCount)
	fmt.Println("warningCount:", warningCount)
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func countWarningsAndErrors(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.Contains(line, "ERR") {
			errorCount++
		} else if strings.Contains(line, "WARN") {
			warningCount++
		}
	}
}
