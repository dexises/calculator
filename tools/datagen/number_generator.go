package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	var objects int
	var fileName string
	flag.IntVar(&objects, "objects", 100000, "number of objects to generate")
	flag.StringVar(&fileName, "file", "data.json", "name of the output JSON file")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	var result int

	data := make([]Data, objects)
	for i := 0; i < objects; i++ {
		data[i] = Data{
			A: rand.Intn(21) - 10,
			B: rand.Intn(21) - 10,
		}

		result += data[i].A + data[i].B
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling data to JSON: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing data to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d objects, wrote to %s and result is %d\n", objects, fileName, result)
}
