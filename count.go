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
	traverseLogs()
}

func traverseLogs() {
	// traverse all .txt files

	files, err := ioutil.ReadDir("pfs/lb-pachy-project")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// only look at .txt files
		if strings.HasSuffix(file.Name(), ".txt") {
			content, err := ioutil.ReadFile("pfs/lb-pachy-project/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			countWarningsAndErrors(string(content))
		}
	}
}

//TODO: Edge cases?
func countWarningsAndErrors(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.Contains(line, "ERR") {
			errorCount++
		} else if strings.Contains(line, "WARN") {
			warningCount++
		}
	}
	createResultsFile(errorCount, warningCount)
}

func createResultsFile(errorCount int, warningCount int) {
	results := "errorCount: " + fmt.Sprint(errorCount) + "\n" + "warningCount: " + fmt.Sprint(warningCount)
	file := ioutil.WriteFile("/results.txt", []byte(results), 0644)
	if file != nil {
		log.Fatal(file)
	}

}
