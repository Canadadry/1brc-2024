package main

import (
	"app/reader"
	"fmt"
	"io"
	"os"
)

// Define a function type for the Rs functions
type rsFunction func(io.Reader, io.Writer) error

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <R1-R7>")
		os.Exit(1)
	}

	// Map of function names to their implementations
	rsFunctions := map[string]rsFunction{
		"R1": reader.R1,
	}

	// Get the function selection from the command line arguments
	selectedFunction := os.Args[1]
	rsFunc, ok := rsFunctions[selectedFunction]
	if !ok {
		fmt.Fprintf(os.Stderr, "Invalid function selection. Please choose from R1 to R7.\n")
		os.Exit(1)
	}

	// Open the file "measurements.txt" for reading
	file, err := os.Open("measurements.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Call the selected function, reading from the file and writing to stdout
	if err := rsFunc(file, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to process data: %v\n", err)
		os.Exit(1)
	}
}
