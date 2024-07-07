package main

import (
	"flag"
	"fmt"
	"github.com/dexises/calculator/internal/app"
	"log"
	"os"
)

func main() {
	var numGoroutines int
	var fileName string
	flag.IntVar(&numGoroutines, "goroutines", 1, "number of goroutines to use")
	flag.StringVar(&fileName, "file", "data.json", "path to the JSON file")
	flag.Parse()

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatalf("file does not exist: %s", fileName)
	}

	result, err := app.Run(numGoroutines, fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Sum:", result)
}
