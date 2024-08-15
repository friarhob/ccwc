package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"unicode"
)

type stats struct {
	bytes int64
	words int64
	lines int64
	chars int64
}

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
	printError("Usage: ccwc [options] [filepaths]")
	printError("Options:")
	printError("   -c, --bytes : Print the number of bytes.")
	printError("   -l, --lines : Print the number of lines.")
	printError("   -w, --words : Print the number of words.")
	printError("   -m, --chars : Print the number of chars.")

	return
}

func calculateStats(reader bufio.Reader) (stats, error) {
	var results stats

	var prevChar rune = rune(0)

	for {
		curChar, bytesRead, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				return results, nil
			}

			return results, err
		}

		if curChar == '\n' {
			results.lines += 1
		}

		if (unicode.IsSpace(prevChar) || prevChar == rune(0)) && !unicode.IsSpace(curChar) {
			results.words += 1
		}

		results.bytes += int64(bytesRead)
		results.chars += 1

		prevChar = curChar
	}
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

	var filepaths []string

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
			filepaths = append(filepaths, param)
		}
	}

	if !flagBytes && !flagChars && !flagLines && !flagWords {
		flagLines = true
		flagWords = true
		flagBytes = true
	}

	if len(filepaths) == 0 {
		printError("Error: no filepaths specified")
		printHelpMessage()
		return
	}

	for _, filepath := range filepaths {
		var output string

		file, err := os.Open(filepath)

		if err != nil {
			printError("Error reading file: " + filepath)
			return
		}

		calculations, err := calculateStats(*bufio.NewReader(file))

		if err != nil {
			printError("Error reading file: " + filepath)
			return
		}

		if flagLines {
			output += fmt.Sprintf(" %7d", calculations.lines)
		}

		if flagWords {
			output += fmt.Sprintf(" %7d", calculations.words)
		}

		if flagBytes {
			output += fmt.Sprintf(" %7d", calculations.bytes)
		}

		if flagChars {
			output += fmt.Sprintf(" %7d", calculations.chars)
		}

		fmt.Println(output, filepath)
	}

}
