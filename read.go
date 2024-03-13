package main

import (
	"app/reader"
	"fmt"
	"os"
	"runtime/pprof"
)

var cpuProfile = "cpu.profil"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <R1-R7>")
		os.Exit(1)
	}

	f, err := os.Create(cpuProfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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
