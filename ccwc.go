package main

import (
	// "bufio"
	"fmt"
	"os"
)

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)

	return
}

func main() {
	parameters := os.Args[1:]
	if len(parameters) != 2 || parameters[0] != "-c" {
		printError("Usage: ccwc -c filepath")
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
