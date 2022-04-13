package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func OpenFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}
	return f, err
}

func ReadCsv(f io.Reader) [][]string {
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV", err)
	}
	return records
}

func main() {
	scoreCount := 0
	file, _ := OpenFile("package.csv")
	quizzes := ReadCsv(file)
	for _, quiz := range quizzes {
		question := quiz[0]
		expectedAnswer := quiz[1]
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("What is %s: ", question)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if expectedAnswer == text {
			scoreCount++
		}
	}
	fmt.Printf("Your total score is: %d", scoreCount)
}
