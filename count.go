package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	files, err := ioutil.ReadDir("/pfs/lb-pachy-project")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(file.Name())
		content := readFile("/pfs/lb-pachy-project/" + file.Name())
		// fmt.Println(content)
		fileName := file.Name()
		countWarningsAndErrors(content, fileName)
	}
	// fmt.Println("errorCount:", errorCount)
	// fmt.Println("warningCount:", warningCount)
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
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
	// if results.txt exists, update it, otherwise create it
	if _, err := os.Stat("/pfs/out/results.txt"); err == nil {
		createResultsFile(errorCount, warningCount)

	} else {
		updateResultsFile(errorCount, warningCount)
	}

}

func createResultsFile(errorCount int, warningCount int) {
	results := "errorCount: " + fmt.Sprint(errorCount) + "\n" + "warningCount: " + fmt.Sprint(warningCount)
	file := ioutil.WriteFile("/pfs/out/results.txt", []byte(results), 0644)
	if file != nil {
		log.Fatal(file)
	}

}

// if results.txt exists, update it with the new counts

func updateResultsFile(errorCount int, warningCount int) {
	results := "errorCount: " + fmt.Sprint(errorCount) + "\n" + "warningCount: " + fmt.Sprint(warningCount)
	file := ioutil.WriteFile("/pfs/out/results.txt", []byte(results), 0644)
	if file != nil {
		log.Fatal(file)
	}

}
