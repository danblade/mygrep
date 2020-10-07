package main

import (
	"bufio"
	"fmt"
	"io"
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

func main() { //myGrep
	var r io.Reader

	if len(os.Args) == 3 {
		file, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		r = file
	} else if len(os.Args) == 2 {
		r = os.Stdin
	} else {
		fmt.Println("Usage: mygrep <string to search> <file to examine>")
		os.Exit(0)
	}

	findString := os.Args[1]
	fmt.Printf("looking for \"%s\".\n", findString)

	re, err := regexp.Compile(findString)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	lines, err := findMatchingLines(r, re)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
