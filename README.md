# Data Calculation

This project is a Go application designed to read a JSON file containing an array of objects, compute the sum of specific fields, and print the result. The application is designed to utilize goroutines for parallel processing to improve performance on large datasets.

## Features

### Implemented as per the Technical Specification (TS)

- Reads data from a JSON file.
- Computes the sum of fields `a` and `b` for all objects in the JSON file.
- Utilizes goroutines for parallel processing to handle large data sets.
- Configurable number of goroutines via command-line arguments.
- Displays the total sum in the console.

### Additional Features

- Generates test data for testing purposes.
- Comprehensive testing for verifying the functionality of each component.
- Makefile for streamlined build, run, test, and generate operations.

## Requirements

- Go 1.20 or later.

## Project Structure
your_project/ <br>
├ cmd/      <br>
│   └ main.go           
├ internal/<br>
│   ├ app/<br>
│   │   └ app.go  
│   ├ data/<br>
│   │   └ data.go     
│   │   └ models.go  
│   └ worker/<br>
│       └ worker.go      
├ tools/<br>
│   └ datagen/<br>
│       └ number_generator.go      
├ go.mod<br>
├ Makefile<br>
└ README.md

## Usage

### Build the Application

To build the application, run:

```bash
make build
```

### Run the Application

To run the application, use:

```bash
make run GOROUTINES=<number_of_goroutines> FILE=<path_to_json_file>
```

For example, to run the application with 4 goroutines and a data file named data.json:
```bash
make run GOROUTINES=4 FILE=data.json
```

### Run Tests
To run the tests, use:

```bash
make test
```

### Generate Test Data

To generate test data, run:
```bash
make generate FILE=<name_for_json_file>
```

Example:
```bash
make generate FILE=data.json
```

## Example JSON File

### Here’s an example of the JSON file format expected by the application:
```json
[
    {
        "a": 1,
        "b": 3
    },
    {
        "a": 5,
        "b": -9
    },
    {
        "a": -2,
        "b": 4
    }
]
```

## Contributing

Feel free to fork this repository and submit pull requests.

## Author

https://t.me/omashev_adlet