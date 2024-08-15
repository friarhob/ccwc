package main

import (
	// "bufio"
	"fmt"
	"os"
	"reflect"
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
	validParameters := []string{"-c", "--bytes"}

	parameters := os.Args[1:]
	if len(parameters) != 2 || !isInSlice(parameters[0], validParameters) {
		printError("Usage: ccwc [options] [filepath]")
		printError("Options:")
		printError("   -c, --bytes : Print the number of bytes.")
		return
	}

	filepath := parameters[len(parameters)-1]

	fileinfo, err := os.Stat(filepath)
	if err != nil {
		printError("Error reading file:" + filepath)
		return
	}

	fmt.Printf("%8d %s\n", fileinfo.Size(), filepath)

	return

	/*
			file, err := os.Open(filepath)
			if err != nil {
				printError("Error opening file: " + filepath)
				return
			}
			defer file.Close()

		    scanner := bufio.NewScanner(file)
		    for scanner.Scan() {
		        fmt.Println(scanner.Text()) // Print each line
		    }

		    if err := scanner.Err(); err != nil {
		        fmt.Println("Error reading file:", err)
		    }
	*/
}
