package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math"
	"os"
)

// Define a struct to hold the parsed YAML.
type Data map[string]map[string]interface{}

func main() {
	// Define the file path
	filePath := "./test/e2e/scores.yaml"

	// Read the content of the YAML file
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		os.Exit(1)
	}

	// Parse the YAML content
	var data Data
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		panic(err)
	}

	// Compute the average score
	var total float64
	var count float64
	var numGood int
	for _, category := range data {
		for _, value := range category {
			// Check if the value is numeric and add it to the total
			if score, ok := value.(int); ok {
				total += float64(score)
				if score >= 4 {
					numGood++
				}
				count++
			}
		}
	}

	if count == 0 {
		fmt.Println("No numeric scores found.")
		return
	}

	average := total / count
	output := map[string]interface{}{
		"total":           count,
		"average_score":   oneDecimal(average),
		"accuracy":        oneDecimal(average * 20),
		"percentage_good": oneDecimal(float64(numGood) / count * 100),
	}
	b, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(b))
}

func oneDecimal(float float64) float64 {
	return math.Round(float*10) / 10
}
