package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

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

	scanner := bufio.NewScanner(bytes.NewReader(dat))
	var re = regexp.MustCompile(findString)

	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} //https://golang.org/pkg/bufio/#example_Scanner_lines

	// match, _ := regexp.MatchString("([a-z]+)", findString)
	// if !match {
	// 	fmt.Printf("Unable to find %s.", findString)
	// 	os.Exit(1)
	// }
	// r, _ := regexp.Compile("([a-z]+)") //https://gobyexample.com/regular-expressions

	// fmt.Println(r.FindAllString(findString, -1))

}
