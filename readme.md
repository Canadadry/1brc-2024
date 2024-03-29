# The 1 Billion Row Challenge

Welcome to the 1 Billion Row Challenge, a performance-oriented project aimed at processing a massive dataset consisting of up to 1 billion rows. This project is designed to explore and compare different Go programming techniques and optimizations through several implementations (R1 to R7) of a data processing function, ranging from trivial to highly optimized solutions.

## Project Structure

The project is composed of the following key components:

 - `reader/reader.go` : Main Processing Functions (R1-R7), these are different implementations of the data processing function. Each version, from R1 to R7, represents a unique approach with varying levels of efficiency and optimization.

 - `generator/gen.go`: A utility program designed to generate datasets of arbitrary size. It takes the number of rows to create as its first command-line argument.

 - `measurements.txt`: The default dataset file used by the processing functions. Generated by create.go, this file contains simulated station temperature data.

 - main.go: The entry point of the application, which allows the user to select and execute one of the R1-R7 functions against the dataset.

## Getting Started

### Prerequisites

 - Go programming language (version 1.13 or later recommended)
 - Basic understanding of command-line operations
 - Sufficient disk space for large datasets (generating 1 billion rows requires significant storage)

### Generating Data

To generate a dataset, use the create.go script as follows:

```bash
go run generator/gen.go <number-of-rows>
```


Replace <number-of-rows> with the desired dataset size. For example, to create a file with 1 million rows, you would run:

```bash
go run generator/gen.go 1000000
```

This will generate a file named measurements.txt in the current directory, populated with the specified number of rows.

### Running the Challenge

To execute a specific implementation against the generated dataset, use the main.go script with the desired function as an argument:

```bash
time go run main.go Rx
```

Replace R1 with any function from R1 to R7 to test different implementations. Each function will process the measurements.txt file and output the results to stdout.

###Comparing Performance

The main goal of the 1 Billion Row Challenge is to compare the performance of different approaches to processing large datasets. Users are encouraged to analyze the runtime, memory usage, and CPU efficiency of each implementation (R1-R7) and consider why certain strategies perform better than others.

## Launching Tests and Benchmarks

### Running Unit Tests

To run the unit tests for all implementations, use the go test command in the project directory:

```bash
go test -v
```

This will execute all defined unit tests, verifying the correctness of each implementation.

### Running Benchmarks

The project includes benchmarks to measure the performance of each implementation. To run these benchmarks, use the go test command with the -bench flag. For example, to benchmark the R1 implementation, run:

```bash
go test -bench=.
```

This will execute all benchmarks whose names match the regular expression, providing insights into the performance of each implementation across varying dataset sizes.
