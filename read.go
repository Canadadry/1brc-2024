package main

import (
	"app/reader"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <R1-R7>")
		os.Exit(1)
	}

	file, err := os.Open("measurements.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := reader.Read(os.Args[1], file, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to process data: %v\n", err)
		os.Exit(1)
	}
}
