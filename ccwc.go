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

func main() {
	bytesParameters := []string{"-c", "--bytes"}
	linesParameters := []string{"-l", "--lines"}
	wordsParameters := []string{"-w", "--words"}
	charsParameters := []string{"-m", "--chars"}

	validParameters := append(append(append(bytesParameters, linesParameters...), wordsParameters...), charsParameters...)

	parameters := os.Args[1:]
	if len(parameters) != 2 || !isInSlice(parameters[0], validParameters) {
		printError("Usage: ccwc [option] [filepath]")
		printError("Options:")
		printError("   -c, --bytes : Print the number of bytes.")
		printError("   -l, --lines : Print the number of lines.")
		printError("   -w, --words : Print the number of words.")
		printError("   -m, --chars : Print the number of chars.")
		return
	}

	filepath := parameters[len(parameters)-1]

	var outputValue int64 = 0

	if isInSlice(parameters[0], bytesParameters) {

		fileinfo, err := os.Stat(filepath)
		if err != nil {
			printError("Error reading file: " + filepath)
			return
		}

		outputValue = fileinfo.Size()
	} else if isInSlice(parameters[0], linesParameters) || isInSlice(parameters[0], wordsParameters) {
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

		if isInSlice(parameters[0], linesParameters) {
			outputValue = lines
		} else if isInSlice(parameters[0], wordsParameters) {
			outputValue = int64(words)
		}

		if err := scanner.Err(); err != nil {
			printError("Error reading file: " + filepath)
			return
		}
	} else if isInSlice(parameters[0], charsParameters) {
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

		outputValue = chars

		if err := scanner.Err(); err != nil {
			printError("Error reading file: " + filepath)
			return
		}

	}

	fmt.Printf("%8d %s\n", outputValue, filepath)
}
