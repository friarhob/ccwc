package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)

	return
}

func isInSlice(elem interface{}, slice interface{}) bool {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic("Provided 'slice' is not a slice")
	}

	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), elem) {
			return true
		}
	}
	return false
}

func printHelpMessage() {
	printError("Usage: ccwc [options] [filepath]")
	printError("Options:")
	printError("   -c, --bytes : Print the number of bytes.")
	printError("   -l, --lines : Print the number of lines.")
	printError("   -w, --words : Print the number of words.")
	printError("   -m, --chars : Print the number of chars.")

	return
}

func main() {
	bytesParameters := []string{"-c", "--bytes"}
	linesParameters := []string{"-l", "--lines"}
	wordsParameters := []string{"-w", "--words"}
	charsParameters := []string{"-m", "--chars"}
	helpParameters := []string{"-h", "--help"}

	calculateBytes := false
	calculateLines := false
	calculateWords := false
	calculateChars := false

	parameters := os.Args[1:]

	var filepath string

	for _, param := range parameters {
		if isInSlice(param, bytesParameters) {
			calculateBytes = true
		} else if isInSlice(param, linesParameters) {
			calculateLines = true
		} else if isInSlice(param, wordsParameters) {
			calculateWords = true
		} else if isInSlice(param, charsParameters) {
			calculateChars = true
		} else if isInSlice(param, helpParameters) {
			printHelpMessage()
			return
		} else {
			filepath = param
		}
	}

	if !calculateBytes && !calculateChars && !calculateLines && !calculateWords {
		calculateLines = true
		calculateWords = true
		calculateBytes = true
	}

	var output string

	if calculateLines || calculateWords {
		file, err := os.Open(filepath)
		if err != nil {
			printError("Error opening file: " + filepath)
			return
		}
		defer file.Close()

		var lines int64 = 0
		var words int = 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			lines += 1
			words += len(strings.Fields(line))
		}

		if calculateLines {
			output += fmt.Sprintf(" %7d", lines)
		}
		if calculateWords {
			output += fmt.Sprintf(" %7d", words)
		}

		if err := scanner.Err(); err != nil {
			printError("Error reading file: " + filepath)
			return
		}
	}

	if calculateBytes {

		fileinfo, err := os.Stat(filepath)
		if err != nil {
			printError("Error reading file: " + filepath)
			return
		}

		output += fmt.Sprintf(" %7d", fileinfo.Size())
	}

	if calculateChars {
		file, err := os.Open(filepath)
		if err != nil {
			printError("Error opening file: " + filepath)
			return
		}
		defer file.Close()

		var chars int64 = 0

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			chars += 1
		}

		output += fmt.Sprintf(" %7d", chars)

		if err := scanner.Err(); err != nil {
			printError("Error reading file: " + filepath)
			return
		}
	}

	fmt.Println(output, filepath)
}
