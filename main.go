package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func findMatchingLines(r io.Reader, re *regexp.Regexp) ([]string, error) {
	scanner := bufio.NewScanner(r)
	lineMatches := []string{}
	for scanner.Scan() {

		if re.MatchString(scanner.Text()) {
			lineMatches = append(lineMatches, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lineMatches, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("ERROR: File not found!")
		fmt.Println("Usage: mygrep <string to search> <file to examine>")
		os.Exit(1)
	} //https://golangcode.com/check-if-a-file-exists/

	return !info.IsDir()
}

//myGrep

func main() {

	args := os.Args[1:]
	findString := args[0]
	examFile := args[1]

	//make sure the user inputs arguments
	if len(os.Args) > 2 {
		fmt.Printf("looking for \"%s\".\n", findString)
	} else {
		fmt.Println("Usage: mygrep <string to search> <file to examine>")
		os.Exit(0)
	}

	//be sure file is available
	if fileExists(examFile) {
		fmt.Println("Examining file:", examFile)
		fmt.Println("")
	}

	dat, err := ioutil.ReadFile(examFile)
	if err != nil {
		fmt.Println("File read error: ", err)
		os.Exit(1)
	} //modified from https://gobyexample.com/reading-files

	r := bytes.NewReader(dat)
	re := regexp.MustCompile(findString)
	lines, err := findMatchingLines(r, re)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} //https://golang.org/pkg/bufio/#example_Scanner_lines

	for _, l := range lines {
		fmt.Println(l)
	}
}
