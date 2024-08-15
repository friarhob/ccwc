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

func calculateLinesWords(filepath string) (int64, int64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return -1, -1, err
	}
	defer file.Close()

	var lines int64 = 0
	var words int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lines += 1
		words += int64(len(strings.Fields(line)))
	}

	if err := scanner.Err(); err != nil {
		return -1, -1, err
	}

	return lines, words, nil
}

func calculateBytes(filepath string) (int64, error) {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		return -1, err
	}

	return fileinfo.Size(), nil
}

func calculateChars(filepath string) (int64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var chars int64 = 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		chars += 1
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return chars, nil
}

func main() {
	bytesParameters := []string{"-c", "--bytes"}
	linesParameters := []string{"-l", "--lines"}
	wordsParameters := []string{"-w", "--words"}
	charsParameters := []string{"-m", "--chars"}
	helpParameters := []string{"-h", "--help"}

	flagBytes := false
	flagLines := false
	flagWords := false
	flagChars := false

	parameters := os.Args[1:]

	var filepath string

	for _, param := range parameters {
		if isInSlice(param, bytesParameters) {
			flagBytes = true
		} else if isInSlice(param, linesParameters) {
			flagLines = true
		} else if isInSlice(param, wordsParameters) {
			flagWords = true
		} else if isInSlice(param, charsParameters) {
			flagChars = true
		} else if isInSlice(param, helpParameters) {
			printHelpMessage()
			return
		} else {
			filepath = param
		}
	}

	if !flagBytes && !flagChars && !flagLines && !flagWords {
		flagLines = true
		flagWords = true
		flagBytes = true
	}

	var output string

	if flagLines || flagWords {
		lines, words, err := calculateLinesWords(filepath)

		if err != nil {
			printError("Error reading file " + filepath)
			return
		}

		if flagLines {
			output += fmt.Sprintf(" %7d", lines)
		}
		if flagWords {
			output += fmt.Sprintf(" %7d", words)
		}
	}

	if flagBytes {
		bytes, err := calculateBytes(filepath)
		if err != nil {
			printError("Error reading file: " + filepath)
			return
		}

		output += fmt.Sprintf(" %7d", bytes)
	}

	if flagChars {
		chars, err := calculateChars(filepath)
		if err != nil {
			printError("Error reading file: " + filepath)
			return
		}

		output += fmt.Sprintf(" %7d", chars)
	}

	fmt.Println(output, filepath)
}
